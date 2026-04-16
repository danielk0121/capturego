# todo: 아이콘 생성 시스템 구축

## 작업 목표
SVG 마스터 아이콘 1개에서 앱 전반에 필요한 모든 아이콘을 자동 생성하는 스크립트를 구축한다.
winresizer-go의 `build-work/gen_icon/` 방식(sharp + Node.js)을 그대로 차용한다.

## 생성할 아이콘 목록

| 아이콘 | 출력 경로 | 용도 | 비고 |
|--------|-----------|------|------|
| `AppIcon.icns` | `build-work/AppIcon.icns` | `.app` 독바 / Finder | iconset 10종 → iconutil |
| `tray_icon.png` | `app/ui/tray_icon.png` | 시스템 트레이 메뉴바 | 22px (+ 44px @2x) |
| `favicon.png` | `app/ui/static/favicon.png` | 웹앱 파비콘 (브라우저 탭) | 32px |
| `favicon.ico` | `app/ui/static/favicon.ico` | IE/구형 브라우저 호환 | 16+32px 멀티 |

> **알림센터 아이콘**: `osascript display notification`은 자동으로 `.app` 아이콘을 사용하므로 별도 파일 불필요.
> **독바 아이콘**: `.icns` 하나로 독바·Finder·스팟라이트 모두 커버.

## 세부 작업

- [x] 마스터 SVG 디자인: `build-work/gen_icon/capturego_icon.svg`
  - 짙은 블루-그레이 배경 + 캡처 코너 프레임 + 크로스헤어 조준점 모티프
- [x] `build-work/gen_icon/gen_icon.js` 스크립트 작성
  - sharp로 SVG → PNG 각 해상도 변환
  - iconutil로 `.icns` 생성 → `build-work/AppIcon.icns` 출력
  - tray_icon.png (22px, @2x 44px) → `app/ui/tray_icon.png`
  - favicon.png (32px) → `app/server/static/favicon.png`
- [x] `build-work/gen_icon/package.json` 작성 (sharp 의존성)
- [x] `build.sh` 에 gen_icon 선행 실행 단계 추가
- [x] 웹앱 HTML에 `<link rel="icon">` 태그 추가
- [x] `app/server/web_server.go` 에 `/static/` 라우트 추가 (embed.FS)

## 참고
- 선행 프로젝트: `/Users/user/ws/winresizer-go/build-work/gen_icon/`
- iconset 규격: 16, 32, 64, 128, 256, 512, 1024px (10개 파일)
- tray 아이콘: macOS는 Template Image 방식(흑백 PNG)을 권장하나, systray 라이브러리에 따라 다름
