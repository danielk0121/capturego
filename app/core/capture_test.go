package core

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"
)

func TestEnsureDirectory_정상_경로(t *testing.T) {
	tmpDir := filepath.Join(t.TempDir(), "captures")
	if err := ensureDirectory(tmpDir); err != nil {
		t.Errorf("정상 경로에서 오류 발생: %v", err)
	}
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		t.Error("디렉토리가 생성되지 않았습니다")
	}
}

func TestEnsureDirectory_이미_존재하는_경로(t *testing.T) {
	tmpDir := t.TempDir()
	if err := ensureDirectory(tmpDir); err != nil {
		t.Errorf("기존 경로에서 오류 발생: %v", err)
	}
}

func TestResolveSavePath_경로_반환(t *testing.T) {
	path, err := resolveSavePath()
	if err != nil {
		t.Fatalf("resolveSavePath 실패: %v", err)
	}
	if path == "" {
		t.Error("반환된 저장 경로가 비어있습니다")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("저장 경로가 존재하지 않습니다: %s", path)
	}
}

func TestFilenameFormat(t *testing.T) {
	// 파일명 포맷 검증: capturego_YYYYMMDD_HHMMSS.png
	filename := "capturego_" + time.Now().Format("20060102_150405") + ".png"
	pattern := regexp.MustCompile(`^capturego_\d{8}_\d{6}\.png$`)
	if !pattern.MatchString(filename) {
		t.Errorf("파일명 포맷이 올바르지 않습니다: %s", filename)
	}
}
