# todo: 웹앱 HTML/CSS/JS 별도 파일 관리

## 작업 목표
`app/server/html.go`의 인라인 문자열로 관리되던 HTML/CSS/JS를
`app/server/static/` 하위 별도 파일로 분리하여 유지보수성을 높인다.

## 현황
- `app/server/html.go`: `indexHTML()`, `licenseHTML()` — 인라인 문자열로 전체 HTML 보관
- `app/server/web_server.go`: `//go:embed static` 디렉티브로 `static/` 폴더 embed 중

## 세부 작업

- [x] `app/server/static/index.html` 생성 (설정 페이지)
- [x] `app/server/static/license.html` 생성 (라이선스 페이지)
- [x] `app/server/static/style.css` 생성 (공통 CSS)
- [x] `app/server/static/app.js` 생성 (설정 페이지 JS)
- [x] `app/server/web_server.go` 라우트 변경: `serveIndex`, `serveLicense`를 embed.FS에서 직접 서빙
- [x] `app/server/html.go` 삭제 (인라인 함수 제거)

## 파일 구조 (목표)
```
app/server/
├── static/
│   ├── favicon.png
│   ├── style.css       ← 공통 CSS (다크모드 포함)
│   ├── app.js          ← 설정 페이지 JS (i18n 포함)
│   ├── index.html      ← 설정 페이지
│   └── license.html    ← 라이선스 페이지
├── web_server.go
└── web_server_test.go
```

## 참고
- `//go:embed static` 디렉티브는 유지 (`web_server.go`)
- HTML에서 CSS/JS는 `/static/style.css`, `/static/app.js` 경로로 참조
- 노캐시 미들웨어는 유지 (정적 파일 포함)
