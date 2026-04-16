# todo: 디폴트 설정값 변경

## 작업 목표
앱 최초 실행 시 적용되는 기본 설정값을 변경한다.

## 변경 내용

| 항목 | 기존 | 변경 |
|------|------|------|
| 저장 폴더 | `~/Pictures/CaptureGo` | `~/screenshot_capturego` |
| 듀얼 세이브 단축키 | `ctrl+shift+1` | `cmd+shift+6` |
| 스크롤 캡쳐 단축키 | `ctrl+shift+2` | `cmd+shift+7` |

## 세부 작업

- [ ] `app/config/config_manager.go`
  - `defaultConfig()` 함수 내 값 변경
    - `SaveDirectory`: `filepath.Join(homeDir(), "screenshot_capturego")`
    - `HotkeyCapture`: `"cmd+shift+6"`
    - `HotkeyScroll`: `"cmd+shift+7"`
- [ ] `app/server/static/app.js`
  - i18n placeholder 값도 새 기본값 반영 (예시 문자열)
- [ ] `app/config/config_manager_test.go`
  - 기본값 검증 테스트 업데이트

## 주의사항
- 이미 설정 파일이 존재하는 사용자에게는 영향 없음 (최초 실행 시에만 적용)
- 기존 사용자가 기본값으로 리셋하고 싶을 경우 config.json 삭제 후 재실행

## 작업 결과

- `config_manager.go`: `defaultConfig()` 내 `SaveDirectory`, `HotkeyCapture`, `HotkeyScroll` 변경
- `app.js`: i18n ko/en 양쪽 placeholder 반영
- `config_manager_test.go`: 기본 단축키 검증 `"cmd+shift+6"` 으로 업데이트
- `go test ./config/...` 통과 확인
