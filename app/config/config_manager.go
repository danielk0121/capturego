package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

//go:embed default-config.json
var defaultConfigJSON []byte

// Config 앱 전체 설정 스키마
type Config struct {
	// 캡처 파일 저장 디렉토리 경로
	SaveDirectory string `json:"save_directory"`
	// 듀얼 세이브 캡처 단축키 (예: "ctrl+shift+1")
	HotkeyCapture string `json:"hotkey_capture"`
	// 스크롤 캡처 단축키
	HotkeyScroll string `json:"hotkey_scroll"`
	// 누적 캡처 횟수
	CaptureCount int `json:"capture_count"`
	// 라이선스 인증 여부
	LicenseActivated bool `json:"license_activated"`
	// Nagware 팝업 비활성화 여부
	NagwareDisabled bool `json:"nagware_disabled"`
	// 라이선스 키
	LicenseKey string `json:"license_key"`
	// 다크모드 여부 (false: 라이트모드, true: 다크모드)
	DarkMode bool `json:"dark_mode"`
	// UI 언어 ("ko" 또는 "en")
	Language string `json:"language"`
}

var (
	configPath    string
	currentConfig *Config
)

// homeDir 현재 사용자의 홈 디렉토리를 반환한다.
// os/user.Current()를 우선 사용하고, 실패 시 HOME 환경변수로 폴백한다.
func homeDir() string {
	if u, err := user.Current(); err == nil && u.HomeDir != "" {
		return u.HomeDir
	}
	return os.Getenv("HOME")
}

// defaultConfig default-config.json을 파싱해 기본 설정값을 반환한다.
// save_directory의 선행 '~'는 런타임 홈 디렉토리로 확장한다.
func defaultConfig() *Config {
	cfg := &Config{}
	if err := json.Unmarshal(defaultConfigJSON, cfg); err != nil {
		panic(fmt.Sprintf("default-config.json 파싱 실패: %v", err))
	}
	if cfg.SaveDirectory != "" && cfg.SaveDirectory[0] == '~' {
		cfg.SaveDirectory = filepath.Join(homeDir(), cfg.SaveDirectory[1:])
	}
	return cfg
}

// ResetToDefault 설정을 default-config.json 기본값으로 초기화하고 저장한다.
// capture_count, license_activated, license_key, nagware_disabled 는 유지한다.
func ResetToDefault() error {
	def := defaultConfig()
	if currentConfig != nil {
		def.CaptureCount     = currentConfig.CaptureCount
		def.LicenseActivated = currentConfig.LicenseActivated
		def.LicenseKey       = currentConfig.LicenseKey
		def.NagwareDisabled  = currentConfig.NagwareDisabled
	}
	currentConfig = def
	return save(currentConfig)
}

// configFilePath 설정 파일 경로를 반환한다
func configFilePath() string {
	return filepath.Join(homeDir(), "Library", "Application Support", "CaptureGo", "config.json")
}

// Init 설정 파일을 로드하거나 최초 실행 시 기본 설정 파일을 생성한다
func Init() error {
	// 테스트에서 미리 configPath를 설정한 경우 덮어쓰지 않는다
	if configPath == "" {
		configPath = configFilePath()
	}

	// 설정 디렉토리 생성
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("설정 디렉토리 생성 실패: %w", err)
	}

	// 설정 파일이 없으면 기본값으로 생성
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		currentConfig = defaultConfig()
		return save(currentConfig)
	}

	return load()
}

// load 파일에서 설정을 읽는다
func load() error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("설정 파일 읽기 실패: %w", err)
	}

	cfg := defaultConfig()
	if err := json.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("설정 파일 파싱 실패: %w", err)
	}
	currentConfig = cfg
	return nil
}

// save 설정을 파일에 저장한다
func save(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("설정 직렬화 실패: %w", err)
	}
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("설정 파일 저장 실패: %w", err)
	}
	return nil
}

// Get 현재 설정 복사본을 반환한다
func Get() Config {
	if currentConfig == nil {
		return *defaultConfig()
	}
	return *currentConfig
}

// Update 설정을 갱신하고 파일에 저장한다
func Update(cfg Config) error {
	currentConfig = &cfg
	return save(currentConfig)
}

// IncrementCaptureCount 누적 캡처 횟수를 1 증가시키고 저장한다
func IncrementCaptureCount() error {
	if currentConfig == nil {
		return fmt.Errorf("설정이 초기화되지 않았습니다")
	}
	currentConfig.CaptureCount++
	return save(currentConfig)
}
