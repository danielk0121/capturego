# todo: 앱 데이터 경로 통일

## 작업 목표
- 설정 파일(`config.json`)과 로그 파일(`capturego.log`)이 저장되는 기본 디렉토리를 `~/Library/Application Support/CaptureGo`로 통일한다.
- 기존 로그 경로(`~/Library/Logs/CaptureGo`) 사용을 중단하고, 설정 디렉토리 내 `logs` 서브 디렉토리를 사용하도록 변경한다.

## 세부 작업
- [x] `app/utils/logger.go`: 로그 디렉토리를 `~/Library/Application Support/CaptureGo/logs`로 변경
- [ ] `app/config/config_manager.go`: 로그와 설정을 위한 베이스 경로 상수화 (선택 사항 - 현 단계에서는 각각 유지하되 경로만 맞춤)
- [x] `InitLogger` 수정: 로그 파일 생성 전 디렉토리 생성 로직 확인

## 작업 결과
- 로그 파일 저장 경로를 설정 파일과 동일한 `Application Support/CaptureGo/logs` 하위로 변경 완료.
