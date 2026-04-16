//go:build e2e

package core

import (
	"capturego/config"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// e2e 테스트 실행 방법:
//   go test ./core/... -tags e2e -v -run E2E
//
// 실제 macOS 권한(화면 기록, 손쉬운 사용)이 필요하다.

func TestE2E_듀얼세이브_캡처_파일생성(t *testing.T) {
	tmpDir := t.TempDir()

	// 설정 초기화 — 임시 저장 경로 사용
	configPath = filepath.Join(t.TempDir(), "config.json")
	currentConfig = nil
	if err := config.Init(); err != nil {
		t.Fatalf("설정 초기화 실패: %v", err)
	}
	cfg := config.Get()
	cfg.SaveDirectory = tmpDir
	config.Update(cfg)

	t.Log("캡처 영역을 마우스로 지정하세요...")
	if err := DualSaveCapture(); err != nil {
		t.Fatalf("DualSaveCapture 실패: %v", err)
	}

	// 파일 생성 확인
	time.Sleep(500 * time.Millisecond)
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("저장 디렉토리 읽기 실패: %v", err)
	}
	if len(entries) == 0 {
		t.Error("캡처 파일이 생성되지 않았습니다")
	} else {
		t.Logf("캡처 파일 생성 확인: %s", entries[0].Name())
	}
}

func TestE2E_스크롤_캡처_병합파일_생성(t *testing.T) {
	tmpDir := t.TempDir()

	configPath = filepath.Join(t.TempDir(), "config.json")
	currentConfig = nil
	config.Init()
	cfg := config.Get()
	cfg.SaveDirectory = tmpDir
	config.Update(cfg)

	t.Log("스크롤 캡처: 영역을 지정하고 스크롤하세요...")
	if err := ScrollCapture(); err != nil {
		t.Fatalf("ScrollCapture 실패: %v", err)
	}

	time.Sleep(500 * time.Millisecond)
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("저장 디렉토리 읽기 실패: %v", err)
	}
	if len(entries) == 0 {
		t.Error("스크롤 캡처 파일이 생성되지 않았습니다")
	} else {
		t.Logf("스크롤 캡처 파일 생성 확인: %s", entries[0].Name())
	}
}

func TestE2E_Nagware_임계값_도달시_알림(t *testing.T) {
	configPath = filepath.Join(t.TempDir(), "config.json")
	currentConfig = nil
	config.Init()

	// 임계값까지 카운트 올리기
	cfg := config.Get()
	cfg.CaptureCount = nagwareThreshold
	config.Update(cfg)

	t.Log("Nagware 알림이 표시되어야 합니다...")
	CheckNagware()
	t.Log("알림 표시 완료 (수동 확인 필요)")
}

func TestE2E_설정_저장_경로_변경후_캡처_위치확인(t *testing.T) {
	tmpDir1 := t.TempDir()
	tmpDir2 := t.TempDir()

	configPath = filepath.Join(t.TempDir(), "config.json")
	currentConfig = nil
	config.Init()

	// 첫 번째 경로로 캡처
	cfg := config.Get()
	cfg.SaveDirectory = tmpDir1
	config.Update(cfg)

	t.Log("첫 번째 경로로 캡처하세요...")
	DualSaveCapture()

	// 경로 변경 후 캡처
	cfg.SaveDirectory = tmpDir2
	config.Update(cfg)

	t.Log("두 번째 경로로 캡처하세요...")
	DualSaveCapture()

	// 각 경로에 파일이 있는지 확인
	for i, dir := range []string{tmpDir1, tmpDir2} {
		entries, _ := os.ReadDir(dir)
		if len(entries) == 0 {
			t.Errorf("경로 %d (%s)에 파일이 없습니다", i+1, dir)
		} else {
			t.Logf("경로 %d 파일 확인: %s", i+1, entries[0].Name())
		}
	}
	fmt.Println("경로 변경 테스트 완료")
}
