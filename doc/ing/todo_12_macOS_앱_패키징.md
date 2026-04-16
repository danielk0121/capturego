# todo: macOS 앱 패키징 (.app 번들)

## 작업 목표
Go 바이너리를 macOS 표준 `.app` 번들로 패키징하여 일반 사용자가 더블클릭으로 실행할 수 있게 한다.

## 세부 작업
- [ ] `Info.plist` 작성
  - `LSUIElement = true` 설정 (Dock 미표시, 메뉴바 전용)
  - 앱 이름, 번들 ID, 버전 정보
  - 필요 권한 선언 (화면 기록, 손쉬운 사용)
- [ ] 앱 아이콘 제작 및 `AppIcon.icns` 생성
- [ ] `build-work/build.sh`: 빌드 → `.app` 번들 생성 스크립트
- [ ] Apple Silicon (`arm64`) 타겟 빌드 확인
- [ ] 빌드 결과물 `build-work/` 에 출력

## 참고
- 타겟: macOS Apple Silicon (arm64)
- Dock 미표시: `Info.plist`의 `LSUIElement` 키로 제어
