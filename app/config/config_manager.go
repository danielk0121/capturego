package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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
}

var (
	configPath    string
	currentConfig *Config
)

// defaultConfig 기본 설정값을 반환한다
func defaultConfig() *Config {
	homeDir := os.Getenv("HOME")
	return &Config{
		SaveDirectory:    filepath.Join(homeDir, "Pictures", "CaptureGo"),
		HotkeyCapture:   "ctrl+shift+1",
		HotkeyScroll:    "ctrl+shift+2",
		CaptureCount:    0,
		LicenseActivated: false,
		NagwareDisabled: false,
		LicenseKey:      "",
	}
}

// configFilePath 설정 파일 경로를 반환한다
func configFilePath() string {
	homeDir := os.Getenv("HOME")
	return filepath.Join(homeDir, "Library", "Application Support", "CaptureGo", "config.json")
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
