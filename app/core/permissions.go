package core

import (
	"capturego/utils"
	"os"
	"os/exec"
)

// PermissionStatus 권한 상태를 담는 구조체
type PermissionStatus struct {
	ScreenRecording bool `json:"screen_recording"`
	Accessibility   bool `json:"accessibility"`
}

// QueryPermissions 현재 macOS 권한 상태를 조회한다
func QueryPermissions() PermissionStatus {
	return PermissionStatus{
		ScreenRecording: queryScreenRecording(),
		Accessibility:   queryAccessibility(),
	}
}

// queryScreenRecording 화면 기록 권한 상태를 반환한다
func queryScreenRecording() bool {
	tmpFile := os.TempDir() + "/capturego_perm_query.png"
	defer os.Remove(tmpFile)

	if err := exec.Command("screencapture", "-x", "-m", tmpFile).Run(); err != nil {
		return false
	}
	_, err := os.Stat(tmpFile)
	return err == nil
}

// queryAccessibility 손쉬운 사용 권한 상태를 반환한다 (osascript로 간접 확인)
func queryAccessibility() bool {
	// AXIsProcessTrusted를 osascript로 간접 조회
	script := `tell application "System Events" to return UI elements enabled`
	out, err := exec.Command("osascript", "-e", script).Output()
	if err != nil {
		return false
	}
	result := string(out)
	return len(result) > 0 && result[0] == 't' // "true\n"
}

// CheckPermissions 필요한 macOS 권한을 확인하고, 미부여 시 안내 알림을 표시한다
func CheckPermissions() {
	checkScreenRecording()
	checkAccessibility()
}

// checkScreenRecording 화면 기록 권한을 확인한다.
// screencapture 테스트 호출로 권한 여부를 판단한다.
func checkScreenRecording() {
	// 임시 파일로 screencapture 테스트 실행
	tmpFile := os.TempDir() + "/capturego_permission_test.png"
	defer os.Remove(tmpFile)

	// -x: 소리 없음, -m: 메뉴바 캡처 (사용자 인터랙션 불필요)
	err := exec.Command("screencapture", "-x", "-m", tmpFile).Run()
	if err != nil {
		utils.Warn("화면 기록 권한 없음 (screencapture 실패): %v", err)
		showPermissionNotification(
			"화면 기록 권한이 필요합니다",
			"시스템 환경설정 → 개인 정보 보호 및 보안 → 화면 기록에서 CaptureGo를 허용해 주세요.",
			"x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture",
		)
		return
	}

	// 파일이 생성되었는지 확인
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		utils.Warn("화면 기록 권한 없음 (캡처 파일 미생성)")
		showPermissionNotification(
			"화면 기록 권한이 필요합니다",
			"시스템 환경설정 → 개인 정보 보호 및 보안 → 화면 기록에서 CaptureGo를 허용해 주세요.",
			"x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture",
		)
		return
	}

	utils.Info("화면 기록 권한: 확인됨")
}

// checkAccessibility 손쉬운 사용 권한을 확인한다.
// 글로벌 단축키 등록 가능 여부로 판단한다.
func checkAccessibility() {
	// 손쉬운 사용 권한은 단축키 등록 실패로만 판단할 수 있으므로
	// 여기서는 안내만 제공하고, 실제 판단은 HotkeyManager.Start() 결과로 처리한다
	utils.Info("손쉬운 사용 권한: 단축키 등록 시 자동 확인됩니다")
}

// NotifyAccessibilityRequired 손쉬운 사용 권한이 필요할 때 알림을 표시한다.
// HotkeyManager.Start() 실패 시 호출된다.
func NotifyAccessibilityRequired() {
	showPermissionNotification(
		"손쉬운 사용 권한이 필요합니다",
		"시스템 환경설정 → 개인 정보 보호 및 보안 → 손쉬운 사용에서 CaptureGo를 허용한 후 앱을 재시작하세요.",
		"x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility",
	)
}

// showPermissionNotification macOS 알림과 함께 설정 패널로 이동하는 버튼을 제공한다
func showPermissionNotification(title, message, settingsURL string) {
	script := `set btn to button returned of (display alert "` + title + `" ` +
		`message "` + message + `" ` +
		`buttons {"나중에", "설정 열기"} default button "설정 열기" ` +
		`as critical)
if btn is "설정 열기" then
  open location "` + settingsURL + `"
end if`

	if err := exec.Command("osascript", "-e", script).Run(); err != nil {
		utils.Warn("권한 알림 표시 실패: %v", err)
	}
}
