package core

import (
	"testing"

	"golang.design/x/hotkey"
)

func TestParseHotkeyString_정상_입력(t *testing.T) {
	cases := []struct {
		input    string
		wantMods int
		wantKey  hotkey.Key
	}{
		{"ctrl+shift+1", 2, hotkey.Key1},
		{"ctrl+shift+2", 2, hotkey.Key2},
		{"ctrl+a", 1, hotkey.KeyA},
		{"cmd+shift+s", 2, hotkey.KeyS},
	}

	for _, tc := range cases {
		mods, key, err := parseHotkeyString(tc.input)
		if err != nil {
			t.Errorf("'%s': 파싱 실패: %v", tc.input, err)
			continue
		}
		if len(mods) != tc.wantMods {
			t.Errorf("'%s': modifier 수가 %d여야 하는데 %d입니다", tc.input, tc.wantMods, len(mods))
		}
		if key != tc.wantKey {
			t.Errorf("'%s': key가 %v여야 하는데 %v입니다", tc.input, tc.wantKey, key)
		}
	}
}

func TestParseHotkeyString_잘못된_입력(t *testing.T) {
	cases := []string{
		"ctrl",         // key 없음
		"ctrl+shift+!", // 지원하지 않는 키
		"win+1",        // 알 수 없는 modifier
	}

	for _, input := range cases {
		_, _, err := parseHotkeyString(input)
		if err == nil {
			t.Errorf("'%s': 오류가 발생해야 하는데 nil이 반환되었습니다", input)
		}
	}
}
