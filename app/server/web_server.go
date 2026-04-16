package server

import (
	"capturego/config"
	"capturego/core"
	"capturego/utils"
	"context"
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed static
var staticFiles embed.FS

const defaultPort = ":18420"

// HotkeyReloader 단축키 재등록 인터페이스
type HotkeyReloader interface {
	Reload(captureKey, scrollKey string) error
}

// WebServer Gin 기반 로컬 HTTP 서버
type WebServer struct {
	engine    *gin.Engine
	srv       *http.Server
	hotkeyMgr HotkeyReloader
}

// New 웹서버 인스턴스를 생성한다
func New(hotkeyMgr HotkeyReloader) *WebServer {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	ws := &WebServer{engine: r, hotkeyMgr: hotkeyMgr}
	ws.registerRoutes()
	return ws
}

// Start 서버를 백그라운드 고루틴에서 시작한다
func (ws *WebServer) Start() {
	ws.srv = &http.Server{
		Addr:    defaultPort,
		Handler: ws.engine,
	}
	go func() {
		utils.Info("웹서버 시작: http://localhost%s", defaultPort)
		if err := ws.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Error("웹서버 오류: %v", err)
		}
	}()
}

// Stop 서버를 정상 종료한다
func (ws *WebServer) Stop() {
	if ws.srv == nil {
		return
	}
	if err := ws.srv.Shutdown(context.Background()); err != nil {
		utils.Error("웹서버 종료 오류: %v", err)
	}
	utils.Info("웹서버 종료 완료")
}

// Port 서버가 바인딩된 포트 문자열을 반환한다
func (ws *WebServer) Port() string {
	return defaultPort
}

func (ws *WebServer) registerRoutes() {
	// 노캐시 미들웨어 전역 적용
	ws.engine.Use(noCacheMiddleware())

	ws.engine.GET("/", serveIndex)
	ws.engine.GET("/license", serveLicense)
	ws.engine.GET("/api/config", getConfig)
	ws.engine.POST("/api/config", ws.postConfig)
	ws.engine.GET("/api/permissions", getPermissions)
	ws.engine.GET("/api/buildtime", getBuildtime)

	// 정적 파일 서빙 (favicon 등)
	sub, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic("static FS 초기화 실패: " + err.Error())
	}
	ws.engine.StaticFS("/static", http.FS(sub))
}

// noCacheMiddleware 모든 응답에 캐시 비활성화 헤더를 추가한다
func noCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Next()
	}
}

// serveIndex 설정 UI HTML을 서빙한다
func serveIndex(c *gin.Context) {
	data, err := staticFiles.ReadFile("static/index.html")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

// serveLicense 오픈소스 라이선스 페이지를 서빙한다
func serveLicense(c *gin.Context) {
	data, err := staticFiles.ReadFile("static/license.html")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

// getBuildtime 빌드타임 문자열을 JSON으로 반환한다
func getBuildtime(c *gin.Context) {
	data, err := staticFiles.ReadFile("static/buildtime.txt")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"buildtime": "v-dev"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buildtime": strings.TrimSpace(string(data))})
}

// getPermissions macOS 권한 상태를 JSON으로 반환한다
func getPermissions(c *gin.Context) {
	status := core.QueryPermissions()
	c.JSON(http.StatusOK, status)
}

// getConfig 현재 설정값을 JSON으로 반환한다
func getConfig(c *gin.Context) {
	cfg := config.Get()
	c.JSON(http.StatusOK, gin.H{
		"save_directory":    cfg.SaveDirectory,
		"hotkey_capture":    cfg.HotkeyCapture,
		"hotkey_scroll":     cfg.HotkeyScroll,
		"capture_count":     cfg.CaptureCount,
		"license_activated": cfg.LicenseActivated,
		"nagware_disabled":  cfg.NagwareDisabled,
	})
}

// postConfig 설정값을 저장한다
func (ws *WebServer) postConfig(c *gin.Context) {
	var body struct {
		SaveDirectory string `json:"save_directory"`
		HotkeyCapture string `json:"hotkey_capture"`
		HotkeyScroll  string `json:"hotkey_scroll"`
		LicenseKey    string `json:"license_key"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식입니다"})
		return
	}

	cfg := config.Get()
	hotkeyChanged := false

	if body.SaveDirectory != "" {
		cfg.SaveDirectory = body.SaveDirectory
	}
	if body.HotkeyCapture != "" && body.HotkeyCapture != cfg.HotkeyCapture {
		cfg.HotkeyCapture = body.HotkeyCapture
		hotkeyChanged = true
	}
	if body.HotkeyScroll != "" && body.HotkeyScroll != cfg.HotkeyScroll {
		cfg.HotkeyScroll = body.HotkeyScroll
		hotkeyChanged = true
	}
	// 라이선스 키 입력 시 별도 검증 처리
	if body.LicenseKey != "" {
		if err := core.ActivateLicense(body.LicenseKey); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := config.Update(cfg); err != nil {
		utils.Error("설정 저장 실패: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "설정 저장에 실패했습니다"})
		return
	}

	// 단축키가 변경된 경우 즉시 재등록
	if hotkeyChanged && ws.hotkeyMgr != nil {
		if err := ws.hotkeyMgr.Reload(cfg.HotkeyCapture, cfg.HotkeyScroll); err != nil {
			utils.Warn("단축키 재등록 실패: %v", err)
		} else {
			utils.Info("단축키 재등록 완료: 캡처=%s, 스크롤=%s", cfg.HotkeyCapture, cfg.HotkeyScroll)
		}
	}

	utils.Info("설정 저장 완료")
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
