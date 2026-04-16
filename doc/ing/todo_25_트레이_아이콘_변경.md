# todo: 트레이 아이콘 변경 (독바 디자인 기반 흑백 버전)

## 작업 목표
현재 "캡쳐고" 텍스트 아이콘 대신, 독바 아이콘(AppIcon)과 동일한 디자인의 흰색 위주 흑백 아이콘을 트레이에 사용한다.

## 현황
- 현재 트레이 아이콘: `app/ui/tray_icon.png` (22px), `app/ui/tray_icon@2x.png` (44px) — "캡쳐고" 텍스트 기반
- 독바 아이콘: `build-work/gen_icon/capturego_icon.svg` (컬러 SVG)

## 변경 방향
- 독바 SVG를 기반으로 **흰색 위주 흑백 버전** SVG를 신규 생성
  - 배경: 투명
  - 아이콘 도형: 흰색(#FFFFFF) + 회색 그림자/테두리 최소화
  - macOS 메뉴바 다크/라이트 모드 모두 대응 (흰색이면 다크 모드에서 잘 보임)
- `gen_icon.js` 에 트레이 흑백 아이콘 생성 단계 추가

## 세부 작업

- [ ] `build-work/gen_icon/tray_icon_mono.svg` 신규 생성
  - 독바 SVG 디자인 기반, 모든 색상을 흰색(#FFFFFF)으로 통일
  - 배경 투명
- [ ] `build-work/gen_icon/gen_icon.js`
  - `tray_icon_mono.svg` → `app/ui/tray_icon.png` (22px) 생성 단계 추가
  - `tray_icon_mono.svg` → `app/ui/tray_icon@2x.png` (44px) 생성 단계 추가
- [ ] `app/main.go`
  - 트레이 아이콘 설정 코드에서 `systray.SetTitle("캡쳐고")` 제거 (아이콘으로 대체)
  - `systray.SetIcon(iconData)` 로 PNG 바이트를 직접 로드하도록 변경
  - `app/ui/tray_icon.png` 또는 `@2x` 를 `go:embed` 로 내장

## 아이콘 디자인 참고
- macOS 메뉴바 아이콘 가이드라인: 템플릿 이미지(흑백) 권장
- 크기: 22×22px (1x), 44×44px (2x / Retina)
- 여백: 상하좌우 2px 내외

## 작업 결과

- `build-work/gen_icon/tray_icon_mono.svg`: 흰색 크로스헤어 + 코너 프레임 디자인 신규 생성 (투명 배경)
- `build-work/gen_icon/gen_icon.js`: TRAY_SVG 상수 추가, 트레이 아이콘 생성 시 `tray_icon_mono.svg` 사용
- `app/main.go`: `systray.SetTitle("캡쳐고")` → `systray.SetIcon(trayIcon)` 변경, `//go:embed ui/tray_icon.png` 추가
- `app/ui/tray_icon.png` / `tray_icon@2x.png`: 흑백 모노 아이콘으로 재생성 (22px / 44px)
- `go build ./...` 통과 확인

## 추가 작업
- [ ] 트레이 아이콘이 투명 배경으로 되어 있음. 흰색 배경으로 변경 필요. 잘 안보임

## 추가 작업 결과

- `build-work/gen_icon/tray_icon_template.svg`: 검정색(#000000) 버전 신규 생성 (macOS 템플릿 이미지용)
- `build-work/gen_icon/gen_icon.js`: 템플릿 아이콘 생성 단계 추가 (`tray_icon_template.png`, `@2x`)
- `app/ui/tray_icon_template.png` / `tray_icon_template@2x.png`: 검정색 템플릿 아이콘 생성 (22px / 44px)
- `app/main.go`: `systray.SetIcon` → `systray.SetTemplateIcon(trayIconTemplate, trayIcon)` 변경
  - macOS가 다크/라이트 모드에 따라 자동으로 색상 반전 처리
