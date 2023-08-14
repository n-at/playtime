(() => {

    window.FlashButtonIcon = (btnId, oldBtnCls, oldIconCls, newBtnCls, newIconCls) => {
        const btn = document.getElementById(btnId);
        if (!btn) {
            return;
        }
        const icon = btn.firstElementChild;
        if (!icon) {
            return;
        }
        if (oldBtnCls) {
            btn.classList.remove(...oldBtnCls);
        }
        if (newBtnCls) {
            btn.classList.add(...newBtnCls)
        }
        if (oldIconCls) {
            icon.classList.remove(...oldIconCls);
        }
        if (newIconCls) {
            icon.classList.add(...newIconCls);
        }
        setTimeout(() => {
            if (newBtnCls) {
                btn.classList.remove(...newBtnCls);
            }
            if (oldBtnCls) {
                btn.classList.add(...oldBtnCls);
            }
            if (newIconCls) {
                icon.classList.remove(...newIconCls);
            }
            if (oldIconCls) {
                icon.classList.add(...oldIconCls);
            }
        }, 1000);
    }

})();
