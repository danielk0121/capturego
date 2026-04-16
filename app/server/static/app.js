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

applyLang();
loadConfig();
