// gen_icon.js: SVG → 각 용도별 아이콘 자동 생성
//
// 생성 목록:
//   - AppIcon.icns      build-work/AppIcon.icns      (.app 독바/Finder)
//   - tray_icon.png     app/ui/tray_icon.png          (시스템 트레이 메뉴바, 22px + @2x 44px)
//   - favicon.png       app/ui/static/favicon.png     (웹앱 파비콘, 32px)
//
// 실행: npm run gen  (또는 node gen_icon.js)
// 의존: sharp, macOS iconutil (Xcode Command Line Tools 포함)

const sharp = require('sharp');
const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const SCRIPT_DIR  = __dirname;
const REPO_ROOT   = path.resolve(SCRIPT_DIR, '..', '..');
const SVG_PATH    = path.join(SCRIPT_DIR, 'capturego_icon.svg');
const TRAY_SVG          = path.join(SCRIPT_DIR, 'tray_icon_mono.svg');
const TRAY_TEMPLATE_SVG = path.join(SCRIPT_DIR, 'tray_icon_template.svg');
const ICONSET_DIR = path.join(SCRIPT_DIR, 'out', 'icon.iconset');
const UI_DIR      = path.join(REPO_ROOT, 'app', 'ui');
const STATIC_DIR  = path.join(REPO_ROOT, 'app', 'server', 'static');

// macOS iconset 규격 (10종)
const iconsetSpecs = [
  { size: 16,   filename: 'icon_16x16.png' },
  { size: 32,   filename: 'icon_16x16@2x.png' },
  { size: 32,   filename: 'icon_32x32.png' },
  { size: 64,   filename: 'icon_32x32@2x.png' },
  { size: 128,  filename: 'icon_128x128.png' },
  { size: 256,  filename: 'icon_128x128@2x.png' },
  { size: 256,  filename: 'icon_256x256.png' },
  { size: 512,  filename: 'icon_256x256@2x.png' },
  { size: 512,  filename: 'icon_512x512.png' },
  { size: 1024, filename: 'icon_512x512@2x.png' },
];

async function renderSvg(svgBuf, size) {
  // 작은 사이즈(< 64px)는 큰 해상도로 먼저 렌더링 후 다운스케일해 품질을 높인다
  if (size < 64) {
    const buf = await sharp(svgBuf).resize(size * 4, size * 4).png().toBuffer();
    return sharp(buf).resize(size, size).png().toBuffer();
  }
  return sharp(svgBuf).resize(size, size).png().toBuffer();
}

async function main() {
  const svgBuf = fs.readFileSync(SVG_PATH);
  fs.mkdirSync(ICONSET_DIR, { recursive: true });
  fs.mkdirSync(STATIC_DIR, { recursive: true });

  // ── 1. AppIcon.icns ──────────────────────────────────────────────────────
  console.log('==> [1/3] iconset PNG 생성');
  for (const spec of iconsetSpecs) {
    const buf = await renderSvg(svgBuf, spec.size);
    const outPath = path.join(ICONSET_DIR, spec.filename);
    await sharp(buf).toFile(outPath);
    console.log(`  ${spec.filename} (${spec.size}px)`);
  }

  const icnsOut = path.join(REPO_ROOT, 'build-work', 'AppIcon.icns');
  execSync(`iconutil -c icns "${ICONSET_DIR}" -o "${icnsOut}"`);
  console.log(`  AppIcon.icns → ${icnsOut}`);

  // ── 2. 트레이 아이콘 ──────────────────────────────────────────────────────
  console.log('==> [2/3] 트레이 아이콘 생성');

  // 2-a. 흰색 모노 (흰색 배경 없는 원본, 보관용)
  const traySvgBuf = fs.readFileSync(TRAY_SVG);
  const trayOut    = path.join(UI_DIR, 'tray_icon.png');
  const tray2xOut  = path.join(UI_DIR, 'tray_icon@2x.png');

  const tray1xBuf = await renderSvg(traySvgBuf, 22);
  await sharp(tray1xBuf).toFile(trayOut);
  console.log(`  tray_icon.png (22px) → ${trayOut}`);

  const tray2xBuf = await renderSvg(traySvgBuf, 44);
  await sharp(tray2xBuf).toFile(tray2xOut);
  console.log(`  tray_icon@2x.png (44px) → ${tray2xOut}`);

  // 2-b. 검정색 템플릿 (macOS SetTemplateIcon용 — 다크/라이트 자동 대응)
  const trayTemplateSvgBuf = fs.readFileSync(TRAY_TEMPLATE_SVG);
  const trayTemplateOut    = path.join(UI_DIR, 'tray_icon_template.png');
  const trayTemplate2xOut  = path.join(UI_DIR, 'tray_icon_template@2x.png');

  const trayTemplate1xBuf = await renderSvg(trayTemplateSvgBuf, 22);
  await sharp(trayTemplate1xBuf).toFile(trayTemplateOut);
  console.log(`  tray_icon_template.png (22px) → ${trayTemplateOut}`);

  const trayTemplate2xBuf = await renderSvg(trayTemplateSvgBuf, 44);
  await sharp(trayTemplate2xBuf).toFile(trayTemplate2xOut);
  console.log(`  tray_icon_template@2x.png (44px) → ${trayTemplate2xOut}`);

  // ── 3. 파비콘 ────────────────────────────────────────────────────────────
  console.log('==> [3/3] 파비콘 생성');
  const faviconOut = path.join(STATIC_DIR, 'favicon.png');
  const faviconBuf = await renderSvg(svgBuf, 32);
  await sharp(faviconBuf).toFile(faviconOut);
  console.log(`  favicon.png (32px) → ${faviconOut}`);

  console.log('\n==> 완료');
}

main().catch(err => { console.error(err); process.exit(1); });
