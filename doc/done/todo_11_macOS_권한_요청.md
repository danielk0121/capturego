# todo: macOS 권한 요청 처리

## 작업 목표
앱 최초 실행 시 필요한 macOS 권한이 없을 경우, 사용자에게 안내하고 시스템 환경설정으로 유도한다.

## 세부 작업
- [ ] 화면 기록 권한 확인 로직 (CGo 또는 `screencapture` 테스트 호출)
- [ ] 손쉬운 사용 권한 확인 로직 (단축키 등록 성공 여부로 판단)
- [ ] 권한 미부여 시 macOS 네이티브 알림으로 안내 메시지 표시
- [ ] 알림 클릭 시 시스템 환경설정 해당 패널로 직접 이동
- [ ] 권한 부여 후 앱 재시작 안내

## 참고
- 화면 기록: `Privacy & Security > Screen Recording`
- 손쉬운 사용: `Privacy & Security > Accessibility`
