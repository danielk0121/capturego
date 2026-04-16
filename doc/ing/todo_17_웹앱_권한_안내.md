# todo: 웹앱 권한 상태 표시 및 초기 안내

## 작업 목표
1. 웹앱에 macOS 권한(화면 기록, 손쉬운 사용) 승인 여부를 실시간으로 표시한다.
2. 웹앱 첫 실행 시(또는 권한 미부여 시) 권한 허용 안내 가이드 모달을 표시한다.

## 세부 작업

- [x] `app/core/permissions.go`: `QueryPermissions()` 함수 추가 — 화면 기록, 손쉬운 사용 상태 반환
- [x] `app/server/web_server.go`: `/api/permissions` GET 엔드포인트 추가
- [x] `app/server/static/index.html`: 권한 상태 카드 추가, 첫 실행 안내 모달 추가
- [x] `app/server/static/style.css`: 권한 상태 배지 스타일, 모달 스타일 추가
- [x] `app/server/static/app.js`: i18n 문자열 추가, 권한 조회/표시 로직, 모달 표시 로직 추가

## 구현 상세

### 권한 조회 (`permissions.go`)
- `QueryScreenRecording() bool`: screencapture 테스트로 화면 기록 권한 확인
- `QueryAccessibility() bool`: CGo `AXIsProcessTrusted()` 호출로 손쉬운 사용 권한 확인
- `QueryPermissions() (screen bool, accessibility bool)`: 두 권한 상태 반환

### API (`web_server.go`)
```
GET /api/permissions
→ { "screen_recording": bool, "accessibility": bool }
```

### UI (index.html)
- 권한 상태 카드: 각 권한 항목에 ✓(녹색) / ✗(빨간) 배지 + "설정 열기" 버튼
- 첫 실행 모달: localStorage `cg_welcomed` 없고 권한 미부여 시 표시, 단계별 안내, 닫기 시 `cg_welcomed=1` 저장

## 참고
- 화면 기록 권한 확인: screencapture -x -m tmpFile 실행 후 파일 존재 여부
- 손쉬운 사용 권한 확인: osascript `tell application "System Events" to return UI elements enabled`
- 설정 열기 URL: `x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture`
- 설정 열기 URL: `x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility`

## 작업 결과

### 구현 내용
- `QueryPermissions() PermissionStatus` 추가: 화면 기록(screencapture 테스트), 손쉬운 사용(osascript UI elements enabled) 각각 실시간 조회
- `GET /api/permissions` 엔드포인트: `{"screen_recording": bool, "accessibility": bool}` 반환
- 웹앱 권한 카드: 각 권한 항목에 허용/미허용 색상 배지 + 미허용 시 "설정 열기" 링크(URL scheme)
- 첫 실행 모달: 권한 미부여 상태이고 `localStorage.cg_welcomed` 미설정 시 자동 표시; 단계별 안내(클릭 가능 링크 포함); 닫기 시 `cg_welcomed=1` 저장
- 언어 전환 시 권한 패널도 해당 언어로 재렌더링
- i18n: ko/en 모두 지원
