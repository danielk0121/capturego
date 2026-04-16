package main

import (
	_ "embed"
	"fmt"
	"os/exec"

	"capturego/config"
	"capturego/core"
	"capturego/server"
	"capturego/utils"

	"github.com/getlantern/systray"
)

//go:embed ui/tray_icon.png
var trayIcon []byte

var (
	webServer *server.WebServer
	hotkeyMgr *core.HotkeyManager
)

func main() {
	// 로거 초기화
	if err := utils.InitLogger(); err != nil {
		panic(err)
	}
	defer utils.CloseLogger()

	// 설정 파일 초기화
	if err := config.Init(); err != nil {
		utils.Error("설정 초기화 실패: %v", err)
		panic(err)
	}
	utils.Info("설정 로드 완료: 저장 경로=%s", config.Get().SaveDirectory)

	// 시스템 트레이 실행 (macOS 요구사항: 메인 스레드 점유)
	systray.Run(onTrayReady, onTrayExit)
}

// onTrayReady 트레이 아이콘이 준비된 후 호출되는 콜백
func onTrayReady() {
	utils.Info("CaptureGo 시작")

	systray.SetIcon(trayIcon)
	systray.SetTooltip("CaptureGo — 듀얼 세이브 캡처")

	// 트레이 메뉴 구성
	mSettings := systray.AddMenuItem("설정 (Settings)", "브라우저로 설정 UI 열기")
	mSupport := systray.AddMenuItem("개발자 응원 (Support)", "후원 페이지로 이동")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("종료 (Quit)", "CaptureGo 종료")

	// 웹서버 시작
	webServer = server.New()
	webServer.Start()

	// macOS 권한 확인 (백그라운드에서 실행)
	go core.CheckPermissions()

	// 글로벌 단축키 등록
	hotkeyMgr = core.NewHotkeyManager()
	cfg := config.Get()
	if err := hotkeyMgr.Start(cfg.HotkeyCapture, cfg.HotkeyScroll); err != nil {
		utils.Warn("단축키 등록 실패: %v", err)
		go core.NotifyAccessibilityRequired()
	}

	// 메뉴 이벤트 처리 (별도 고루틴)
	go func() {
		for {
			select {
			case <-mSettings.ClickedCh:
				utils.Info("트레이: 설정 열기 클릭")
				openBrowser(fmt.Sprintf("http://localhost%s", webServer.Port()))
			case <-mSupport.ClickedCh:
				utils.Info("트레이: 개발자 응원하기 클릭")
				core.OpenSupportPage()
			case <-mQuit.ClickedCh:
				utils.Info("트레이: 앱 종료 클릭")
				systray.Quit()
			}
		}
	}()
}

// onTrayExit 트레이 종료 시 정리 작업
func onTrayExit() {
	if hotkeyMgr != nil {
		hotkeyMgr.Stop()
	}
	if webServer != nil {
		webServer.Stop()
	}
	utils.Info("CaptureGo 종료")
}

// openBrowser macOS 기본 브라우저로 URL을 연다
func openBrowser(url string) {
	if err := exec.Command("open", url).Start(); err != nil {
		utils.Error("브라우저 열기 실패: %v", err)
	}
}
