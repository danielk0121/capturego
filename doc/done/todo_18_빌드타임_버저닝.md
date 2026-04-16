# todo: 빌드타임 기반 버저닝

## 작업 목표
빌드 시각(KST)을 기준으로 버전 정보를 앱 번들·웹앱 푸터·DMG 파일명에 일관되게 반영한다.

## 세부 작업

- [ ] `build-work/build.sh`: 빌드 시각(yyyyMMdd-HHmm KST) 계산 및 아래 항목에 적용
  - `CaptureGo.app/Contents/buildtime-yyyyMMdd-HHmm.txt` 파일 생성
  - DMG 파일명을 `CaptureGo-yyyyMMdd-HHmm.dmg` 형식으로 변경
- [ ] `app/server/static/index.html`: footer에 `v-yyyyMMdd-HHmm-kst` 문자열 표기
- [ ] `app/server/static/app.js`: 빌드타임 문자열을 `/api/buildtime` 또는 정적 파일에서 읽어 렌더링
- [ ] `app/server/web_server.go`: `/api/buildtime` 엔드포인트 추가 (buildtime.txt 읽어 반환)
- [ ] `app/core/buildtime.go`: 빌드타임 문자열을 embed 또는 런타임에 읽는 유틸 함수
- [ ] `README.md`: 버저닝 규칙 섹션 추가

## 구현 상세

### 빌드 시각 계산 (`build.sh`)
```bash
BUILD_TIME=$(TZ=Asia/Seoul date +"%Y%m%d-%H%M")
# → yyyyMMdd-HHmm (KST)
```

### buildtime.txt
- 위치: `CaptureGo.app/Contents/buildtime-${BUILD_TIME}.txt`
- 내용: `${BUILD_TIME}` (단순 문자열)
- 용도: 배포된 앱 번들의 빌드 시각 추적

### DMG 파일명
```
CaptureGo-yyyyMMdd-HHmm.dmg
```

### 웹앱 footer
- `app/server/static/` 에 `buildtime.txt` 파일을 빌드 시 생성 (embed.FS로 서빙)
- 앱 실행 시 해당 파일을 읽어 `/api/buildtime` 으로 반환
- `app.js` 초기화 시 fetch → footer에 `v-yyyyMMdd-HHmm-kst` 표기

### README 버저닝 규칙
- 버전 식별자: `yyyyMMdd-HHmm` (KST 기준 빌드 시각)
- SemVer 대신 빌드타임 기반 버저닝 사용
- DMG 파일명·앱 번들 내 buildtime.txt·웹앱 footer가 동일 시각을 공유

## 작업 결과

- `build-work/build.sh`:
  - `BUILDTIME="v-$(TZ=Asia/Seoul date '+%Y%m%d-%H%M')-kst"` 계산 (기존 todo_22에서 완료)
  - `CaptureGo.app/Contents/buildtime.txt` 복사 추가 (앱 번들 내 빌드 시각 추적)
  - DMG 파일명 `CaptureGo-${BUILDTIME}.dmg` 형식으로 변경 (기존 `CaptureGo-1.0.0.dmg` → buildtime 기반)
- `app/server/static/buildtime.txt`, `/api/buildtime`, 웹앱 footer 렌더링: todo_19/22에서 완료
