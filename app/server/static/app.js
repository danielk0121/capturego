const i18n = {
  ko: {
    title_suffix: '설정',
    section_status: '현황',
    label_capture_count: '누적 캡처 횟수',
    label_license: '라이선스',
    section_permissions: 'macOS 권한',
    perm_loading: '권한 확인 중...',
    perm_screen_recording: '화면 기록',
    perm_screen_recording_desc: '스크린샷 캡처에 필요합니다',
    perm_accessibility: '손쉬운 사용',
    perm_accessibility_desc: '글로벌 단축키 감지에 필요합니다',
    perm_granted: '허용됨 ✓',
    perm_denied: '미허용 ✗',
    perm_open_settings: '설정 열기',
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
    modal_title: '권한 설정이 필요합니다',
    modal_sub: 'CaptureGo가 정상 작동하려면 아래 두 가지 권한이 필요합니다.',
    modal_step1: '<b>화면 기록 권한</b>: 시스템 설정 → 개인 정보 보호 및 보안 → <a href="x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture" onclick="openPref(this)">화면 기록</a> → CaptureGo 허용',
    modal_step2: '<b>손쉬운 사용 권한</b>: 시스템 설정 → 개인 정보 보호 및 보안 → <a href="x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility" onclick="openPref(this)">손쉬운 사용</a> → CaptureGo 허용',
    modal_step3: '권한 부여 후 <b>앱을 재시작</b>하면 모든 기능이 활성화됩니다.',
    modal_btn_close: '확인했습니다',
  },
  en: {
    title_suffix: 'Settings',
    section_status: 'Status',
    label_capture_count: 'Total Captures',
    label_license: 'License',
    section_permissions: 'macOS Permissions',
    perm_loading: 'Checking permissions...',
    perm_screen_recording: 'Screen Recording',
    perm_screen_recording_desc: 'Required for screenshot capture',
    perm_accessibility: 'Accessibility',
    perm_accessibility_desc: 'Required for global hotkey detection',
    perm_granted: 'Granted ✓',
    perm_denied: 'Not Granted ✗',
    perm_open_settings: 'Open Settings',
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
    modal_title: 'Permissions Required',
    modal_sub: 'CaptureGo needs the following two permissions to work properly.',
    modal_step1: '<b>Screen Recording</b>: System Settings → Privacy & Security → <a href="x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture" onclick="openPref(this)">Screen Recording</a> → Allow CaptureGo',
    modal_step2: '<b>Accessibility</b>: System Settings → Privacy & Security → <a href="x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility" onclick="openPref(this)">Accessibility</a> → Allow CaptureGo',
    modal_step3: 'After granting permissions, <b>restart the app</b> to activate all features.',
    modal_btn_close: 'Got it',
  },
};

let lang = localStorage.getItem('cg_lang') || 'en';

function applyLang() {
  const t = i18n[lang];
  requestAnimationFrame(() => {
    document.querySelectorAll('[data-i18n]').forEach(el => {
      const key = el.getAttribute('data-i18n');
      if (t[key] !== undefined) {
        if (key.startsWith('modal_step')) {
          el.innerHTML = t[key];
        } else {
          el.textContent = t[key];
        }
      }
    });
    document.querySelectorAll('[data-i18n-placeholder]').forEach(el => {
      const key = el.getAttribute('data-i18n-placeholder');
      if (t[key] !== undefined) el.placeholder = t[key];
    });
    document.documentElement.lang = lang;
    // 언어 버튼 active 상태
    document.getElementById('btnEN').classList.toggle('active', lang === 'en');
    document.getElementById('btnKO').classList.toggle('active', lang === 'ko');
    // 라이선스 상태 텍스트 갱신
    const licEl = document.getElementById('licenseStatus');
    if (licEl && licEl.dataset.activated === '1') {
      licEl.textContent = t.license_activated;
    } else if (licEl && licEl.dataset.activated === '0') {
      licEl.textContent = t.license_inactive;
    }
  });
}

function setLang(targetLang) {
  lang = targetLang;
  localStorage.setItem('cg_lang', lang);
  applyLang();
  // 캐시된 권한 데이터로 재렌더링 (fetch 없이)
  if (cachedPermData !== null) {
    renderPermissions(cachedPermData);
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
    licEl.textContent = cfg.license_activated
      ? i18n[lang].license_activated
      : i18n[lang].license_inactive;
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

// 권한 설정 열기 (앱 내 URL scheme 사용)
function openPref(el) {
  // <a> 클릭 시 기본 동작(브라우저 이동)을 막고 /api/open-url로 위임
  // 단순히 링크를 열도록 처리 — 브라우저가 직접 URL scheme을 처리한다
}

let cachedPermData = null;

function renderPermissions(data) {
  const t = i18n[lang];
  const container = document.getElementById('permList');

  const SCREEN_URL = 'x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture';
  const ACCESS_URL = 'x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility';

  const rows = [
    {
      label: t.perm_screen_recording,
      desc: t.perm_screen_recording_desc,
      granted: data.screen_recording,
      settingsUrl: SCREEN_URL,
    },
    {
      label: t.perm_accessibility,
      desc: t.perm_accessibility_desc,
      granted: data.accessibility,
      settingsUrl: ACCESS_URL,
    },
  ];

  container.innerHTML = rows.map(row => `
    <div class="perm-row">
      <div class="perm-row-left">
        <span>${row.label}</span>
        <span class="perm-row-desc">${row.desc}</span>
      </div>
      <div class="perm-row-right">
        <span class="perm-badge ${row.granted ? 'granted' : 'denied'}">
          ${row.granted ? t.perm_granted : t.perm_denied}
        </span>
        ${!row.granted ? `<a class="perm-open-btn" href="${row.settingsUrl}">${t.perm_open_settings}</a>` : ''}
      </div>
    </div>
  `).join('');
}

async function loadPermissions() {
  const t = i18n[lang];
  const container = document.getElementById('permList');
  try {
    const res = await fetch('/api/permissions');
    const data = await res.json();
    cachedPermData = data;
    renderPermissions(data);

    // 첫 실행 모달: 권한 미부여 + cg_welcomed 미설정 시 표시
    const anyDenied = !data.screen_recording || !data.accessibility;
    if (anyDenied && !localStorage.getItem('cg_welcomed')) {
      document.getElementById('permModal').classList.remove('hidden');
    }
  } catch (e) {
    container.innerHTML = `<span class="perm-loading">${t.perm_loading}</span>`;
  }
}

function closePermModal() {
  localStorage.setItem('cg_welcomed', '1');
  document.getElementById('permModal').classList.add('hidden');
}

async function loadBuildtime() {
  try {
    const res = await fetch('/api/buildtime');
    const data = await res.json();
    const el = document.getElementById('buildtime');
    if (el && data.buildtime) el.textContent = data.buildtime;
  } catch (e) {
    // v-dev 유지
  }
}

applyLang();
loadConfig();
loadPermissions();
loadBuildtime();
