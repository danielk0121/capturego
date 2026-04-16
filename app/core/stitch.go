package core

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// StitchImages 이미지 파일 목록을 세로로 병합한다.
// 인접한 이미지 사이의 겹치는 픽셀 영역을 계산해 절삭 후 이어붙인다.
func StitchImages(paths []string, outputPath string) error {
	if len(paths) == 0 {
		return fmt.Errorf("병합할 이미지가 없습니다")
	}

	images, err := loadImages(paths)
	if err != nil {
		return err
	}

	// 겹침 영역 계산 후 유효 높이 목록 산출
	heights := computeEffectiveHeights(images)

	totalHeight := 0
	for _, h := range heights {
		totalHeight += h
	}
	width := images[0].Bounds().Dx()

	// 결과 이미지 생성
	result := image.NewRGBA(image.Rect(0, 0, width, totalHeight))
	y := 0
	for i, img := range images {
		src := img.(*image.NRGBA)
		cropRect := image.Rect(0, 0, width, heights[i])
		dst := result.SubImage(image.Rect(0, y, width, y+heights[i])).(*image.RGBA)
		draw.Draw(dst, dst.Bounds(), src, cropRect.Min, draw.Src)
		y += heights[i]
	}

	// 파일 저장
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("출력 파일 생성 실패: %w", err)
	}
	defer f.Close()

	if err := png.Encode(f, result); err != nil {
		return fmt.Errorf("PNG 인코딩 실패: %w", err)
	}
	return nil
}

// loadImages 파일 경로 목록에서 이미지를 디코딩하여 반환한다
func loadImages(paths []string) ([]image.Image, error) {
	images := make([]image.Image, 0, len(paths))
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return nil, fmt.Errorf("이미지 파일 열기 실패 '%s': %w", p, err)
		}
		img, _, err := image.Decode(f)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("이미지 디코딩 실패 '%s': %w", p, err)
		}

		// draw.Draw 호환을 위해 *image.NRGBA로 변환
		b := img.Bounds()
		nrgba := image.NewNRGBA(b)
		draw.Draw(nrgba, b, img, b.Min, draw.Src)
		images = append(images, nrgba)
	}
	return images, nil
}

// computeEffectiveHeights 인접 이미지 간 겹치는 행을 제거한 유효 높이를 반환한다.
// 마지막 이미지는 전체 높이를 사용한다.
func computeEffectiveHeights(images []image.Image) []int {
	heights := make([]int, len(images))
	for i, img := range images {
		heights[i] = img.Bounds().Dy()
	}

	for i := 0; i < len(images)-1; i++ {
		overlap := findOverlap(images[i], images[i+1])
		if overlap > 0 {
			heights[i] -= overlap
		}
	}
	return heights
}

// findOverlap 두 이미지 간 겹치는 행 수를 반환한다.
// images[i]의 하단 k행이 images[i+1]의 상단 k행과 일치하는 최대 k를 찾는다.
func findOverlap(top, bottom image.Image) int {
	topH := top.Bounds().Dy()
	botH := bottom.Bounds().Dy()
	w := top.Bounds().Dx()
	if bw := bottom.Bounds().Dx(); bw < w {
		w = bw
	}

	// 탐색할 최대 겹침 범위 (전체 높이의 60% 이하)
	maxOverlap := topH * 60 / 100
	if maxOverlap > botH {
		maxOverlap = botH
	}

	best := 0
	for k := maxOverlap; k >= 4; k-- {
		if rowsMatch(top, topH-k, bottom, 0, k, w) {
			best = k
			break
		}
	}
	return best
}

// rowsMatch top 이미지의 startTop 행부터 count행이 bottom 이미지의 startBot 행부터와 일치하는지 확인한다.
// 픽셀 오차 허용값(tolerance)을 적용해 JPEG 압축 노이즈를 허용한다.
func rowsMatch(top image.Image, startTop int, bottom image.Image, startBot int, count int, w int) bool {
	const tolerance = 8
	const maxMismatch = 3 // 행당 최대 불일치 픽셀 수

	for row := 0; row < count; row++ {
		mismatch := 0
		for x := 0; x < w; x++ {
			r1, g1, b1, _ := top.At(x, startTop+row).RGBA()
			r2, g2, b2, _ := bottom.At(x, startBot+row).RGBA()
			if diff(r1, r2) > tolerance*257 ||
				diff(g1, g2) > tolerance*257 ||
				diff(b1, b2) > tolerance*257 {
				mismatch++
				if mismatch > maxMismatch {
					return false
				}
			}
		}
	}
	return true
}

func diff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}
