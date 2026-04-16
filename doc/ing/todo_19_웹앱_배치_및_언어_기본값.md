# todo: 웹앱 배치 순서 변경 및 언어 기본값 EN

## 작업 목표
설정 웹앱 카드 배치 순서를 변경하고, 언어 기본값을 영어(EN)로 설정한다.
언어 전환 버튼을 EN / 한국어 2개로 분리하여 우상단 배치하고, 제목의 ⚡ 이모지를 제거한다.

## 세부 작업

- [ ] `app/server/static/index.html`
  - `<h1>` 에서 `⚡` 제거 → `CaptureGo Settings` 형태로 변경 (`data-i18n` 유지)
  - 카드 순서 변경: macOS 권한 → 저장 경로 → 단축키 → 현황 → 라이선스
  - 언어 버튼을 2개로 분리: `<button id="btnEN">EN</button>` / `<button id="btnKO">한국어</button>`
- [ ] `app/server/static/app.js`
  - 언어 기본값 `'ko'` → `'en'` 으로 변경
  - `toggleLang()` 제거, `setLang(targetLang)` 함수로 교체
  - EN/한국어 버튼 각각 active 스타일 적용 (현재 선택된 언어 버튼 강조)
- [ ] `app/server/static/style.css`
  - 언어 버튼 2개 나란히 배치 스타일 추가 (`.lang-btn-group`)
  - active 버튼 강조 스타일 (`.lang-btn.active`)

## 작업 결과

- `index.html`: ⚡ 제거, 카드 순서 변경(권한→저장경로→단축키→현황→라이선스), 언어 버튼 2개(EN/한국어) 분리, footer 변경
- `app.js`: 언어 기본값 `'en'`, `toggleLang()` → `setLang(targetLang)` 교체, `requestAnimationFrame` 일괄 DOM 업데이트, 권한 데이터 캐시(`cachedPermData`), `renderPermissions()` 분리, buildtime fetch 추가
- `style.css`: `.lang-btn-group` 추가, `.lang-btn.active` 강조 스타일 추가
- `web_server.go`: `/api/buildtime` 엔드포인트 추가, `strings` import 추가
- `static/buildtime.txt`: 기본값 `v-dev` 파일 생성 (embed 대상)
