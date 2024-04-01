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

    window.addEventListener('load', () => {
        document.getElementById('auto-save-enabled').addEventListener('change', autoSaveOptionsVisibility);
        autoSaveOptionsVisibility();
    });

})();
