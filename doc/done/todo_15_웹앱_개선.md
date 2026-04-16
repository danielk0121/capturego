# todo: 웹앱 개선 (다국어·다크모드·푸터·라이선스·노캐시)

## 작업 목표
설정 웹앱 UI를 다음 5가지 항목으로 개선한다.

## 세부 작업

- [x] **다국어 지원** (한국어 / 영어)
  - 언어 토글 버튼 (🇰🇷 / 🇺🇸) 상단 배치
  - 모든 레이블·플레이스홀더·버튼 문자열을 i18n 딕셔너리로 관리
  - localStorage에 언어 선택값 저장 (새로고침 유지)

- [x] **다크모드 지원**
  - `@media (prefers-color-scheme: dark)` CSS 미디어 쿼리로 자동 전환
  - 별도 토글 불필요 (시스템 설정 따름)

- [x] **웹앱 하단 개발자 정보 (Footer)**
  - 앱 이름, 개발자 이메일 표시
  - 스타일: 작은 폰트, 중앙 정렬, 회색 계열

- [x] **라이선스 정보 페이지**
  - `/license` 경로 추가 (gin 라우트)
  - 오픈소스 라이브러리 목록 표시 (Go 모듈 기반)
  - 설정 페이지 푸터에 링크 추가

- [x] **노캐시 적용**
  - 모든 응답에 `Cache-Control: no-cache, no-store, must-revalidate` 헤더 추가
  - gin 미들웨어로 전역 적용

## 참고
- 웹앱 HTML: `app/server/html.go` (인라인 문자열)
- 웹 서버: `app/server/web_server.go`
- 개발자 이메일: maedk10@gmail.com
