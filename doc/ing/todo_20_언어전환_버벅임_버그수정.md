# todo: 언어 전환 버벅임 버그 수정

## 작업 목표
EN/한국어 버튼을 누를 때 글자들이 순차적으로 로딩되며 변경되는 버벅임 현상을 수정한다.

## 원인 분석

현재 `applyLang()` 함수는 `document.querySelectorAll('[data-i18n]')`로 DOM 전체를 순회하며
각 요소의 `textContent`를 개별 업데이트한다. 이 과정에서:
1. DOM 업데이트가 하나씩 리플로우를 유발할 수 있음
2. `loadPermissions()` 호출이 async fetch를 포함하여 권한 패널만 늦게 갱신됨
3. 브라우저가 각 textContent 변경마다 레이아웃을 재계산할 가능성

## 수정 방향

- 언어 전환 시 페이지 전체를 **한 번에 교체** 하는 방식으로 변경
  - `document.documentElement.lang` 속성을 `ko` / `en` 으로 변경하고
  - CSS `[lang="en"] [data-ko]`, `[lang="ko"] [data-en]` 패턴으로 숨김/표시 처리
  - 또는: `applyLang()` 내부에서 모든 변경을 `requestAnimationFrame` 콜백 안에서 일괄 처리
- `loadPermissions()` 는 언어 전환 시 fetch 재호출 없이 캐시된 데이터로 재렌더링

## 세부 작업

- [ ] `app/server/static/app.js`
  - `applyLang()` 내 DOM 업데이트를 `requestAnimationFrame` 1회 콜백으로 묶음
  - 권한 데이터를 전역 변수에 캐시, 언어 전환 시 fetch 없이 재렌더링
  - `loadPermissions()` 결과를 `cachedPermData` 변수에 저장
  - `renderPermissions(data)` 함수로 분리하여 언어 전환 시 재활용

## 작업 결과

todo_19 작업과 함께 수정 완료.

- `applyLang()`: 모든 DOM 업데이트를 `requestAnimationFrame` 1회 콜백으로 묶어 일괄 처리
- `cachedPermData` 전역 변수에 권한 데이터 캐시
- `renderPermissions(data)` 함수 분리 → 언어 전환 시 fetch 없이 재렌더링
- `setLang()` 에서 캐시 데이터로 권한 패널 즉시 재렌더링 → fetch 지연 없음
