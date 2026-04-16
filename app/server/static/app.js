const i18n = {
  ko: {
    title_suffix: '설정',
    section_status: '현황',
    label_capture_count: '누적 캡처 횟수',
    label_license: '라이선스 상태',
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
    placeholder_save_dir: '예: /Users/user/screenshot_capturego',
    hint_save_dir: '절대 경로를 입력하세요. 폴더가 없으면 자동 생성됩니다.',
    section_hotkeys: '단축키',
    label_hotkey_capture: '듀얼 세이브 캡처',
    placeholder_hotkey_capture: '입력란 클릭 후 키를 누르세요',
    label_hotkey_scroll: '스크롤 캡처',
    placeholder_hotkey_scroll: '입력란 클릭 후 키를 누르세요',
    hint_hotkey: '입력란을 클릭한 뒤 단축키를 누르면 자동 등록됩니다. (Esc: 취소)',
    section_license: '라이선스 관리',
    label_license_key: '라이선스 키',
    hint_license_key: '키 입력 후 저장하면 Nagware 팝업이 영구 해제됩니다.',
    btn_save: '저장',
    btn_reset: '기본값으로 초기화',
    msg_reset_confirm: '모든 설정을 기본값으로 초기화하시겠습니까?',
    msg_reset_done: '설정이 기본값으로 초기화되었습니다.',
    msg_reset_fail: '초기화에 실패했습니다.',
    license_activated: '인증됨 ✓',
    license_inactive: '미인증',
    msg_saved: '설정이 저장되었습니다.',
    msg_save_fail: '저장에 실패했습니다.',
    msg_load_fail: '설정을 불러오지 못했습니다.',
    msg_save_error: '저장 중 오류가 발생했습니다.',
    footer_dev: '개발자:',
    footer_oss_link: '오픈소스 라이선스',
    btn_license_key: '라이선스 키 관리',
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
    label_license: 'License Status',
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
    placeholder_save_dir: 'e.g. /Users/user/screenshot_capturego',
    hint_save_dir: 'Enter an absolute path. The folder will be created if it does not exist.',
    section_hotkeys: 'Hotkeys',
    label_hotkey_capture: 'Dual Save Capture',
    placeholder_hotkey_capture: 'Click and press your shortcut keys',
    label_hotkey_scroll: 'Scroll Capture',
    placeholder_hotkey_scroll: 'Click and press your shortcut keys',
    hint_hotkey: 'Click the field and press keys to register. (Esc: cancel)',
    section_license: 'License',
    label_license_key: 'License Key',
    hint_license_key: 'Save with a key to permanently dismiss the Nagware popup.',
    btn_save: 'Save',
    btn_reset: 'Reset to Defaults',
    msg_reset_confirm: 'Reset all settings to default values?',
    msg_reset_done: 'Settings have been reset to defaults.',
    msg_reset_fail: 'Failed to reset settings.',
    license_activated: 'Activated ✓',
    license_inactive: 'Not activated',
    msg_saved: 'Settings saved.',
    msg_save_fail: 'Failed to save.',
    msg_load_fail: 'Failed to load settings.',
    msg_save_error: 'An error occurred while saving.',
    footer_dev: 'Developer:',
    footer_oss_link: 'Open Source Licenses',
    btn_license_key: 'License Key',
    modal_title: 'Permissions Required',
    modal_sub: 'CaptureGo needs the following two permissions to work properly.',
    modal_step1: '<b>Screen Recording</b>: System Settings → Privacy & Security → <a href="x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture" onclick="openPref(this)">Screen Recording</a> → Allow CaptureGo',
    modal_step2: '<b>Accessibility</b>: System Settings → Privacy & Security → <a href="x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility" onclick="openPref(this)">Accessibility</a> → Allow CaptureGo',
    modal_step3: 'After granting permissions, <b>restart the app</b> to activate all features.',
    modal_btn_close: 'Got it',
  },
};

// lang, darkMode 변수 및 applyTheme, toggleTheme, applyLang, setLang, loadBuildtime 함수는 common.js의 것을 사용

async function loadConfig() {
  try {
    const cfg = await initPage(i18n, { titleKey: 'title_suffix' });
    if (!cfg) return;

    document.getElementById('saveDir').value = cfg.save_directory || '';
    document.getElementById('hotkeyCapture').value = cfg.hotkey_capture || '';
    document.getElementById('hotkeyScroll').value = cfg.hotkey_scroll || '';
    
    const licEl = document.getElementById('licenseStatus');
    licEl.dataset.activated = cfg.license_activated ? '1' : '0';
    // initPage에서 호출된 applyLang이 처리하겠지만, 텍스트 갱신을 위해 onLangApplied 명시 호출
    onLangApplied(lang, i18n[lang]);
    
    loadPermissions();
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
  
  if (cachedPermData !== null) {
    renderPermissions(cachedPermData);
  }
}

async function saveConfig() {
  const body = {
    save_directory: document.getElementById('saveDir').value.trim(),
    hotkey_capture: document.getElementById('hotkeyCapture').value.trim(),
    hotkey_scroll:  document.getElementById('hotkeyScroll').value.trim(),
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

async function resetConfig() {
  if (!confirm(i18n[lang].msg_reset_confirm)) return;
  try {
    const res = await fetch('/api/config/reset', { method: 'POST' });
    if (res.ok) {
      showStatus(i18n[lang].msg_reset_done, true);
      await loadConfig();
    } else {
      showStatus(i18n[lang].msg_reset_fail, false);
    }
  } catch (e) {
    showStatus(i18n[lang].msg_reset_fail, false);
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
}

let cachedPermData = null;

function renderPermissions(data) {
  const t = i18n[lang];
  
  const updateRow = (id, granted) => {
    const root = document.getElementById(id);
    if (!root) return;
    const badge = root.querySelector('.perm-badge');
    const btn = root.querySelector('.perm-open-btn');
    
    badge.textContent = granted ? t.perm_granted : t.perm_denied;
    badge.className = `perm-badge ${granted ? 'granted' : 'denied'}`;
    
    if (granted) {
      btn.classList.add('hidden');
    } else {
      btn.classList.remove('hidden');
    }
  };

  updateRow('permScreen', data.screen_recording);
  updateRow('permAccess', data.accessibility);
}

async function loadPermissions() {
  try {
    const res = await fetch('/api/permissions');
    const data = await res.json();
    cachedPermData = data;
    renderPermissions(data);

    // 권한 미부여 시 안내 모달 표시
    const anyDenied = !data.screen_recording || !data.accessibility;
    if (anyDenied) {
      document.getElementById('permModal').classList.remove('hidden');
    }
  } catch (e) {
    console.error('Failed to load permissions:', e);
  }
}

function closePermModal() {
  document.getElementById('permModal').classList.add('hidden');
}

// 단축키 input: 키보드 직접 입력으로 단축키 등록
const MODIFIER_KEYS = new Set(['Control', 'Alt', 'Shift', 'Meta', 'OS']);

function attachHotkeyInput(inputEl) {
  let prevValue = '';

  inputEl.addEventListener('focus', () => {
    prevValue = inputEl.value;
  });

  inputEl.addEventListener('keydown', (e) => {
    e.preventDefault();

    // Escape: 이전 값 복원
    if (e.key === 'Escape') {
      inputEl.value = prevValue;
      inputEl.blur();
      return;
    }

    // 단독 modifier 키는 무시
    if (MODIFIER_KEYS.has(e.key)) return;

    const parts = [];
    if (e.ctrlKey)  parts.push('ctrl');
    if (e.altKey)   parts.push('alt');
    if (e.shiftKey) parts.push('shift');
    if (e.metaKey)  parts.push('cmd');

    // 일반 키: 소문자 변환
    const key = e.key.length === 1 ? e.key.toLowerCase() : e.key.toLowerCase();
    parts.push(key);

    inputEl.value = parts.join('+');
  });
}

// 초기화 시작
loadConfig();
attachHotkeyInput(document.getElementById('hotkeyCapture'));
attachHotkeyInput(document.getElementById('hotkeyScroll'));
