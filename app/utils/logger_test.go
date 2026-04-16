package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitLogger(t *testing.T) {
	err := InitLogger()
	if err != nil {
		t.Fatalf("로거 초기화 실패: %v", err)
	}
	defer CloseLogger()

	logPath := filepath.Join(os.Getenv("HOME"), "Library", "Logs", "CaptureGo", "capturego.log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Errorf("로그 파일이 생성되지 않았습니다: %s", logPath)
	}
}

func TestLogLevels(t *testing.T) {
	if err := InitLogger(); err != nil {
		t.Fatalf("로거 초기화 실패: %v", err)
	}
	defer CloseLogger()

	// 패닉 없이 각 레벨 로그가 기록되어야 한다
	Info("테스트 INFO 로그: %s", "정상")
	Warn("테스트 WARN 로그: %d", 42)
	Error("테스트 ERROR 로그: %v", true)
}

func TestLoggerWithoutInit(t *testing.T) {
	// 초기화 없이 호출해도 패닉이 발생하지 않아야 한다
	saved := defaultLogger
	defaultLogger = nil
	defer func() { defaultLogger = saved }()

	Info("초기화 없는 INFO")
	Warn("초기화 없는 WARN")
	Error("초기화 없는 ERROR")
}
