// 공통 UI/i18n/Theme 유틸리티

let lang = 'en';
let darkMode = false;
let pageI18n = {}; // 각 페이지에서 설정한 i18n 객체

const commonI18n = {
  ko: {
    btn_back_settings: '설정으로 돌아가기',
    footer_oss_link: '오픈소스 라이선스',
  },
  en: {
    btn_back_settings: 'Back to Settings',
    footer_oss_link: 'Open Source Licenses',
  }
};

function renderTopBar({ titleKey, backHref, hideTitle }) {
  const topBar = document.getElementById('top-bar');
  if (!topBar) return;
  topBar.innerHTML = `
    <div class="top-bar-left">
      ${backHref ? `<a class="back" href="${backHref}">&#8592; <span data-i18n="btn_back_settings"></span></a>` : ''}
      ${!hideTitle ? `<h1>CaptureGo <span data-i18n="${titleKey}"></span></h1>` : ''}
    </div>
    <div class="lang-btn-group">
      <button class="theme-btn" id="btnTheme" onclick="toggleTheme()" title="Toggle dark/light mode">🌙</button>
      <button class="lang-btn" id="btnEN" onclick="setLang('en')">EN</button>
      <button class="lang-btn" id="btnKO" onclick="setLang('ko')">한국어</button>
    </div>
  `;
}

function renderFooter() {
  const footer = document.getElementById('common-footer');
  if (!footer) return;
  footer.innerHTML = `
    <div>github.com/danielk0121 &mdash; <span id="buildtime">v-dev</span></div>
    <div><a href="/license" data-i18n="footer_oss_link"></a></div>
  `;
}

function applyTheme(dark) {
  darkMode = dark;
  document.body.classList.toggle('dark', dark);
  document.body.classList.toggle('light', !dark);
  const btn = document.getElementById('btnTheme');
  if (btn) btn.textContent = dark ? '☀️' : '🌙';
}

async function toggleTheme() {
  const next = !darkMode;
  applyTheme(next);
  try {
    await fetch('/api/darkmode', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ dark_mode: next }),
    });
  } catch (e) {}
}

function applyLang() {
  const t = pageI18n[lang] || {};
  const c = commonI18n[lang] || {};
  
  document.querySelectorAll('[data-i18n]').forEach(el => {
    const key = el.getAttribute('data-i18n');
    const val = t[key] !== undefined ? t[key] : c[key];
    if (val !== undefined) {
      if (key.startsWith('modal_step')) {
        el.innerHTML = val;
      } else {
        el.textContent = val;
      }
    }
  });

  document.querySelectorAll('[data-i18n-placeholder]').forEach(el => {
    const key = el.getAttribute('data-i18n-placeholder');
    const val = t[key] !== undefined ? t[key] : c[key];
    if (val !== undefined) el.placeholder = val;
  });

  document.documentElement.lang = lang;
  
  const btnEN = document.getElementById('btnEN');
  const btnKO = document.getElementById('btnKO');
  if (btnEN) btnEN.classList.toggle('active', lang === 'en');
  if (btnKO) btnKO.classList.toggle('active', lang === 'ko');

  // 페이지별 추가 처리 (hook)
  if (typeof onLangApplied === 'function') {
    onLangApplied(lang, t);
  }
}

function setLang(targetLang) {
  lang = targetLang;
  applyLang();
  fetch('/api/config', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ language: targetLang }),
  }).catch(() => {});
}

async function loadBuildtime() {
  try {
    const res = await fetch('/api/buildtime');
    const data = await res.json();
    const el = document.getElementById('buildtime');
    if (el && data.buildtime) el.textContent = data.buildtime;
  } catch (e) {}
}

async function initPage(i18nObj, topBarConfig) {
  pageI18n = i18nObj;
  renderTopBar(topBarConfig);
  renderFooter();
  
  try {
    const res = await fetch('/api/config');
    const cfg = await res.json();
    lang = cfg.language || 'en';
    applyTheme(!!cfg.dark_mode);
    applyLang();
    loadBuildtime();
    return cfg;
  } catch (e) {
    applyLang();
    loadBuildtime();
    return null;
  }
}
