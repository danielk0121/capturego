const i18n = {
  ko: {
    page_license_key_title: '라이선스 키 등록',
    section_status: '현황',
    label_capture_count: '누적 캡처 횟수',
    label_license: '라이선스 상태',
    section_license: '라이선스 키 등록',
    label_license_key: '라이선스 키',
    hint_license_key: '키 입력 후 저장하면 Nagware 팝업이 영구 해제됩니다.',
    btn_save: '저장',
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
    label_capture_count: 'Total Captures',
    label_license: 'License Status',
    section_license: 'Register License Key',
    label_license_key: 'License Key',
    hint_license_key: 'Save with a key to permanently dismiss the Nagware popup.',
    btn_save: 'Save',
    license_activated: 'Activated ✓',
    license_inactive: 'Not activated',
    msg_saved: 'License registered successfully.',
    msg_save_fail: 'Failed to save.',
    msg_save_error: 'An error occurred while saving.',
    msg_load_fail: 'Failed to load settings.',
  },
};

// lang, darkMode 변수 등은 common.js 사용

async function loadStatus() {
  try {
    const cfg = await initPage(i18n, { titleKey: 'page_license_key_title', backHref: '/', hideTitle: true });
    if (!cfg) return;

    const countEl = document.getElementById('captureCount');
    if (countEl) countEl.textContent = (cfg.capture_count || 0).toLocaleString();
    
    const licEl = document.getElementById('licenseStatus');
    licEl.dataset.activated = cfg.license_activated ? '1' : '0';
    onLangApplied(lang, i18n[lang]);
  } catch (e) {
    showStatus(i18n[lang].msg_load_fail, false);
  }
}

function onLangApplied(currentLang, t) {
  const licEl = document.getElementById('licenseStatus');
  if (licEl && licEl.dataset.activated === '1') {
    licEl.textContent = t.license_activated;
  } else if (licEl && licEl.dataset.activated === '0') {
    licEl.textContent = t.license_inactive;
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

loadStatus();
