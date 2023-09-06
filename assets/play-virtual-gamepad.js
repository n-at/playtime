(() => {

    window.addEventListener('load', () => {
        if (!window.isMobile()) {
            document.getElementById('play-virtual-gamepad-control').classList.add('d-none');
        } else {
            document.getElementById('play-virtual-gamepad-toggle').addEventListener('click', virtualGamepadToggle);
            virtualGamepadLoad();
        }
    });

    let virtualGamepadVisible = false;

    function virtualGamepadLoad() {
        virtualGamepadVisible = window.isMobile();
        virtualGamepadShow();
    }

    function virtualGamepadToggle() {
        virtualGamepadVisible = !virtualGamepadVisible;
        if (window.localStorage) {
            window.localStorage.playtimeVirtualGamepad = (virtualGamepadVisible ? 'true' : 'false');
        }
        virtualGamepadShow();
    }

    function virtualGamepadShow() {
        EJS_emulator.toggleVirtualGamepad(virtualGamepadVisible);

        const btn = document.getElementById('play-virtual-gamepad-toggle');
        if (virtualGamepadVisible) {
            btn.classList.remove('btn-outline-secondary');
            btn.classList.add('btn-outline-success');
        } else {
            btn.classList.remove('btn-outline-success');
            btn.classList.add('btn-outline-secondary');
        }
    }

})();
