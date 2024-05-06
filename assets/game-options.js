(() => {

    function autoSaveOptionsVisibility() {
        const enabled = document.getElementById('auto-save-enabled');
        const el = document.getElementById('auto-save-options');
        if (enabled.checked) {
            el.classList.remove('d-none');
        } else {
            el.classList.add('d-none');
        }
    }

    function netplayOptionsVisibility() {
        const enabled = document.getElementById("netplay-enabled");
        const el = document.getElementById("netplay-options");
        if (enabled.checked) {
            el.classList.remove('d-none');
        } else {
            el.classList.add('d-none');
        }
    }

    window.addEventListener('load', () => {
        const autoSaveEnabledEl = document.getElementById('auto-save-enabled');
        if (autoSaveEnabledEl) {
            autoSaveEnabledEl.addEventListener('change', autoSaveOptionsVisibility);
            autoSaveOptionsVisibility();
        }

        const netplayEnabledEl = document.getElementById('netplay-enabled');
        if (netplayEnabledEl) {
            netplayEnabledEl.addEventListener('change', netplayOptionsVisibility);
            netplayOptionsVisibility();
        }
    });

})();
