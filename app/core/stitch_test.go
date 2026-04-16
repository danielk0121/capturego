package core

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

// makeTestImage 단색 PNG 이미지 파일을 생성하고 경로를 반환한다
func makeTestImage(t *testing.T, dir, name string, w, h int, c color.RGBA) string {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, c)
		}
	}
	path := filepath.Join(dir, name)
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		t.Fatal(err)
	}
	return path
}

// makeGradientImage 세로 그라데이션 PNG 이미지 파일을 생성한다
func makeGradientImage(t *testing.T, dir, name string, w, h int, startY uint8) string {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		v := startY + uint8(y)
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	path := filepath.Join(dir, name)
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestStitchImages_단일_이미지(t *testing.T) {
	dir := t.TempDir()
	src := makeTestImage(t, dir, "a.png", 100, 50, color.RGBA{255, 0, 0, 255})
	out := filepath.Join(dir, "out.png")

	if err := StitchImages([]string{src}, out); err != nil {
		t.Fatalf("StitchImages 실패: %v", err)
	}

	bounds, err := readImageBounds(out)
	if err != nil {
		t.Fatalf("결과 이미지 읽기 실패: %v", err)
	}
	if bounds.Dy() != 50 {
		t.Errorf("높이가 50이어야 하는데 %d입니다", bounds.Dy())
	}
}

func TestStitchImages_겹침없는_병합(t *testing.T) {
	dir := t.TempDir()
	// 두 이미지가 완전히 다른 색상 → 겹침 0 → 총 높이 = 합산
	a := makeTestImage(t, dir, "a.png", 100, 40, color.RGBA{255, 0, 0, 255})
	b := makeTestImage(t, dir, "b.png", 100, 40, color.RGBA{0, 255, 0, 255})
	out := filepath.Join(dir, "out.png")

	if err := StitchImages([]string{a, b}, out); err != nil {
		t.Fatalf("StitchImages 실패: %v", err)
	}

	bounds, err := readImageBounds(out)
	if err != nil {
		t.Fatalf("결과 이미지 읽기 실패: %v", err)
	}
	// 겹침이 없으므로 총 높이 = 40 + 40 = 80
	if bounds.Dy() != 80 {
		t.Errorf("높이가 80이어야 하는데 %d입니다", bounds.Dy())
	}
}

func TestStitchImages_겹침_감지(t *testing.T) {
	dir := t.TempDir()
	// a: startY=0 (행값 0~29), b: startY=20 (행값 20~49)
	// a 하단 10행(값 20~29) = b 상단 10행(값 20~29) → 최소 10행 이상 겹침 감지 기대
	a := makeGradientImage(t, dir, "a.png", 100, 30, 0)
	b := makeGradientImage(t, dir, "b.png", 100, 30, 20)
	out := filepath.Join(dir, "out.png")

	if err := StitchImages([]string{a, b}, out); err != nil {
		t.Fatalf("StitchImages 실패: %v", err)
	}

	bounds, err := readImageBounds(out)
	if err != nil {
		t.Fatalf("결과 이미지 읽기 실패: %v", err)
	}
	// 겹침이 적어도 10행 이상 감지되어야 하므로 결과 높이 < 60
	totalNoOverlap := 30 + 30
	if bounds.Dy() >= totalNoOverlap {
		t.Errorf("겹침이 감지되지 않았습니다: 높이=%d (기대 < %d)", bounds.Dy(), totalNoOverlap)
	}
}

func TestStitchImages_빈_목록_오류(t *testing.T) {
	out := filepath.Join(t.TempDir(), "out.png")
	if err := StitchImages([]string{}, out); err == nil {
		t.Error("빈 목록에서 오류가 발생해야 합니다")
	}
}

func TestFramesIdentical_동일_이미지(t *testing.T) {
	dir := t.TempDir()
	a := makeTestImage(t, dir, "a.png", 50, 50, color.RGBA{128, 128, 128, 255})
	b := makeTestImage(t, dir, "b.png", 50, 50, color.RGBA{128, 128, 128, 255})

	if !framesIdentical(a, b) {
		t.Error("동일한 이미지를 다르다고 판단했습니다")
	}
}

func TestFramesIdentical_다른_이미지(t *testing.T) {
	dir := t.TempDir()
	a := makeTestImage(t, dir, "a.png", 50, 50, color.RGBA{255, 0, 0, 255})
	b := makeTestImage(t, dir, "b.png", 50, 50, color.RGBA{0, 0, 255, 255})

	if framesIdentical(a, b) {
		t.Error("다른 이미지를 동일하다고 판단했습니다")
	}
}
