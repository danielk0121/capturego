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
| 아이콘 생성 | Node.js + sharp (SVG → PNG/ICNS 자동 생성) |

---

## 프로젝트 구조

```
capturego/
├── app/
│   ├── main.go               # 진입점: 트레이, 웹서버, 단축키 초기화
│   ├── core/                 # 핵심 비즈니스 로직 (캡처, 클립보드, 스크롤 병합)
│   │   ├── capture.go        # 듀얼 세이브 캡처 및 클립보드 복사
│   │   ├── scroll_capture.go # 스크롤 캡처 및 이미지 병합
│   │   ├── stitch.go         # 이미지 스티칭 알고리즘
│   │   ├── hotkey.go         # 글로벌 단축키 등록
│   │   ├── license.go        # 라이선스 키 검증
│   │   └── permissions.go    # macOS 권한 확인
│   ├── server/               # Gin 라우터 및 REST API 핸들러
│   │   ├── web_server.go     # 서버 생성, 라우트, 미들웨어
│   │   └── static/           # go:embed 대상 정적 파일
│   │       ├── index.html    # 설정 페이지
│   │       ├── license.html  # 오픈소스 라이선스 페이지
│   │       ├── style.css     # 공통 CSS (다크모드 포함)
│   │       ├── app.js        # 설정 페이지 JS (i18n 포함)
│   │       └── favicon.png   # 웹앱 파비콘 (32px)
│   ├── ui/                   # 시스템 트레이 자원
│   │   ├── tray_icon.png     # 트레이 아이콘 (22px)
│   │   └── tray_icon@2x.png  # 트레이 아이콘 Retina (44px)
│   ├── config/               # 설정 파일 및 관리자
│   └── utils/                # 로거 등 공통 유틸리티
├── build-work/
│   ├── build.sh              # 빌드 스크립트 (아이콘 생성 → Go 빌드 → 패키징)
│   ├── AppIcon.icns          # 앱 번들 아이콘 (.app 독바/Finder)
│   ├── Info.plist            # macOS 앱 번들 메타데이터
│   └── gen_icon/             # 아이콘 자동 생성 스크립트
│       ├── capturego_icon.svg # 마스터 SVG 아이콘
│       ├── gen_icon.js        # SVG → PNG/ICNS 변환 스크립트
│       └── package.json       # Node.js 의존성 (sharp)
├── doc/                      # 작업 문서 (ing/hold/done/discard)
└── ref/                      # 참고 자료
```

---

## 주요 진입점

| 파일 | 역할 |
|------|------|
| `app/main.go` | 앱 시작, 트레이·웹서버·단축키 초기화 및 생명주기 관리 |
| `app/server/web_server.go` | Gin 서버 생성, 라우트 등록, 노캐시 미들웨어 |
| `app/core/capture.go` | 듀얼 세이브 캡처 실행 및 파일/클립보드 병렬 처리 |
| `app/core/scroll_capture.go` | 스크롤 캡처 루프 및 이미지 스티칭 호출 |
| `app/config/config_manager.go` | JSON 설정 파일 읽기/쓰기, 기본값 관리 |
| `build-work/gen_icon/gen_icon.js` | SVG 마스터에서 ICNS·트레이·파비콘 일괄 생성 |

---

## 실행 중 프로세스 구성

| 고루틴 | 역할 |
|--------|------|
| Main (메인 스레드) | 시스템 트레이 실행 — macOS 요구사항으로 메인 스레드 점유 |
| HotkeyListener | OS 전역 단축키 감지 및 캡처 이벤트 트리거 |
| WebServer | Gin 기반 로컬 HTTP 서버, 설정 UI 및 API 제공 |

---

## 웹 서버 엔드포인트

| 경로 | 메서드 | 설명 |
|------|--------|------|
| `/` | GET | 설정 UI 페이지 (`static/index.html`) |
| `/license` | GET | 오픈소스 라이선스 페이지 (`static/license.html`) |
| `/static/*` | GET | 정적 파일 서빙 (CSS, JS, favicon 등) |
| `/api/config` | GET | 현재 설정값 JSON 반환 |
| `/api/config` | POST | 설정값 갱신 (저장 경로, 단축키, 라이선스 키) |

> 모든 응답에 `Cache-Control: no-cache, no-store, must-revalidate` 헤더 적용.

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
  └─> OS 기본 브라우저 → localhost:18420 (Gin 서버)
        └─> 설정 변경 (저장 경로 / 단축키 / 라이선스)
              └─> POST /api/config → config 파일 갱신
```

**아이콘 생성 흐름**
```
build-work/build.sh 실행
  └─> [0/3] gen_icon.js (Node.js + sharp)
        ├─> SVG → iconset 10종 → iconutil → AppIcon.icns
        ├─> SVG → tray_icon.png (22px) + tray_icon@2x.png (44px)
        └─> SVG → favicon.png (32px)
  └─> [1/3] go build → capturego 바이너리
  └─> [2/3] .app 번들 패키징
  └─> [3/3] 배포 패키지 생성
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

**전체 앱 빌드 (.app 번들 + 아이콘 생성)**

```bash
# Node.js 18+ 및 Xcode Command Line Tools 필요
cd capturego/build-work
bash build.sh
```

> **타겟 환경:** macOS (Apple Silicon 아키텍처 기반 시스템 최적화)

---

## 필수 권한 설정 (macOS)

CaptureGo는 다음 두 가지 권한이 필요합니다.

1. **화면 기록 권한**: 시스템 환경설정 → 개인 정보 보호 및 보안 → 화면 기록 → CaptureGo 허용
2. **손쉬운 사용 권한**: 시스템 환경설정 → 개인 정보 보호 및 보안 → 손쉬운 사용 → CaptureGo 허용 (글로벌 단축키 감지에 필요)

> 권한 부여 후 앱을 재시작해야 정상 동작합니다.
