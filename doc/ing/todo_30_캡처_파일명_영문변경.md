# todo: 캡처 파일명 영문 변경

## 작업 목표
- 스크린샷 저장 시 파일명 접두어를 한글("캡쳐고_")에서 영문("capturego_")으로 변경한다.
- 일반 캡처와 스크롤 캡처 모두 적용한다.

## 세부 작업
- [ ] `app/core/capture.go`: 파일명 접두어 변경 (`캡쳐고_` -> `capturego_`)
- [ ] `app/core/scroll_capture.go`: 파일명 접두어 변경 (`캡쳐고_스크롤_` -> `capturego_scroll_`)
- [ ] `app/core/capture_test.go`: 테스트 코드의 파일명 검증 로직 수정

## 작업 결과
