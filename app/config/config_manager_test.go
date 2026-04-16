package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func setupTestConfigPath(t *testing.T) func() {
	t.Helper()
	// 테스트용 임시 경로로 교체
	tmpDir := t.TempDir()
	origPath := configPath
	origConfig := currentConfig

	configPath = filepath.Join(tmpDir, "config.json")
	currentConfig = nil

	return func() {
		configPath = origPath
		currentConfig = origConfig
	}
}

func TestInit_최초실행시_기본설정_파일_생성(t *testing.T) {
	cleanup := setupTestConfigPath(t)
	defer cleanup()

	if err := Init(); err != nil {
		t.Fatalf("Init 실패: %v", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("설정 파일이 생성되지 않았습니다")
	}

	cfg := Get()
	if cfg.HotkeyCapture != "cmd+shift+6" {
		t.Errorf("기본 단축키가 올바르지 않습니다: %s", cfg.HotkeyCapture)
	}
	if cfg.CaptureCount != 0 {
		t.Errorf("초기 캡처 횟수는 0이어야 합니다: %d", cfg.CaptureCount)
	}
}

func TestInit_기존설정_파일_로드(t *testing.T) {
	cleanup := setupTestConfigPath(t)
	defer cleanup()

	// 미리 설정 파일 작성
	preset := defaultConfig()
	preset.HotkeyCapture = "ctrl+shift+9"
	preset.CaptureCount = 5
	data, _ := json.MarshalIndent(preset, "", "  ")
	os.MkdirAll(filepath.Dir(configPath), 0755)
	os.WriteFile(configPath, data, 0644)

	if err := Init(); err != nil {
		t.Fatalf("Init 실패: %v", err)
	}

	cfg := Get()
	if cfg.HotkeyCapture != "ctrl+shift+9" {
		t.Errorf("저장된 단축키가 로드되지 않았습니다: %s", cfg.HotkeyCapture)
	}
	if cfg.CaptureCount != 5 {
		t.Errorf("저장된 캡처 횟수가 로드되지 않았습니다: %d", cfg.CaptureCount)
	}
}

func TestUpdate_설정_갱신_및_저장(t *testing.T) {
	cleanup := setupTestConfigPath(t)
	defer cleanup()

	Init()

	cfg := Get()
	cfg.SaveDirectory = "/tmp/test_captures"
	if err := Update(cfg); err != nil {
		t.Fatalf("Update 실패: %v", err)
	}

	// 파일에서 다시 읽어 확인
	currentConfig = nil
	load()
	if Get().SaveDirectory != "/tmp/test_captures" {
		t.Errorf("갱신된 저장 경로가 파일에 반영되지 않았습니다")
	}
}

func TestIncrementCaptureCount(t *testing.T) {
	cleanup := setupTestConfigPath(t)
	defer cleanup()

	Init()

	for i := 1; i <= 3; i++ {
		if err := IncrementCaptureCount(); err != nil {
			t.Fatalf("횟수 증가 실패: %v", err)
		}
		if Get().CaptureCount != i {
			t.Errorf("캡처 횟수가 %d여야 하는데 %d입니다", i, Get().CaptureCount)
		}
	}
}
