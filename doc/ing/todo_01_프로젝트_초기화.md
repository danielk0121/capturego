# todo: 프로젝트 초기화

## 작업 목표
Go 모듈 초기화 및 공통 인프라 구성. 이후 모든 기능 개발의 기반이 되는 뼈대를 만든다.

## 세부 작업
- [ ] `app/go.mod` 초기화 (`module capturego`)
- [ ] 의존성 추가: `github.com/gin-gonic/gin`, `github.com/getlantern/systray`, `golang.design/x/hotkey`
- [ ] `app/utils/logger.go` 구현 (공통 로거)
- [ ] `app/config/` 설정 파일 스키마 정의 및 관리자 구현
- [ ] `app/main.go` 진입점 뼈대 작성 (각 컴포넌트 초기화 흐름)
