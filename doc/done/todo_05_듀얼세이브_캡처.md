# todo: 듀얼 세이브 캡처 (F1)

## 작업 목표
글로벌 단축키 입력 시 macOS 네이티브 캡처 도구로 영역 지정 후, 파일 저장과 클립보드 복사를 고루틴으로 동시에 처리한다.

## 세부 작업
- [ ] `app/core/capture.go`: `screencapture` CLI 호출로 영역 지정 캡처 실행
- [ ] `app/core/capture.go`: 파일 저장 고루틴 — 사용자 지정 디렉토리에 이미지 저장
- [ ] `app/core/capture.go`: 클립보드 복사 고루틴 — `osascript` AppleScript로 클립보드 적재
- [ ] 두 고루틴 완료 대기 (`sync.WaitGroup`)
- [ ] 단위 테스트 작성

## 참고
- 캡처 엔진: macOS `screencapture` CLI
- 클립보드 제어: `osascript` (AppleScript)
- 파일명 포맷: `캡쳐고_YYYYMMDD_HHMMSS.png`
