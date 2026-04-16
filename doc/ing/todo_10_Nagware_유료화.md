# todo: 부분 유료화 플로우 — Nagware (F5)

## 작업 목표
기본 기능은 무료 제공하되, 누적 캡처 횟수 도달 시 macOS 네이티브 알림으로 후원을 안내한다. 라이선스 키 인증 완료 시 알림 로직을 영구 비활성화한다.

## 세부 작업
- [ ] `app/core/license.go`: 누적 캡처 횟수 카운트 및 영구 저장
- [ ] `app/core/license.go`: 임계값 도달 시 macOS 네이티브 알림 (`osascript`) 발송
- [ ] `app/core/license.go`: 라이선스 키 검증 로직
- [ ] 라이선스 키 인증 완료 시 알림 플래그 설정 파일에 기록 → 이후 알림 비활성화
- [ ] 설정 UI의 라이선스 키 입력 필드와 연동 (`POST /api/config` 핸들러)
- [ ] 알림 클릭 시 후원 페이지 URL 오픈

## 참고
- 알림: macOS `osascript` 또는 `NSUserNotification`
- 임계값 및 후원 URL은 `config/` 에서 관리
