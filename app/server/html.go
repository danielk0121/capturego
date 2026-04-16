package server

// indexHTML 설정 UI HTML을 반환한다
// go:embed를 사용하지 않고 인라인으로 관리한다 (todo_09 후속 작업에서 분리 예정)
func indexHTML() string {
	return `<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CaptureGo 설정</title>
  <style>
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
           background: #f5f5f7; color: #1d1d1f; min-height: 100vh; padding: 40px 20px; }
    .container { max-width: 560px; margin: 0 auto; }
    h1 { font-size: 24px; font-weight: 600; margin-bottom: 32px; }
    .card { background: white; border-radius: 12px; padding: 24px;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1); margin-bottom: 16px; }
    .card h2 { font-size: 14px; font-weight: 600; color: #6e6e73; margin-bottom: 16px; text-transform: uppercase; }
    label { display: block; font-size: 14px; margin-bottom: 6px; font-weight: 500; }
    input[type="text"] { width: 100%; padding: 10px 12px; border: 1px solid #d2d2d7;
                         border-radius: 8px; font-size: 14px; outline: none; }
    input[type="text"]:focus { border-color: #007aff; box-shadow: 0 0 0 3px rgba(0,122,255,0.15); }
    .form-group { margin-bottom: 16px; }
    .hint { font-size: 12px; color: #6e6e73; margin-top: 4px; }
    button { width: 100%; padding: 12px; background: #007aff; color: white; border: none;
             border-radius: 8px; font-size: 15px; font-weight: 500; cursor: pointer; }
    button:hover { background: #0066d6; }
    #status { text-align: center; margin-top: 12px; font-size: 14px; min-height: 20px; }
    .success { color: #34c759; }
    .error { color: #ff3b30; }
    .stat { display: flex; justify-content: space-between; font-size: 14px; padding: 8px 0;
            border-bottom: 1px solid #f5f5f7; }
    .stat:last-child { border-bottom: none; }
    .stat-val { font-weight: 600; }
  </style>
</head>
<body>
  <div class="container">
    <h1>⚡ CaptureGo 설정</h1>

    <div class="card">
      <h2>현황</h2>
      <div class="stat"><span>누적 캡처 횟수</span><span class="stat-val" id="captureCount">-</span></div>
      <div class="stat"><span>라이선스</span><span class="stat-val" id="licenseStatus">-</span></div>
    </div>

    <div class="card">
      <h2>저장 경로</h2>
      <div class="form-group">
        <label for="saveDir">캡처 파일 저장 폴더</label>
        <input type="text" id="saveDir" placeholder="예: /Users/user/Pictures/CaptureGo">
        <p class="hint">절대 경로를 입력하세요. 폴더가 없으면 자동 생성됩니다.</p>
      </div>
    </div>

    <div class="card">
      <h2>단축키</h2>
      <div class="form-group">
        <label for="hotkeyCapture">듀얼 세이브 캡처</label>
        <input type="text" id="hotkeyCapture" placeholder="예: ctrl+shift+1">
      </div>
      <div class="form-group">
        <label for="hotkeyScroll">스크롤 캡처</label>
        <input type="text" id="hotkeyScroll" placeholder="예: ctrl+shift+2">
        <p class="hint">지원 modifier: ctrl, shift, alt, cmd</p>
      </div>
    </div>

    <div class="card">
      <h2>라이선스</h2>
      <div class="form-group">
        <label for="licenseKey">라이선스 키</label>
        <input type="text" id="licenseKey" placeholder="XXXX-XXXX-XXXX-XXXX">
        <p class="hint">키 입력 후 저장하면 Nagware 팝업이 영구 해제됩니다.</p>
      </div>
    </div>

    <button onclick="saveConfig()">저장</button>
    <p id="status"></p>
  </div>

  <script>
    async function loadConfig() {
      try {
        const res = await fetch('/api/config');
        const cfg = await res.json();
        document.getElementById('saveDir').value = cfg.save_directory || '';
        document.getElementById('hotkeyCapture').value = cfg.hotkey_capture || '';
        document.getElementById('hotkeyScroll').value = cfg.hotkey_scroll || '';
        document.getElementById('captureCount').textContent = cfg.capture_count ?? '-';
        document.getElementById('licenseStatus').textContent =
          cfg.license_activated ? '인증됨 ✓' : '미인증';
      } catch (e) {
        showStatus('설정을 불러오지 못했습니다.', false);
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
          showStatus('설정이 저장되었습니다.', true);
        } else {
          showStatus(data.error || '저장에 실패했습니다.', false);
        }
      } catch (e) {
        showStatus('저장 중 오류가 발생했습니다.', false);
      }
    }

    function showStatus(msg, ok) {
      const el = document.getElementById('status');
      el.textContent = msg;
      el.className = ok ? 'success' : 'error';
      setTimeout(() => { el.textContent = ''; el.className = ''; }, 3000);
    }

    loadConfig();
  </script>
</body>
</html>`
}
