#!/bin/bash
# CaptureGo macOS .app 번들 빌드 스크립트
# 타겟: Apple Silicon (arm64)
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(dirname "$SCRIPT_DIR")"
APP_SRC="$REPO_ROOT/app"
OUT_DIR="$SCRIPT_DIR/dist"
APP_NAME="CaptureGo"
BINARY_NAME="capturego"
BUNDLE="$OUT_DIR/$APP_NAME.app"

echo "=== CaptureGo 빌드 시작 ==="

# 아이콘 생성 (SVG → icns, tray_icon, favicon)
echo "[0/5] 아이콘 생성..."
GEN_ICON_DIR="$SCRIPT_DIR/gen_icon"
if [ ! -d "$GEN_ICON_DIR/node_modules" ]; then
  echo "       npm install 실행 중..."
  cd "$GEN_ICON_DIR" && npm install --silent
fi
node "$GEN_ICON_DIR/gen_icon.js"

# buildtime.txt 생성 (KST 기준 v-yyyyMMdd-HHmm-kst 형식)
echo "[1/5] buildtime.txt 생성..."
BUILDTIME="v-$(TZ=Asia/Seoul date '+%Y%m%d-%H%M')-kst"
echo "$BUILDTIME" > "$APP_SRC/server/static/buildtime.txt"
echo "       buildtime: $BUILDTIME"

# 출력 디렉토리 초기화
rm -rf "$OUT_DIR"
mkdir -p "$BUNDLE/Contents/MacOS"
mkdir -p "$BUNDLE/Contents/Resources"

# Go 바이너리 빌드 (Apple Silicon)
echo "[2/5] Go 바이너리 빌드 (arm64)..."
cd "$APP_SRC"
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 \
  go build -ldflags="-s -w" -o "$BUNDLE/Contents/MacOS/$BINARY_NAME" .

echo "[3/5] Info.plist 복사..."
cp "$SCRIPT_DIR/Info.plist" "$BUNDLE/Contents/Info.plist"

# 앱 아이콘 복사 (있는 경우)
if [ -f "$SCRIPT_DIR/AppIcon.icns" ]; then
  echo "       AppIcon.icns 복사..."
  cp "$SCRIPT_DIR/AppIcon.icns" "$BUNDLE/Contents/Resources/AppIcon.icns"
else
  echo "       [주의] AppIcon.icns 없음 — 기본 아이콘 사용"
fi

echo "[4/5] 빌드 결과 확인..."
ls -lh "$BUNDLE/Contents/MacOS/$BINARY_NAME"
file "$BUNDLE/Contents/MacOS/$BINARY_NAME"

# 앱 번들 내 buildtime.txt 복사
cp "$APP_SRC/server/static/buildtime.txt" "$BUNDLE/Contents/buildtime.txt"

# 코드 서명 (ad-hoc): Apple Developer 계정 없이도 Gatekeeper 통과를 위해 필요
# '-s -' : ad-hoc 서명 (자체 서명, 무료)
# '--deep' : 번들 내 모든 실행파일 재귀 서명
# '--force' : 기존 서명 덮어쓰기
# '--options runtime' : Hardened Runtime 활성화 (Gatekeeper 신뢰성 향상)
echo "[4.5/5] 코드 서명 (ad-hoc)..."
codesign --deep --force --options runtime --sign - "$BUNDLE"
echo "       서명 완료"
codesign --verify --deep --strict "$BUNDLE" && echo "       서명 검증 OK"

# DMG 생성
DMG_NAME="${APP_NAME}-${BUILDTIME}.dmg"
DMG_PATH="$OUT_DIR/$DMG_NAME"

echo ""
echo "[5/5] DMG 생성..."

# 임시 마운트 폴더 구성
DMG_STAGING="$OUT_DIR/dmg_staging"
rm -rf "$DMG_STAGING"
mkdir -p "$DMG_STAGING"
cp -r "$BUNDLE" "$DMG_STAGING/"
# Applications 심볼릭 링크 추가 (드래그 인스톨 안내)
ln -s /Applications "$DMG_STAGING/Applications"

# 읽기/쓰기 임시 DMG 생성 후 압축 변환
TMP_DMG="$OUT_DIR/tmp_rw.dmg"
hdiutil create -volname "$APP_NAME" -srcfolder "$DMG_STAGING" \
  -ov -format UDRW "$TMP_DMG" > /dev/null
hdiutil convert "$TMP_DMG" -format UDZO -imagekey zlib-level=9 \
  -o "$DMG_PATH" > /dev/null
rm -f "$TMP_DMG"
rm -rf "$DMG_STAGING"

ls -lh "$DMG_PATH"

echo ""
echo "=== 빌드 완료 ==="
echo "앱 번들: $BUNDLE"
echo "DMG:     $DMG_PATH"
