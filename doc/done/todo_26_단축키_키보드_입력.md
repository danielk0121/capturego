# todo: 단축키 입력 — 키보드 직접 입력으로 등록

## 작업 목표
단축키 설정 input에서 사용자가 키보드를 직접 누르면 해당 조합이 자동으로 입력되도록 변경한다.
텍스트를 손으로 타이핑하는 방식 대신, 실제 키 입력 이벤트를 감지하여 `cmd+shift+6` 형식 문자열로 변환한다.

## 세부 작업

- [ ] `app/server/static/app.js`: `attachHotkeyInput(inputEl)` 함수 구현
  - `keydown` 이벤트 감지
  - modifier(cmd/ctrl/alt/shift) + 일반 키 조합을 `cmd+shift+6` 형식으로 변환
  - `Escape` 키 입력 시 이전 값 복원
  - Tab/Enter 등 단독 modifier는 무시
- [ ] `app/server/static/index.html`: hotkey input에 readonly 속성 추가 (직접 타이핑 방지)
- [ ] `app/server/static/style.css`: hotkey input에 포커스 시 안내 스타일 추가 (선택)

## 구현 상세

### 키 변환 규칙
| 이벤트 속성 | 문자열 |
|------------|--------|
| `e.metaKey` | `cmd` |
| `e.ctrlKey` | `ctrl` |
| `e.altKey` | `alt` |
| `e.shiftKey` | `shift` |
| `e.key` (일반 키) | 소문자 변환 (숫자/문자) |

순서: `ctrl` → `alt` → `shift` → `cmd` → 일반키 (modifier 간 충돌 방지)

### 단독 modifier 키 목록 (무시)
`Control`, `Alt`, `Shift`, `Meta`, `OS`

### placeholder 변경 (i18n)
- ko: `단축키 입력란을 클릭 후 키를 누르세요`
- en: `Click and press your shortcut keys`

## 작업 결과

- `app/server/static/app.js`: `attachHotkeyInput(inputEl)` 함수 구현
  - `keydown` 이벤트로 modifier+key 조합 감지 → `cmd+shift+6` 형식 문자열 자동 생성
  - `Escape` 입력 시 이전 값 복원 후 blur
  - 단독 modifier 키 입력은 무시
  - i18n placeholder/hint 문구 변경 (ko/en)
- `app/server/static/index.html`: hotkeyCapture, hotkeyScroll input에 `readonly` 추가
- `app/server/static/style.css`: `readonly` input 포커스 시 녹색 테두리 스타일 추가
- e2e 검증: `cmd+shift+6` 입력 → `shift+cmd+6` 표시 정상, `Esc` → 이전 값 복원 정상
