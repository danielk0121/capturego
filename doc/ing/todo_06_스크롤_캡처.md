# todo: 스크롤 캡처 (F2)

## 작업 목표
긴 웹페이지나 소스 코드를 하나의 이미지로 캡처한다. 스크롤하며 연속 촬영 후 겹치는 픽셀 영역을 계산·절삭하여 하나의 세로 이미지로 병합(Stitching)한다.

## 세부 작업
- [ ] `app/core/scroll_capture.go`: 스크롤 캡처 모드 진입 및 영역 지정
- [ ] `app/core/scroll_capture.go`: 스크롤 감지 후 짧은 간격으로 연속 촬영 루프
- [ ] `app/core/stitch.go`: 이미지 간 겹치는 픽셀 영역 계산 알고리즘
- [ ] `app/core/stitch.go`: 절삭 후 세로 방향 이미지 병합(Stitching)
- [ ] 병합 결과 파일 저장 (듀얼 세이브와 동일한 저장 로직 재사용)
- [ ] 단위 테스트 작성 (Stitching 알고리즘 검증)

## 참고
- 이미지 처리: Go 표준 `image` 패키지
