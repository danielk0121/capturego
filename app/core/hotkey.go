package core

import (
	"capturego/utils"
	"context"
	"fmt"
	"strings"

	"golang.design/x/hotkey"
)

// HotkeyManager 글로벌 단축키 등록 및 해제를 관리한다
type HotkeyManager struct {
	captureHK *hotkey.Hotkey
	scrollHK  *hotkey.Hotkey
	cancel    context.CancelFunc
}

// NewHotkeyManager 새로운 HotkeyManager를 생성한다
func NewHotkeyManager() *HotkeyManager {
	return &HotkeyManager{}
}

// Start 설정에서 단축키를 읽어 등록하고 이벤트 리스닝을 시작한다
func (m *HotkeyManager) Start(captureKey, scrollKey string) error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel

	// 듀얼 세이브 캡처 단축키 등록
	captureHK, err := parseAndRegister(captureKey)
	if err != nil {
		cancel()
		return fmt.Errorf("캡처 단축키 등록 실패 ('%s'): %w", captureKey, err)
	}
	m.captureHK = captureHK

	// 스크롤 캡처 단축키 등록
	scrollHK, err := parseAndRegister(scrollKey)
	if err != nil {
		m.captureHK.Unregister()
		cancel()
		return fmt.Errorf("스크롤 단축키 등록 실패 ('%s'): %w", scrollKey, err)
	}
	m.scrollHK = scrollHK

	utils.Info("단축키 등록 완료: 캡처=%s, 스크롤=%s", captureKey, scrollKey)

	// 이벤트 리스닝 고루틴
	go m.listenCapture(ctx)
	go m.listenScroll(ctx)

	return nil
}

// Stop 등록된 단축키를 모두 해제한다
func (m *HotkeyManager) Stop() {
	if m.cancel != nil {
		m.cancel()
	}
	if m.captureHK != nil {
		m.captureHK.Unregister()
	}
	if m.scrollHK != nil {
		m.scrollHK.Unregister()
	}
	utils.Info("단축키 해제 완료")
}

// Reload 단축키를 새 설정으로 재등록한다
func (m *HotkeyManager) Reload(captureKey, scrollKey string) error {
	m.Stop()
	return m.Start(captureKey, scrollKey)
}

func (m *HotkeyManager) listenCapture(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-m.captureHK.Keydown():
			utils.Info("캡처 단축키 입력")
			if err := DualSaveCapture(); err != nil {
				utils.Error("듀얼 세이브 캡처 실패: %v", err)
			}
		}
	}
}

func (m *HotkeyManager) listenScroll(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-m.scrollHK.Keydown():
			utils.Info("스크롤 캡처 단축키 입력")
			// TODO: 스크롤 캡처 호출 (todo_06)
		}
	}
}

// parseAndRegister 문자열 단축키를 파싱하여 hotkey를 등록한다
// 지원 형식: "ctrl+shift+1", "ctrl+shift+a" 등
func parseAndRegister(keyStr string) (*hotkey.Hotkey, error) {
	mods, key, err := parseHotkeyString(keyStr)
	if err != nil {
		return nil, err
	}

	hk := hotkey.New(mods, key)
	if err := hk.Register(); err != nil {
		return nil, fmt.Errorf("단축키 등록 실패: %w", err)
	}
	return hk, nil
}

// parseHotkeyString "ctrl+shift+1" 형태의 문자열을 modifier와 key로 분리한다
func parseHotkeyString(s string) ([]hotkey.Modifier, hotkey.Key, error) {
	parts := strings.Split(strings.ToLower(s), "+")
	if len(parts) < 2 {
		return nil, 0, fmt.Errorf("단축키 형식 오류: '%s' (예: ctrl+shift+1)", s)
	}

	var mods []hotkey.Modifier
	keyPart := parts[len(parts)-1]

	for _, part := range parts[:len(parts)-1] {
		switch part {
		case "ctrl", "control":
			mods = append(mods, hotkey.ModCtrl)
		case "shift":
			mods = append(mods, hotkey.ModShift)
		case "alt", "option":
			mods = append(mods, hotkey.ModOption)
		case "cmd", "command", "super":
			mods = append(mods, hotkey.ModCmd)
		default:
			return nil, 0, fmt.Errorf("알 수 없는 modifier: '%s'", part)
		}
	}

	key, err := parseKey(keyPart)
	if err != nil {
		return nil, 0, err
	}

	return mods, key, nil
}

// parseKey 키 문자열을 hotkey.Key로 변환한다
func parseKey(k string) (hotkey.Key, error) {
	keyMap := map[string]hotkey.Key{
		"0": hotkey.Key0, "1": hotkey.Key1, "2": hotkey.Key2,
		"3": hotkey.Key3, "4": hotkey.Key4, "5": hotkey.Key5,
		"6": hotkey.Key6, "7": hotkey.Key7, "8": hotkey.Key8, "9": hotkey.Key9,
		"a": hotkey.KeyA, "b": hotkey.KeyB, "c": hotkey.KeyC, "d": hotkey.KeyD,
		"e": hotkey.KeyE, "f": hotkey.KeyF, "g": hotkey.KeyG, "h": hotkey.KeyH,
		"i": hotkey.KeyI, "j": hotkey.KeyJ, "k": hotkey.KeyK, "l": hotkey.KeyL,
		"m": hotkey.KeyM, "n": hotkey.KeyN, "o": hotkey.KeyO, "p": hotkey.KeyP,
		"q": hotkey.KeyQ, "r": hotkey.KeyR, "s": hotkey.KeyS, "t": hotkey.KeyT,
		"u": hotkey.KeyU, "v": hotkey.KeyV, "w": hotkey.KeyW, "x": hotkey.KeyX,
		"y": hotkey.KeyY, "z": hotkey.KeyZ,
	}

	if key, ok := keyMap[k]; ok {
		return key, nil
	}
	return 0, fmt.Errorf("지원하지 않는 키: '%s'", k)
}
