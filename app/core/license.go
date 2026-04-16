package core

import (
	"capturego/config"
	"capturego/utils"
	"fmt"
	"os/exec"
	"strings"
)

const (
	// Nagware 팝업을 처음 표시할 누적 캡처 횟수
	nagwareThreshold = 50
	// 이후 반복 표시 주기
	nagwareRepeatInterval = 20

	// 후원 페이지 URL
	supportURL = "https://github.com/danielk0121/capturego"
)

// CheckNagware 누적 캡처 횟수를 확인해 필요 시 macOS 네이티브 알림을 표시한다
func CheckNagware() {
	cfg := config.Get()
	if cfg.NagwareDisabled || cfg.LicenseActivated {
		return
	}

	count := cfg.CaptureCount
	if count < nagwareThreshold {
		return
	}
	if (count-nagwareThreshold)%nagwareRepeatInterval != 0 {
		return
	}

	utils.Info("Nagware 알림 표시 (누적 캡처 %d회)", count)
	showNagwareNotification(count)
}

// ActivateLicense 라이선스 키를 검증하고 인증 완료 시 Nagware를 영구 비활성화한다
func ActivateLicense(key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("라이선스 키가 비어있습니다")
	}

	if !validateLicenseKey(key) {
		return fmt.Errorf("유효하지 않은 라이선스 키입니다")
	}

	cfg := config.Get()
	cfg.LicenseKey = key
	cfg.LicenseActivated = true
	cfg.NagwareDisabled = true

	if err := config.Update(cfg); err != nil {
		return fmt.Errorf("라이선스 저장 실패: %w", err)
	}

	utils.Info("라이선스 인증 완료: Nagware 비활성화")
	showActivationNotification()
	return nil
}

// validateLicenseKey 라이선스 키의 형식을 검증한다.
// 형식: XXXX-XXXX-XXXX-XXXX (영문 대문자 + 숫자)
func validateLicenseKey(key string) bool {
	parts := strings.Split(key, "-")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if len(part) != 4 {
			return false
		}
		for _, ch := range part {
			if !((ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')) {
				return false
			}
		}
	}
	return true
}

// showNagwareNotification macOS 네이티브 알림으로 후원 안내를 표시한다
func showNagwareNotification(count int) {
	script := fmt.Sprintf(`display notification "CaptureGo를 %d회 사용하셨습니다! 개발자에게 커피 한 잔 어떠세요? ☕" `+
		`with title "CaptureGo" subtitle "계속 무료로 사용하시거나, 후원으로 개발을 응원해주세요" `+
		`sound name "default"`,
		count,
	)
	if err := exec.Command("osascript", "-e", script).Run(); err != nil {
		utils.Warn("Nagware 알림 표시 실패: %v", err)
	}
}

// showActivationNotification 라이선스 인증 완료 알림을 표시한다
func showActivationNotification() {
	script := `display notification "라이선스가 인증되었습니다. 앞으로 팝업이 표시되지 않습니다." ` +
		`with title "CaptureGo" subtitle "감사합니다! ☕" sound name "default"`
	if err := exec.Command("osascript", "-e", script).Run(); err != nil {
		utils.Warn("인증 완료 알림 표시 실패: %v", err)
	}
}

// OpenSupportPage 후원 페이지를 기본 브라우저로 연다
func OpenSupportPage() {
	if err := exec.Command("open", supportURL).Start(); err != nil {
		utils.Error("후원 페이지 열기 실패: %v", err)
	}
}
