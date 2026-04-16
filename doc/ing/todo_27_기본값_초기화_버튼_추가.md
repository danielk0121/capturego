# todo: 시스템 기본 설정값으로 초기화 버튼 추가

## 작업 목표
설정 UI에 "기본값으로 초기화" 버튼을 추가한다.
`default-config.json`을 읽기 전용 기본값 소스로 도입하고,
초기화 버튼 클릭 시 해당 파일의 값으로 `config.json`을 덮어쓴다.

## 현황
- 기본값이 `config_manager.go`의 `defaultConfig()` 함수에 하드코딩되어 있음
- `default-config.json` 파일 없음
- 설정 UI에 초기화 버튼 없음

## 변경 방향

### 1. `app/config/default-config.json` 신규 생성
- `defaultConfig()` 함수의 하드코딩 값과 동일한 내용으로 작성
- 읽기 전용 참조 파일로 사용 (앱이 직접 수정하지 않음)
- `go:embed`로 바이너리에 내장

### 2. `app/config/config_manager.go` 수정
- `default-config.json`을 embed로 로드
- `defaultConfig()` 함수가 embed된 JSON을 파싱해서 반환하도록 변경
- `ResetToDefault()` 함수 신규 추가: default-config.json 값으로 config.json 덮어쓰기

### 3. 설정 UI (웹) 수정
- 설정 페이지 하단에 "기본값으로 초기화" 버튼 추가
- 클릭 시 확인 다이얼로그 → POST /api/config/reset 호출

### 4. `app/server/` API 엔드포인트 추가
- `POST /api/config/reset`: `config.ResetToDefault()` 호출 후 200 반환

## 세부 작업

- [ ] `app/config/default-config.json` 생성
- [ ] `app/config/config_manager.go`
  - `//go:embed default-config.json` 추가
  - `defaultConfig()` → embed JSON 파싱으로 변경
  - `ResetToDefault()` 함수 추가
- [ ] `app/server/` 라우터에 `POST /api/config/reset` 엔드포인트 추가
- [ ] 설정 UI HTML/JS에 초기화 버튼 추가
- [ ] e2e 테스트: 초기화 버튼 클릭 → 설정값이 기본값으로 복원되는지 확인

## 작업 결과
