package server

// indexHTML 설정 UI HTML을 반환한다
func indexHTML() string {
	return `<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CaptureGo 설정</title>
  <link rel="icon" type="image/png" href="/static/favicon.png">
  <style>
    :root {
      --bg: #f5f5f7;
      --surface: #ffffff;
      --text: #1d1d1f;
      --subtext: #6e6e73;
      --border: #d2d2d7;
      --divider: #f5f5f7;
      --accent: #007aff;
      --accent-hover: #0066d6;
      --shadow: rgba(0,0,0,0.1);
    }
    @media (prefers-color-scheme: dark) {
      :root {
        --bg: #1c1c1e;
        --surface: #2c2c2e;
        --text: #f5f5f7;
        --subtext: #8e8e93;
        --border: #3a3a3c;
        --divider: #3a3a3c;
        --accent: #0a84ff;
        --accent-hover: #409cff;
        --shadow: rgba(0,0,0,0.4);
      }
    }
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
      background: var(--bg); color: var(--text);
      min-height: 100vh; padding: 40px 20px;
    }
    .container { max-width: 560px; margin: 0 auto; }
    .top-bar { display: flex; justify-content: space-between; align-items: center; margin-bottom: 32px; }
    h1 { font-size: 24px; font-weight: 600; }
    .lang-btn {
      background: var(--surface); border: 1px solid var(--border);
      color: var(--text); border-radius: 8px; padding: 6px 12px;
      font-size: 13px; cursor: pointer;
    }
    .lang-btn:hover { background: var(--border); }
    .card {
      background: var(--surface); border-radius: 12px; padding: 24px;
      box-shadow: 0 1px 3px var(--shadow); margin-bottom: 16px;
    }
    .card h2 { font-size: 14px; font-weight: 600; color: var(--subtext); margin-bottom: 16px; text-transform: uppercase; }
    label { display: block; font-size: 14px; margin-bottom: 6px; font-weight: 500; }
    input[type="text"] {
      width: 100%; padding: 10px 12px; border: 1px solid var(--border);
      border-radius: 8px; font-size: 14px; outline: none;
      background: var(--surface); color: var(--text);
    }
    input[type="text"]:focus { border-color: var(--accent); box-shadow: 0 0 0 3px rgba(0,122,255,0.15); }
    .form-group { margin-bottom: 16px; }
    .hint { font-size: 12px; color: var(--subtext); margin-top: 4px; }
    button.save-btn {
      width: 100%; padding: 12px; background: var(--accent); color: white; border: none;
      border-radius: 8px; font-size: 15px; font-weight: 500; cursor: pointer;
    }
    button.save-btn:hover { background: var(--accent-hover); }
    #status { text-align: center; margin-top: 12px; font-size: 14px; min-height: 20px; }
    .success { color: #34c759; }
    .error { color: #ff3b30; }
    .stat { display: flex; justify-content: space-between; font-size: 14px; padding: 8px 0; border-bottom: 1px solid var(--divider); }
    .stat:last-child { border-bottom: none; }
    .stat-val { font-weight: 600; }
    footer {
      text-align: center; margin-top: 32px; padding-bottom: 24px;
      font-size: 12px; color: var(--subtext); line-height: 1.8;
    }
    footer a { color: var(--subtext); text-decoration: none; }
    footer a:hover { color: var(--accent); }
  </style>
</head>
<body>
  <div class="container">
    <div class="top-bar">
      <h1>⚡ CaptureGo <span data-i18n="title_suffix"></span></h1>
      <button class="lang-btn" onclick="toggleLang()" id="langBtn">EN</button>
    </div>

    <div class="card">
      <h2 data-i18n="section_status"></h2>
      <div class="stat"><span data-i18n="label_capture_count"></span><span class="stat-val" id="captureCount">-</span></div>
      <div class="stat"><span data-i18n="label_license"></span><span class="stat-val" id="licenseStatus">-</span></div>
    </div>

    <div class="card">
      <h2 data-i18n="section_save_path"></h2>
      <div class="form-group">
        <label for="saveDir" data-i18n="label_save_dir"></label>
        <input type="text" id="saveDir" data-i18n-placeholder="placeholder_save_dir">
        <p class="hint" data-i18n="hint_save_dir"></p>
      </div>
    </div>

    <div class="card">
      <h2 data-i18n="section_hotkeys"></h2>
      <div class="form-group">
        <label for="hotkeyCapture" data-i18n="label_hotkey_capture"></label>
        <input type="text" id="hotkeyCapture" data-i18n-placeholder="placeholder_hotkey_capture">
      </div>
      <div class="form-group">
        <label for="hotkeyScroll" data-i18n="label_hotkey_scroll"></label>
        <input type="text" id="hotkeyScroll" data-i18n-placeholder="placeholder_hotkey_scroll">
        <p class="hint" data-i18n="hint_hotkey"></p>
      </div>
    </div>

    <div class="card">
      <h2 data-i18n="section_license"></h2>
      <div class="form-group">
        <label for="licenseKey" data-i18n="label_license_key"></label>
        <input type="text" id="licenseKey" placeholder="XXXX-XXXX-XXXX-XXXX">
        <p class="hint" data-i18n="hint_license_key"></p>
      </div>
    </div>

    <button class="save-btn" onclick="saveConfig()" data-i18n="btn_save"></button>
    <p id="status"></p>

    <footer>
      <div>CaptureGo &mdash; <span data-i18n="footer_dev"></span> <a href="mailto:maedk10@gmail.com">maedk10@gmail.com</a></div>
      <div><a href="/license" data-i18n="footer_oss_link"></a></div>
    </footer>
  </div>

  <script>
    const i18n = {
      ko: {
        title_suffix: '설정',
        section_status: '현황',
        label_capture_count: '누적 캡처 횟수',
        label_license: '라이선스',
        section_save_path: '저장 경로',
        label_save_dir: '캡처 파일 저장 폴더',
        placeholder_save_dir: '예: /Users/user/Pictures/CaptureGo',
        hint_save_dir: '절대 경로를 입력하세요. 폴더가 없으면 자동 생성됩니다.',
        section_hotkeys: '단축키',
        label_hotkey_capture: '듀얼 세이브 캡처',
        placeholder_hotkey_capture: '예: ctrl+shift+1',
        label_hotkey_scroll: '스크롤 캡처',
        placeholder_hotkey_scroll: '예: ctrl+shift+2',
        hint_hotkey: '지원 modifier: ctrl, shift, alt, cmd',
        section_license: '라이선스',
        label_license_key: '라이선스 키',
        hint_license_key: '키 입력 후 저장하면 Nagware 팝업이 영구 해제됩니다.',
        btn_save: '저장',
        license_activated: '인증됨 ✓',
        license_inactive: '미인증',
        msg_saved: '설정이 저장되었습니다.',
        msg_save_fail: '저장에 실패했습니다.',
        msg_load_fail: '설정을 불러오지 못했습니다.',
        msg_save_error: '저장 중 오류가 발생했습니다.',
        footer_dev: '개발자:',
        footer_oss_link: '오픈소스 라이선스',
      },
      en: {
        title_suffix: 'Settings',
        section_status: 'Status',
        label_capture_count: 'Total Captures',
        label_license: 'License',
        section_save_path: 'Save Path',
        label_save_dir: 'Capture Save Folder',
        placeholder_save_dir: 'e.g. /Users/user/Pictures/CaptureGo',
        hint_save_dir: 'Enter an absolute path. The folder will be created if it does not exist.',
        section_hotkeys: 'Hotkeys',
        label_hotkey_capture: 'Dual Save Capture',
        placeholder_hotkey_capture: 'e.g. ctrl+shift+1',
        label_hotkey_scroll: 'Scroll Capture',
        placeholder_hotkey_scroll: 'e.g. ctrl+shift+2',
        hint_hotkey: 'Supported modifiers: ctrl, shift, alt, cmd',
        section_license: 'License',
        label_license_key: 'License Key',
        hint_license_key: 'Save with a key to permanently dismiss the Nagware popup.',
        btn_save: 'Save',
        license_activated: 'Activated ✓',
        license_inactive: 'Not activated',
        msg_saved: 'Settings saved.',
        msg_save_fail: 'Failed to save.',
        msg_load_fail: 'Failed to load settings.',
        msg_save_error: 'An error occurred while saving.',
        footer_dev: 'Developer:',
        footer_oss_link: 'Open Source Licenses',
      },
    };

    let lang = localStorage.getItem('cg_lang') || 'ko';

    function applyLang() {
      const t = i18n[lang];
      document.querySelectorAll('[data-i18n]').forEach(el => {
        const key = el.getAttribute('data-i18n');
        if (t[key] !== undefined) el.textContent = t[key];
      });
      document.querySelectorAll('[data-i18n-placeholder]').forEach(el => {
        const key = el.getAttribute('data-i18n-placeholder');
        if (t[key] !== undefined) el.placeholder = t[key];
      });
      document.getElementById('langBtn').textContent = lang === 'ko' ? 'EN' : '한국어';
      document.documentElement.lang = lang;
    }

    function toggleLang() {
      lang = lang === 'ko' ? 'en' : 'ko';
      localStorage.setItem('cg_lang', lang);
      applyLang();
      // 라이선스 상태 텍스트 갱신
      const licEl = document.getElementById('licenseStatus');
      if (licEl.dataset.activated === '1') {
        licEl.textContent = i18n[lang].license_activated;
      } else if (licEl.dataset.activated === '0') {
        licEl.textContent = i18n[lang].license_inactive;
      }
    }

    async function loadConfig() {
      try {
        const res = await fetch('/api/config');
        const cfg = await res.json();
        document.getElementById('saveDir').value = cfg.save_directory || '';
        document.getElementById('hotkeyCapture').value = cfg.hotkey_capture || '';
        document.getElementById('hotkeyScroll').value = cfg.hotkey_scroll || '';
        document.getElementById('captureCount').textContent = cfg.capture_count ?? '-';
        const licEl = document.getElementById('licenseStatus');
        licEl.dataset.activated = cfg.license_activated ? '1' : '0';
        licEl.textContent = cfg.license_activated ? i18n[lang].license_activated : i18n[lang].license_inactive;
      } catch (e) {
        showStatus(i18n[lang].msg_load_fail, false);
      }
    }

    async function saveConfig() {
      const body = {
        save_directory: document.getElementById('saveDir').value.trim(),
        hotkey_capture: document.getElementById('hotkeyCapture').value.trim(),
        hotkey_scroll:  document.getElementById('hotkeyScroll').value.trim(),
        license_key:    document.getElementById('licenseKey').value.trim(),
      };
      try {
        const res = await fetch('/api/config', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(body),
        });
        const data = await res.json();
        if (res.ok) {
          showStatus(i18n[lang].msg_saved, true);
        } else {
          showStatus(data.error || i18n[lang].msg_save_fail, false);
        }
      } catch (e) {
        showStatus(i18n[lang].msg_save_error, false);
      }
    }

    function showStatus(msg, ok) {
      const el = document.getElementById('status');
      el.textContent = msg;
      el.className = ok ? 'success' : 'error';
      setTimeout(() => { el.textContent = ''; el.className = ''; }, 3000);
    }

    applyLang();
    loadConfig();
  </script>
</body>
</html>`
}

// licenseHTML 오픈소스 라이선스 페이지 HTML을 반환한다
func licenseHTML() string {
	return `<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CaptureGo — 오픈소스 라이선스</title>
  <link rel="icon" type="image/png" href="/static/favicon.png">
  <style>
    :root {
      --bg: #f5f5f7; --surface: #ffffff; --text: #1d1d1f;
      --subtext: #6e6e73; --border: #d2d2d7; --accent: #007aff;
      --shadow: rgba(0,0,0,0.1);
    }
    @media (prefers-color-scheme: dark) {
      :root {
        --bg: #1c1c1e; --surface: #2c2c2e; --text: #f5f5f7;
        --subtext: #8e8e93; --border: #3a3a3c; --accent: #0a84ff;
        --shadow: rgba(0,0,0,0.4);
      }
    }
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
      background: var(--bg); color: var(--text);
      min-height: 100vh; padding: 40px 20px;
    }
    .container { max-width: 700px; margin: 0 auto; }
    .back { display: inline-block; margin-bottom: 24px; font-size: 14px; color: var(--accent); text-decoration: none; }
    .back:hover { text-decoration: underline; }
    h1 { font-size: 24px; font-weight: 600; margin-bottom: 8px; }
    .intro { font-size: 14px; color: var(--subtext); margin-bottom: 32px; }
    .card {
      background: var(--surface); border-radius: 12px; padding: 20px 24px;
      box-shadow: 0 1px 3px var(--shadow); margin-bottom: 12px;
    }
    .card h2 { font-size: 15px; font-weight: 600; margin-bottom: 4px; }
    .card .meta { font-size: 13px; color: var(--subtext); margin-bottom: 8px; }
    .card .spdx {
      display: inline-block; font-size: 12px; font-weight: 500;
      background: var(--bg); border: 1px solid var(--border);
      border-radius: 4px; padding: 2px 8px; color: var(--subtext);
    }
    footer {
      text-align: center; margin-top: 32px; padding-bottom: 24px;
      font-size: 12px; color: var(--subtext);
    }
  </style>
</head>
<body>
  <div class="container">
    <a class="back" href="/">← 설정으로 돌아가기</a>
    <h1>오픈소스 라이선스</h1>
    <p class="intro">CaptureGo는 아래 오픈소스 라이브러리를 사용합니다.</p>

    <div class="card">
      <h2>gin-gonic/gin</h2>
      <div class="meta">HTTP web framework written in Go</div>
      <span class="spdx">MIT</span>
    </div>

    <div class="card">
      <h2>getlantern/systray</h2>
      <div class="meta">Cross-platform system tray library</div>
      <span class="spdx">Apache-2.0</span>
    </div>

    <div class="card">
      <h2>golang-design/hotkey</h2>
      <div class="meta">Global hotkey registration for Go</div>
      <span class="spdx">MIT</span>
    </div>

    <div class="card">
      <h2>nfnt/resize</h2>
      <div class="meta">Pure Go image resizing</div>
      <span class="spdx">ISC</span>
    </div>

    <div class="card">
      <h2>Go standard library</h2>
      <div class="meta">The Go programming language standard library</div>
      <span class="spdx">BSD-3-Clause</span>
    </div>

    <footer>CaptureGo &mdash; <a href="mailto:maedk10@gmail.com" style="color:var(--subtext)">maedk10@gmail.com</a></footer>
  </div>
</body>
</html>`
}
