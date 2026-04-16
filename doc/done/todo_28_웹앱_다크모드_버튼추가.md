## 개요
- 웹앱 > 다크모드/라이트모드 우상단에 버튼 추가
- 기본값: 라이트모드
- 저장버튼 없이, 누르자 마자 로컬 설정 파일에 저장하여 적용
- 보통 웹앱들이 사용하는 로컬스토리지 사용하는 방법이 아님
- config json 파일을 사용

## 작업 결과

- `config/config_manager.go`: `DarkMode bool` 필드 추가 (`json:"dark_mode"`)
- `server/web_server.go`: `POST /api/darkmode` 엔드포인트 추가, `/api/config` 응답에 `dark_mode` 포함
- `static/style.css`: `body.dark`, `body.light` 강제 테마 클래스 추가, `.theme-btn` 스타일 추가
- `static/index.html`: 우상단 언어 버튼 옆 `🌙/☀️` 토글 버튼 추가
- `static/app.js`: `applyTheme()`, `toggleTheme()` 함수 추가, `loadConfig()` 시 `dark_mode` 읽어 적용
