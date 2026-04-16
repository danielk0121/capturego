const i18n = {
  ko: {
    page_license_key_title: '라이선스 키 등록',
    section_status: '현황',
    label_license: '라이선스',
    section_license: '라이선스 키 등록',
    label_license_key: '라이선스 키',
    hint_license_key: '키 입력 후 저장하면 Nagware 팝업이 영구 해제됩니다.',
    btn_save: '저장',
    btn_back_settings: '설정으로 돌아가기',
    license_activated: '인증됨 ✓',
    license_inactive: '미인증',
    msg_saved: '라이선스가 등록되었습니다.',
    msg_save_fail: '저장에 실패했습니다.',
    msg_save_error: '저장 중 오류가 발생했습니다.',
    msg_load_fail: '설정을 불러오지 못했습니다.',
  },
  en: {
    page_license_key_title: 'License Key',
    section_status: 'Status',
    label_license: 'License',
    section_license: 'Register License Key',
    label_license_key: 'License Key',
    hint_license_key: 'Save with a key to permanently dismiss the Nagware popup.',
    btn_save: 'Save',
    btn_back_settings: 'Back to Settings',
    license_activated: 'Activated ✓',
    license_inactive: 'Not activated',
    msg_saved: 'License registered successfully.',
    msg_save_fail: 'Failed to save.',
    msg_save_error: 'An error occurred while saving.',
    msg_load_fail: 'Failed to load settings.',
  },
};

let lang = localStorage.getItem('cg_lang') || 'en';

function applyLang() {
  const t = i18n[lang];
  requestAnimationFrame(() => {
    document.querySelectorAll('[data-i18n]').forEach(el => {
      const key = el.getAttribute('data-i18n');
      if (t[key] !== undefined) el.textContent = t[key];
    });
    document.documentElement.lang = lang;
    document.getElementById('btnEN').classList.toggle('active', lang === 'en');
    document.getElementById('btnKO').classList.toggle('active', lang === 'ko');
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
}

async function loadStatus() {
  try {
    const res = await fetch('/api/config');
    const cfg = await res.json();
    const licEl = document.getElementById('licenseStatus');
    licEl.dataset.activated = cfg.license_activated ? '1' : '0';
    licEl.textContent = cfg.license_activated
      ? i18n[lang].license_activated
      : i18n[lang].license_inactive;
    // 다크모드 적용
    if (cfg.dark_mode) {
      document.body.classList.add('dark');
      document.body.classList.remove('light');
    } else {
      document.body.classList.add('light');
      document.body.classList.remove('dark');
    }
  } catch (e) {
    showStatus(i18n[lang].msg_load_fail, false);
  }
}

async function saveLicenseKey() {
  const key = document.getElementById('licenseKey').value.trim();
  if (!key) return;
  try {
    const res = await fetch('/api/config', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ license_key: key }),
    });
    const data = await res.json();
    if (res.ok) {
      showStatus(i18n[lang].msg_saved, true);
      loadStatus();
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

async function loadBuildtime() {
  try {
    const res = await fetch('/api/buildtime');
    const data = await res.json();
    const el = document.getElementById('buildtime');
    if (el && data.buildtime) el.textContent = data.buildtime;
  } catch (e) {}
}

applyLang();
loadStatus();
loadBuildtime();
