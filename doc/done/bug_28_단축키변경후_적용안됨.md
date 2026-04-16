## 개요
- 웹앱 > 단축키 변경 > 저장 > 새로 고침 > 저장 됨
- 하지만, 저장 이전에 설정한 단축키로 실행됨 > 버그

## 원인 분석
- `postConfig`에서 설정을 파일에 저장(`config.Update`)만 하고, 실행 중인 `HotkeyManager.Reload()`를 호출하지 않았음
- 설정 파일엔 새 단축키가 저장되지만, 앱 재시작 전까지는 이전 단축키가 그대로 유지됨

## 작업 결과
- `WebServer`에 `HotkeyReloader` 인터페이스를 주입받도록 변경 (`server/web_server.go`)
- `postConfig`를 `WebServer` 메서드로 변경하여 단축키 변경 감지 시 `Reload()` 즉시 호출
- `main.go`에서 `hotkeyMgr`을 먼저 생성한 후 `server.New(hotkeyMgr)`로 주입
- 단축키 변경 저장 즉시 새 단축키가 적용됨 (앱 재시작 불필요)
