## 개요
- 다른 맥 PC 에 설치 하고 실행하니까, 이 앱은 손상되었다면서, 실행할 수 없다고 한다.
- 맥 미니 2024년 모델 apple m4 pro 칩, OS Tahoe 26.3
- 설치한 버전 : CaptureGo-1.1.0.dmg
- Commit 10aa07b

## 재현
- 에러 메세지 : ‘CaptureGo.app’은(는) 손상되었기 때문에 열 수 없습니다. 해당 항목을 휴지통으로 이동해야 합니다.
- 깃헙 릴리즈 페이지에서 dmg 파일 다운로드
- 다운로드 후 실행 > 맥에서 에러 발생
- 추정 : 아마도, 온라인에서 다운로드 받은 dmg 파일이라서 에러 발생하는 것으로 보임.

## 원인 분석
- macOS Gatekeeper는 인터넷에서 다운로드한 앱에 `com.apple.quarantine` 속성을 자동으로 부여함
- 앱 실행 시 Gatekeeper가 코드 서명을 검증하며, 서명이 없으면 "손상됨" 오류를 표시
- 기존 `build.sh`에는 코드 서명 단계가 없었음

## 작업 결과
- `build-work/build.sh`에 ad-hoc 코드 서명 단계([4.5/5]) 추가
- `codesign --deep --force --options runtime --sign -` 사용
  - `--sign -` : ad-hoc 서명 (Apple Developer 계정 불필요, 무료)
  - `--deep` : 번들 내 모든 실행파일 재귀 서명
  - `--options runtime` : Hardened Runtime 활성화
- 서명 후 `codesign --verify --deep --strict`로 검증 확인
- DMG 생성 전에 서명하므로, DMG에 포함된 앱도 서명된 상태로 배포됨

### 참고: ad-hoc 서명의 한계
- ad-hoc 서명은 Apple의 공식 검증을 받지 않아, 타 PC에서 여전히 경고창이 뜰 수 있음
- 완전한 해결을 위해선 Apple Developer Program($99/년) 가입 후 공증(notarize)이 필요
- 그러나 ad-hoc 서명만으로도 "손상됨" 오류는 해소되며, 시스템 환경설정 > 개인정보 및 보안에서 "그래도 열기"로 실행 가능

