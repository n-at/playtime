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

    /**
     * @param {string} cls
     * @param {string} text
     * @param {number} delay
     * @constructor
     */
    window.ShowToastMessage = function(cls, text, delay = 1000) {
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

        if (!delay) {
            delay = 1000;
        }

        document.getElementById('notifications').append(el);
        bootstrap.Toast.getOrCreateInstance(el, {delay: delay}).show();
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

    window.NetplayLoadClientName = defaultName => {
        if (!window.localStorage) {
            return;
        }
        if (window.localStorage.playtimeNetplayName) {
            return window.localStorage.playtimeNetplayName;
        } else {
            NetplaySaveClientName(defaultName);
            return defaultName;
        }
    };

    window.NetplaySaveClientName = name => {
        if (!window.localStorage) {
            return;
        }
        window.localStorage.playtimeNetplayName = name;
    };

    window.NetplayLoadClientPlayer = (gameId, sessionId, defaultPlayer) => {
        if (!window.localStorage || !window.localStorage.playtimeNetplayPlayer) {
            return defaultPlayer;
        }
        try {
            const players = JSON.parse(window.localStorage.playtimeNetplayPlayer);
            if (players[gameId] !== undefined && players[gameId][sessionId] !== undefined) {
                return players[gameId][sessionId];
            }
        } catch (e) {
            console.error(e);
        }
        window.NetplaySaveClientPlayer(gameId, sessionId, defaultPlayer);
        return defaultPlayer;
    };

    window.NetplaySaveClientPlayer = (gameId, sessionId, player) => {
        if (!window.localStorage) {
            return;
        }
        let players = {};
        if (window.localStorage.playtimeNetplayPlayer) {
            try {
                players = JSON.parse(window.localStorage.playtimeNetplayPlayer);
            } catch (e) {
                console.error(e);
            }
        }
        if (players[gameId] === undefined) {
            players[gameId] = {};
        }
        players[gameId][sessionId] = player;
        window.localStorage.playtimeNetplayPlayer = JSON.stringify(players);
    };

})();
