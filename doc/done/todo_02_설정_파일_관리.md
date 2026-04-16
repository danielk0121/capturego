# todo: 설정 파일 관리

## 작업 목표
앱의 모든 설정값(저장 경로, 단축키, 라이선스 상태, 누적 캡처 횟수 등)을 로컬 파일로 영구 저장하고 런타임에 로드한다.

## 세부 작업
- [ ] `app/config/default_config.json`: 기본 설정값 정의
- [ ] `app/config/config_manager.go`: 설정 파일 로드 / 저장 / 갱신 함수 구현
- [ ] 설정 파일 저장 경로: `~/Library/Application Support/CaptureGo/config.json`
- [ ] 앱 최초 실행 시 기본 설정 파일 자동 생성
- [ ] 설정 스키마 정의 (Go 구조체)
  - 캡처 저장 디렉토리 경로
  - 글로벌 단축키 (듀얼 세이브 / 스크롤 캡처)
  - 누적 캡처 횟수
  - 라이선스 인증 여부
  - Nagware 팝업 비활성화 여부
- [ ] 단위 테스트 작성
