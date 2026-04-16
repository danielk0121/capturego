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

# 출력 디렉토리 초기화
rm -rf "$OUT_DIR"
mkdir -p "$BUNDLE/Contents/MacOS"
mkdir -p "$BUNDLE/Contents/Resources"

# Go 바이너리 빌드 (Apple Silicon)
echo "[1/3] Go 바이너리 빌드 (arm64)..."
cd "$APP_SRC"
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 \
  go build -ldflags="-s -w" -o "$BUNDLE/Contents/MacOS/$BINARY_NAME" .

echo "[2/3] Info.plist 복사..."
cp "$SCRIPT_DIR/Info.plist" "$BUNDLE/Contents/Info.plist"

# 앱 아이콘 복사 (있는 경우)
if [ -f "$SCRIPT_DIR/AppIcon.icns" ]; then
  echo "       AppIcon.icns 복사..."
  cp "$SCRIPT_DIR/AppIcon.icns" "$BUNDLE/Contents/Resources/AppIcon.icns"
else
  echo "       [주의] AppIcon.icns 없음 — 기본 아이콘 사용"
fi

echo "[3/3] 빌드 결과 확인..."
ls -lh "$BUNDLE/Contents/MacOS/$BINARY_NAME"
file "$BUNDLE/Contents/MacOS/$BINARY_NAME"

echo ""
echo "=== 빌드 완료 ==="
echo "경로: $BUNDLE"
echo ""
echo "실행하려면:"
echo "  open \"$BUNDLE\""
echo ""
echo "또는 /Applications 로 이동:"
echo "  cp -r \"$BUNDLE\" /Applications/"
