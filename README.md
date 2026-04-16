# CaptureGo (캡쳐고)

macOS 환경에서 스크린샷 캡처 시 **파일 저장과 클립보드 복사를 동시에 수행**하는 초경량 백그라운드 유틸리티.

> 순수 Go 언어와 네이티브 OS 기능만을 활용해 극강의 실행 속도와 안정성을 확보한다. 무거운 프론트엔드 프레임워크 및 불필요한 서드파티 라이브러리를 철저히 배제한다.

---

## 핵심 기능

| 기능 | 설명 |
|------|------|
| **듀얼 세이브 캡처** | 글로벌 단축키 → 영역 지정 → 파일 저장 + 클립보드 복사 동시 처리 |
| **스크롤 캡처** | 긴 페이지 연속 촬영 후 자체 알고리즘으로 이미지 병합(Stitching) |
| **설정 UI** | 트레이 메뉴 → `localhost` 브라우저 UI에서 저장 경로 및 단축키 설정 |
| **시스템 트레이** | Dock 미표시, 상단 메뉴바 상주 |
| **부분 유료화** | 기본 기능 무료 제공 / 누적 캡처 시 후원 안내 팝업 / 라이선스 키로 영구 해제 |

---

## 기술 스택

| 구분 | 기술 |
|------|------|
| Core & Backend | Go |
| Local Web Server | Gin Framework |
| Frontend (설정 UI) | Vanilla JS / CSS / HTML5 (`go:embed` 내장) |
| 시스템 트레이 | getlantern/systray |
| 글로벌 단축키 | golang.design/x/hotkey |
| 캡처 엔진 | macOS `screencapture` CLI |
| 클립보드 제어 | macOS `osascript` (AppleScript) |
| 이미지 처리 | Go 표준 `image` 패키지 |

---

## 프로젝트 구조

```
capturego/
├── app/
│   ├── main.go               # 진입점: 트레이, 웹서버, 단축키 초기화
│   ├── core/                 # 핵심 비즈니스 로직 (캡처, 클립보드, 스크롤 병합)
│   ├── server/               # Gin 라우터 및 REST API 핸들러
│   ├── ui/                   # 시스템 트레이 + 설정 웹 자원 (go:embed)
│   │   ├── static/           # JS, CSS
│   │   └── templates/        # HTML 템플릿
│   ├── config/               # 설정 파일 및 관리자
│   └── utils/                # 로거 등 공통 유틸리티
├── doc/                      # 작업 문서 (ing/hold/done/discard)
└── ref/                      # 참고 자료
```

---

## 실행 중 프로세스 구성

| 고루틴 | 역할 |
|--------|------|
| Main (메인 스레드) | 시스템 트레이 실행 — macOS 요구사항으로 메인 스레드 점유 |
| HotkeyListener | OS 전역 단축키 감지 및 캡처 이벤트 트리거 |
| WebServer | Gin 기반 로컬 HTTP 서버, 설정 UI 및 API 제공 |

---

## 데이터 플로우

**단축키 캡처 흐름**
```
글로벌 단축키 입력
  └─> HotkeyListener 감지
        └─> screencapture CLI 실행 (영역 지정 모드)
              └─> 영역 지정 완료
                    ├─> [Goroutine A] 파일 저장 (지정 디렉토리)
                    └─> [Goroutine B] 클립보드 복사 (osascript)
```

**설정 변경 흐름**
```
트레이 메뉴 → '설정 열기'
  └─> OS 기본 브라우저 → localhost (Gin 서버)
        └─> 설정 변경 (저장 경로 / 단축키 / 라이선스)
              └─> REST API → config 파일 갱신
```

---

## 설치 및 실행 (Quick Guide)

```bash
# 저장소 클론
git clone https://github.com/danielk0121/capturego.git
cd capturego/app

# 의존성 설치
go mod download

# 빌드
go build -o capturego .

# 실행
./capturego
```

> **타겟 환경:** macOS (Apple Silicon 아키텍처 기반 시스템 최적화)

---

## 필수 권한 설정 (macOS)

CaptureGo는 다음 두 가지 권한이 필요합니다.

1. **화면 기록 권한**: 시스템 환경설정 → 개인 정보 보호 및 보안 → 화면 기록 → CaptureGo 허용
2. **손쉬운 사용 권한**: 시스템 환경설정 → 개인 정보 보호 및 보안 → 손쉬운 사용 → CaptureGo 허용 (글로벌 단축키 감지에 필요)

> 권한 부여 후 앱을 재시작해야 정상 동작합니다.
