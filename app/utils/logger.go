package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 로그 레벨 정의
type LogLevel int

const (
	LevelInfo  LogLevel = iota
	LevelWarn
	LevelError
)

// Logger 앱 전역 로거
type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	file        *os.File
}

var defaultLogger *Logger

// InitLogger 로거를 초기화하고 로그 파일을 생성한다
func InitLogger() error {
	logDir := filepath.Join(os.Getenv("HOME"), "Library", "Logs", "CaptureGo")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("로그 디렉토리 생성 실패: %w", err)
	}

	logPath := filepath.Join(logDir, "capturego.log")
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("로그 파일 열기 실패: %w", err)
	}

	// 파일과 표준 출력에 동시 기록
	mw := io.MultiWriter(f, os.Stdout)
	flags := log.Ldate | log.Ltime | log.Lmsgprefix

	defaultLogger = &Logger{
		infoLogger:  log.New(mw, "[INFO]  ", flags),
		warnLogger:  log.New(mw, "[WARN]  ", flags),
		errorLogger: log.New(mw, "[ERROR] ", flags),
		file:        f,
	}

	Info("로거 초기화 완료: %s", logPath)
	return nil
}

// CloseLogger 앱 종료 시 로그 파일을 닫는다
func CloseLogger() {
	if defaultLogger != nil && defaultLogger.file != nil {
		Info("로거 종료")
		defaultLogger.file.Close()
	}
}

// rotateIfNeeded 파일이 10MB를 초과하면 날짜 기반 백업 후 새 파일 생성
func rotateIfNeeded() {
	if defaultLogger == nil || defaultLogger.file == nil {
		return
	}
	info, err := defaultLogger.file.Stat()
	if err != nil || info.Size() < 10*1024*1024 {
		return
	}

	home := os.Getenv("HOME")
	logDir := filepath.Join(home, "Library", "Application Support", "CaptureGo", "logs")
	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(logDir, fmt.Sprintf("capturego_%s.log", timestamp))

	defaultLogger.file.Close()
	os.Rename(filepath.Join(logDir, "capturego.log"), backupPath)

	f, err := os.OpenFile(filepath.Join(logDir, "capturego.log"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	mw := io.MultiWriter(f, os.Stdout)
	defaultLogger.file = f
	defaultLogger.infoLogger.SetOutput(mw)
	defaultLogger.warnLogger.SetOutput(mw)
	defaultLogger.errorLogger.SetOutput(mw)
}

// Info INFO 레벨 로그를 기록한다
func Info(format string, args ...any) {
	if defaultLogger == nil {
		fmt.Printf("[INFO] "+format+"\n", args...)
		return
	}
	rotateIfNeeded()
	defaultLogger.infoLogger.Printf(" "+format, args...)
}

// Warn WARN 레벨 로그를 기록한다
func Warn(format string, args ...any) {
	if defaultLogger == nil {
		fmt.Printf("[WARN] "+format+"\n", args...)
		return
	}
	rotateIfNeeded()
	defaultLogger.warnLogger.Printf(" "+format, args...)
}

// Error ERROR 레벨 로그를 기록한다
func Error(format string, args ...any) {
	if defaultLogger == nil {
		fmt.Printf("[ERROR] "+format+"\n", args...)
		return
	}
	rotateIfNeeded()
	defaultLogger.errorLogger.Printf(" "+format, args...)
}
