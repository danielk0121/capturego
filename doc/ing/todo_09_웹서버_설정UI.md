# todo: 로컬 웹서버 및 설정 UI (F3)

## 작업 목표
Gin 기반 로컬 HTTP 서버를 실행하고, 브라우저에서 접근 가능한 설정 UI를 제공한다. `go:embed`로 웹 자원을 바이너리에 내장한다.

## 세부 작업
- [ ] `app/server/web_server.go`: Gin 서버 초기화 및 `localhost` 포트 바인딩
- [ ] `app/server/handlers.go`: REST API 핸들러 구현
  - `GET /` — 설정 UI HTML 서빙
  - `GET /api/config` — 현재 설정값 반환
  - `POST /api/config` — 설정값 저장
- [ ] `app/ui/assets.go`: `go:embed`로 `static/`, `templates/` 바이너리 내장
- [ ] `app/ui/templates/index.html`: 설정 화면 HTML
- [ ] `app/ui/static/app.js`: 설정 저장/불러오기 API 연동
- [ ] `app/ui/static/style.css`: 기본 스타일
- [ ] 설정 항목: 캡처 저장 경로 / 글로벌 단축키 / 라이선스 키
