package core

import (
	"capturego/config"
	"capturego/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

// DualSaveCapture 영역 지정 캡처 후 파일 저장과 클립보드 복사를 동시에 처리한다
func DualSaveCapture() error {
	// 저장 경로 준비
	savePath, err := resolveSavePath()
	if err != nil {
		return fmt.Errorf("저장 경로 준비 실패: %w", err)
	}

	// 파일명 생성: 캡쳐고_YYYYMMDD_HHMMSS.png
	filename := fmt.Sprintf("캡쳐고_%s.png", time.Now().Format("20060102_150405"))
	filePath := filepath.Join(savePath, filename)

	utils.Info("캡처 시작: %s", filePath)

	// screencapture CLI로 영역 지정 캡처 실행
	// -i: 인터랙티브 모드 (영역 지정), -x: 소리 끄기
	cmd := exec.Command("screencapture", "-i", "-x", filePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("screencapture 실행 실패: %w", err)
	}

	// 사용자가 캡처를 취소한 경우 (파일이 생성되지 않음)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		utils.Info("캡처 취소됨")
		return nil
	}

	// 파일 저장과 클립보드 복사를 고루틴으로 동시 처리
	var wg sync.WaitGroup
	var saveErr, clipErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		// 파일은 이미 screencapture가 저장했으므로 존재 확인만 수행
		if _, err := os.Stat(filePath); err != nil {
			saveErr = fmt.Errorf("저장 파일 확인 실패: %w", err)
			return
		}
		utils.Info("파일 저장 완료: %s", filePath)
	}()

	go func() {
		defer wg.Done()
		if err := copyImageToClipboard(filePath); err != nil {
			clipErr = fmt.Errorf("클립보드 복사 실패: %w", err)
			utils.Error("클립보드 복사 실패: %v", err)
			return
		}
		utils.Info("클립보드 복사 완료")
	}()

	wg.Wait()

	if saveErr != nil {
		return saveErr
	}
	if clipErr != nil {
		// 클립보드 실패는 경고만 기록하고 진행 (파일은 이미 저장됨)
		utils.Warn("클립보드 복사에 실패했지만 파일은 저장되었습니다: %v", clipErr)
	}

	// 누적 캡처 횟수 증가
	if err := config.IncrementCaptureCount(); err != nil {
		utils.Warn("캡처 횟수 기록 실패: %v", err)
	}

	utils.Info("듀얼 세이브 완료: 총 %d회", config.Get().CaptureCount)
	return nil
}

// copyImageToClipboard osascript를 사용해 이미지 파일을 클립보드에 복사한다
func copyImageToClipboard(filePath string) error {
	script := fmt.Sprintf(
		`set the clipboard to (read (POSIX file "%s") as TIFF picture)`,
		filePath,
	)
	cmd := exec.Command("osascript", "-e", script)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("osascript 오류: %s: %w", string(out), err)
	}
	return nil
}

// resolveSavePath 설정된 저장 경로를 반환하고, 유효하지 않으면 기본 경로로 폴백한다
func resolveSavePath() (string, error) {
	cfg := config.Get()
	savePath := cfg.SaveDirectory

	if err := ensureDirectory(savePath); err != nil {
		// 유효하지 않으면 기본 경로로 폴백
		defaultPath := filepath.Join(os.Getenv("HOME"), "Pictures", "CaptureGo")
		utils.Warn("저장 경로 '%s' 사용 불가, 기본 경로로 폴백: %s", savePath, defaultPath)

		if err2 := ensureDirectory(defaultPath); err2 != nil {
			return "", fmt.Errorf("기본 저장 경로 생성 실패: %w", err2)
		}
		return defaultPath, nil
	}
	return savePath, nil
}

// ensureDirectory 디렉토리가 없으면 생성하고, 쓰기 권한을 확인한다
func ensureDirectory(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("디렉토리 생성 실패: %w", err)
	}
	// 쓰기 권한 확인
	testFile := filepath.Join(path, ".write_test")
	f, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("쓰기 권한 없음: %w", err)
	}
	f.Close()
	os.Remove(testFile)
	return nil
}
