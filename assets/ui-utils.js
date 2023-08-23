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
    };

    ///////////////////////////////////////////////////////////////////////////
    // Toast messages
    ///////////////////////////////////////////////////////////////////////////

    window.ShowToastMessage = function(cls, text) {
        const el = document.createElement('div');
        el.classList.add('toast', 'align-items-center', 'border-0', `text-bg-${cls}`);
        el.role = 'alert';

        const flexEl = document.createElement('div');
        flexEl.classList.add('d-flex');
        el.append(flexEl);

        const bodyEl = document.createElement('div');
        bodyEl.classList.add('toast-body');
        bodyEl.innerText = text;
        flexEl.append(bodyEl);

        const btnEl = document.createElement('button');
        btnEl.type = 'button';
        btnEl.classList.add('btn-close', 'btn-close-white', 'me-2', 'm-auto');
        btnEl.setAttribute('data-bs-dismiss', 'toast');
        btnEl.ariaLabel = 'Close';
        flexEl.append(btnEl);

        document.getElementById('notifications').append(el);
        bootstrap.Toast.getOrCreateInstance(el, {delay: 1000}).show();
    };

    ///////////////////////////////////////////////////////////////////////////
    // Display utils
    ///////////////////////////////////////////////////////////////////////////

    window.NetplayPlayerDisplay = (clientId, hostId, player) => {
        if (clientId === hostId) {
            return 'host';
        }
        if (player === -1) {
            return 'spectator'
        }
        return `player ${player + 1}`;
    };

})();
