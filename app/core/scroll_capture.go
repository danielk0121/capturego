package core

import (
	"capturego/utils"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	// 연속 촬영 간격
	scrollCaptureInterval = 300 * time.Millisecond
	// 최대 촬영 장수
	maxScrollFrames = 30
)

// ScrollCapture 스크롤 캡처를 실행한다.
// 사용자가 영역을 지정하면 스크롤하며 연속 촬영 후 이미지를 병합해 저장한다.
func ScrollCapture() error {
	savePath, err := resolveSavePath()
	if err != nil {
		return fmt.Errorf("저장 경로 준비 실패: %w", err)
	}

	// 임시 작업 디렉토리 생성
	tmpDir, err := os.MkdirTemp("", "capturego_scroll_*")
	if err != nil {
		return fmt.Errorf("임시 디렉토리 생성 실패: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	utils.Info("스크롤 캡처 시작: 스크롤하면 자동으로 촬영됩니다")

	frames, err := captureScrollFrames(tmpDir)
	if err != nil {
		return fmt.Errorf("스크롤 프레임 촬영 실패: %w", err)
	}

	if len(frames) == 0 {
		utils.Info("스크롤 캡처 취소됨")
		return nil
	}

	outputName := fmt.Sprintf("캡쳐고_스크롤_%s.png", time.Now().Format("20060102_150405"))
	outputPath := filepath.Join(savePath, outputName)

	if len(frames) == 1 {
		// 단일 프레임은 병합 없이 그대로 복사
		return copyFile(frames[0], outputPath)
	}

	// 이미지 병합
	utils.Info("이미지 병합 중: %d장", len(frames))
	if err := StitchImages(frames, outputPath); err != nil {
		return fmt.Errorf("이미지 병합 실패: %w", err)
	}

	utils.Info("스크롤 캡처 완료: %s", outputPath)
	return nil
}

// captureScrollFrames 영역 지정 후 스크롤 간격으로 연속 촬영한다.
func captureScrollFrames(tmpDir string) ([]string, error) {
	var frames []string

	// 첫 번째 프레임: 인터랙티브 영역 지정
	firstFrame := filepath.Join(tmpDir, "frame_000.png")
	if err := exec.Command("screencapture", "-i", "-x", firstFrame).Run(); err != nil {
		return nil, fmt.Errorf("첫 프레임 캡처 실패: %w", err)
	}
	if _, err := os.Stat(firstFrame); os.IsNotExist(err) {
		return nil, nil // 취소됨
	}
	frames = append(frames, firstFrame)

	// 캡처된 이미지의 크기를 읽어 반복 캡처 영역으로 사용
	bounds, err := readImageBounds(firstFrame)
	if err != nil {
		utils.Warn("캡처 영역 감지 실패, 단일 프레임으로 처리: %v", err)
		return frames, nil
	}

	utils.Info("첫 프레임 캡처 완료. 스크롤하세요 (최대 %d장, %v 간격)", maxScrollFrames, scrollCaptureInterval)

	// 이후 프레임: screencapture -R 로 같은 영역 반복 촬영
	for i := 1; i < maxScrollFrames; i++ {
		time.Sleep(scrollCaptureInterval)

		framePath := filepath.Join(tmpDir, fmt.Sprintf("frame_%03d.png", i))
		region := fmt.Sprintf("%d,%d,%d,%d",
			bounds.Min.X, bounds.Min.Y, bounds.Dx(), bounds.Dy())
		if err := exec.Command("screencapture", "-x", "-R", region, framePath).Run(); err != nil {
			utils.Warn("프레임 %d 캡처 실패: %v", i, err)
			break
		}

		// 이전 프레임과 동일 → 스크롤 종료로 판단
		if framesIdentical(frames[len(frames)-1], framePath) {
			utils.Info("스크롤 종료 감지 (프레임 %d)", i)
			os.Remove(framePath)
			break
		}

		frames = append(frames, framePath)
		utils.Info("프레임 %d 캡처", i)
	}

	return frames, nil
}

// readImageBounds PNG 파일의 이미지 크기를 반환한다
func readImageBounds(path string) (image.Rectangle, error) {
	f, err := os.Open(path)
	if err != nil {
		return image.Rectangle{}, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return image.Rectangle{}, err
	}
	return img.Bounds(), nil
}

// framesIdentical 두 프레임이 동일한지 픽셀 샘플링으로 빠르게 비교한다
func framesIdentical(pathA, pathB string) bool {
	fa, err := os.Open(pathA)
	if err != nil {
		return false
	}
	defer fa.Close()
	fb, err := os.Open(pathB)
	if err != nil {
		return false
	}
	defer fb.Close()

	imgA, _, err := image.Decode(fa)
	if err != nil {
		return false
	}
	imgB, _, err := image.Decode(fb)
	if err != nil {
		return false
	}

	bA, bB := imgA.Bounds(), imgB.Bounds()
	if bA != bB {
		return false
	}

	// 16×16 격자 샘플링으로 빠르게 비교
	w, h := bA.Dx(), bA.Dy()
	for row := 0; row < 16; row++ {
		for col := 0; col < 16; col++ {
			x := col * w / 16
			y := row * h / 16
			r1, g1, b1, _ := imgA.At(x, y).RGBA()
			r2, g2, b2, _ := imgB.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 {
				return false
			}
		}
	}
	return true
}

// copyFile 파일을 src에서 dst로 복사한다
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
