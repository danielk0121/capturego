# todo: footer 정보 변경

## 작업 목표
웹앱 footer 텍스트를 변경하고, 오픈소스 라이선스 링크 텍스트를 영어로 통일한다.

## 변경 내용

### footer 텍스트
- 기존: `CaptureGo — 개발자: maedk10@gmail.com`
- 변경: `github.com/danielk0121 v-yyyyMMdd-HHmm-kst`
  - `v-yyyyMMdd-HHmm-kst` 는 `/api/buildtime` 에서 동적으로 읽어 렌더링
  - 빌드타임 정보가 없는 경우 `v-dev` 로 표시

### 오픈소스 라이선스 링크
- 기존 i18n key `footer_oss_link`: `ko="오픈소스 라이선스"` / `en="Open Source Licenses"`
- 변경: 언어 무관하게 항상 영어 `"Open Source Licenses"` 고정

## 세부 작업

- [ ] `app/server/static/index.html`
  - footer 내용을 `github.com/danielk0121 <span id="buildtime">v-dev</span>` 형태로 변경
  - 오픈소스 링크 텍스트: `data-i18n` 제거하고 고정 영어 텍스트 사용
- [ ] `app/server/static/app.js`
  - 초기화 시 `GET /api/buildtime` fetch → `#buildtime` span에 렌더링
  - 실패 시 `v-dev` 유지
  - i18n에서 `footer_dev`, `footer_oss_link` ko 값 영어로 통일 또는 제거
- [ ] `app/server/web_server.go`
  - `/api/buildtime` 엔드포인트 추가
  - `app/server/static/buildtime.txt` 파일을 embed FS에서 읽어 반환
  - 파일 없으면 `"v-dev"` 반환
- [ ] `app/core/buildtime.go` (신규)
  - buildtime.txt 읽기 유틸 함수
- [ ] `build-work/build.sh`
  - 빌드 시 `app/server/static/buildtime.txt` 에 `v-yyyyMMdd-HHmm-kst` 값 기록

## 작업 결과

- `index.html`: footer → `github.com/danielk0121 — <span id="buildtime">v-dev</span>` + 고정 영어 "Open Source Licenses" (todo_19에서 완료)
- `app.js`: `loadBuildtime()` 함수로 `/api/buildtime` fetch → `#buildtime` 렌더링 (todo_19에서 완료)
- `web_server.go`: `/api/buildtime` 엔드포인트 추가 (todo_19에서 완료)
- `build.sh`: `[1/5]` 단계에서 KST 기준 `v-yyyyMMdd-HHmm-kst` 형식으로 `buildtime.txt` 자동 기록

## 추가 작업
- [ ] 오픈소스 라이선스 페이지에도 동일하게 footer 적용

## 추가 작업 결과
