# todo: 웹앱 UI 컴팩트 (전체 높이 1000px 이내)

## 작업 목표
현재 스크롤이 필요한 설정 페이지를 전체 높이 약 1000px 이내의 컴팩트한 레이아웃으로 변경한다.

## 현황 분석
현재 카드 5개 + 버튼 + footer 구성으로 세로 스크롤 필요.
각 카드의 패딩, 여백, 폰트 크기를 줄여 1000px 내에 수납한다.

## 세부 작업

- [ ] `app/server/static/style.css`
  - `body` padding: `40px 20px` → `20px 16px`
  - `.card` padding: `24px` → `14px 18px`
  - `.card` margin-bottom: `16px` → `8px`
  - `.card h2` font-size: `14px` → `12px`, margin-bottom: `12px` → `8px`
  - `.top-bar` margin-bottom: `32px` → `16px`
  - `.form-group` margin-bottom: `16px` → `8px`
  - `input[type="text"]` padding: `10px 12px` → `7px 10px`
  - `.stat` padding: `8px 0` → `5px 0`
  - `.save-btn` padding: `12px` → `10px`
  - `footer` margin-top: `32px` → `16px`, padding-bottom: `24px` → `12px`
  - `h1` font-size: `24px` → `20px`
  - `label` font-size: `14px` → `13px`
  - `.hint` font-size: `12px` → `11px`
  - `.perm-row` padding: `10px 0` → `6px 0`

## 목표 레이아웃 높이 계산 (근사)
| 요소 | 높이(px) |
|------|---------|
| top-bar | 40 |
| macOS 권한 카드 | 110 |
| 저장 경로 카드 | 90 |
| 단축키 카드 | 110 |
| 현황 카드 | 80 |
| 라이선스 카드 | 80 |
| 저장 버튼 + status | 50 |
| footer | 60 |
| 여백 합계 | ~100 |
| **합계** | **~720** |

## 작업 결과

- `style.css` 패딩/마진 일괄 축소:
  - `body` padding `40px 20px` → `20px 16px`
  - `.top-bar` margin-bottom `32px` → `16px`
  - `.card` padding `24px` → `16px`, margin-bottom `16px` → `10px`
  - `.card h2` font-size `14px` → `13px`, margin-bottom `16px` → `10px`
  - `.form-group` margin-bottom `16px` → `10px`
  - `input[type="text"]` padding `10px 12px` → `8px 10px`
  - `.stat` padding `8px 0` → `6px 0`
  - `.perm-row` padding `10px 0` → `8px 0`
  - `.save-btn` padding `12px` → `10px`, font-size `15px` → `14px`
  - `footer` margin-top `32px` → `16px`, padding-bottom `24px` → `12px`
