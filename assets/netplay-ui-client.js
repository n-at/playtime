(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        setupGameDisplaySize();
        window.addEventListener('resize', setupGameDisplaySize);

        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);

        netplay = NetplayClient({
            debug: true,
            gameVideoEl: document.getElementById('game'),
            gameId: window.NetplayGameId,
            sessionId: window.NetplaySessionId,
            host: false,
            turnServerUrl: window.NetplayTurnServerUrl,
            turnServerUser: window.NetplayTurnServerUser,
            turnServerPassword: window.NetplayTurnServerPassword,

            onClientError: errorHandler,
            onSelfNameChanged: selfNameChanged,
            onSelfPlayerChanged: selfPlayerChanged,
            onClientConnected: clientConnected,
            onClientDisconnected: clientDisconnected,
            onClientNameChanged: clientNameChanged,
            onClientPlayerChanged: clientPlayerChanged,
        });
        netplay.connect();
    });

    ///////////////////////////////////////////////////////////////////////////
    // Self name and player
    ///////////////////////////////////////////////////////////////////////////

    function selfNameChanged(name) {
        document.getElementById('netplay-name').value = name;
        document.getElementById('netplay-player').innerText = `${netplay.getName()}: ${_displayPlayer(netplay.getPlayer())}`;
    }

    function selfPlayerChanged(player) {
        document.getElementById('netplay-player').innerText = `${netplay.getName()}: ${_displayPlayer(netplay.getPlayer())}`;

        //TODO give controller
    }

    function changeSelfName() {
        const name = document.getElementById('netplay-name').value;

        if (name.trim().length === 0 || name.length > 32) {
            window.FlashButtonIcon('netplay-name-change', ['btn-outline-secondary'], ['bi-check'], ['btn-outline-danger'], ['bi-x']);
            return;
        }

        netplay.setName(name);

        window.FlashButtonIcon('netplay-name-change', ['btn-outline-secondary'], [], ['btn-outline-success'], []);
    }

    ///////////////////////////////////////////////////////////////////////////
    // Game session clients list
    ///////////////////////////////////////////////////////////////////////////

    function clientConnected(id, name, player) {
        const elId = `netplay-client-${id}`;
        if (document.getElementById(elId)) {
            clientNameChanged(name);
            clientPlayerChanged(player);
            return;
        }

        const el = document.createElement('div');
        el.id = `netplay-client-${id}`;
        el.classList.add('list-group-item');

        const containerEl = document.createElement('div');
        containerEl.classList.add('row');
        el.append(containerEl);

        const hostEl = document.createElement('div');
        hostEl.classList.add('col-1');
        containerEl.append(hostEl);

        if (id === netplay.getHostId()) {
            const el = document.createElement('i');
            el.classList.add('bi', 'bi-star', 'lead', 'text-success');
            hostEl.append(el);
        }

        const nameEl = document.createElement('div');
        nameEl.id = `netplay-client-${id}-name`;
        nameEl.classList.add('lead', 'col-6', 'col-md-8');
        nameEl.innerText = name;
        containerEl.append(nameEl);

        const playerEl = document.createElement('div');
        playerEl.id = `netplay-client-${id}-player`;
        playerEl.classList.add('text-end', 'col-5', 'col-md-3');
        playerEl.innerText = _displayPlayer(player);
        containerEl.append(playerEl);

        document.getElementById('netplay-clients').append(el);
    }

    function clientDisconnected(id) {
        document.getElementById(`netplay-client-${id}`).remove();
    }

    function clientNameChanged(id, name) {
        document.getElementById(`netplay-client-${id}-name`).innerText = name;
    }

    function clientPlayerChanged(id, player) {
        document.getElementById(`netplay-client-${id}-player`).innerText = _displayPlayer(player);
    }

    ///////////////////////////////////////////////////////////////////////////
    // Errors
    ///////////////////////////////////////////////////////////////////////////

    function errorHandler(type, clientId, e) {
        if (['web-socket', 'rtc-offer-send', 'rtc-answer-send', 'rtc-connection', 'rtc-ice-connection', 'rtc-control-channel'].contains(type)) {
            errorScreen(true);
        }
        //TODO error toast
    }

    function errorScreen(display) {
        const el = document.getElementById('error-screen');
        if (display) {
            el.classList.remove('d-none');
        } else {
            el.classList.add('d-none');
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Utils
    ///////////////////////////////////////////////////////////////////////////

    function setupGameDisplaySize() {
        const game = document.getElementById('game');
        game.width = window.innerWidth;
        game.height = window.innerHeight - 50;
    }

    function _displayPlayer(player) {
        return player === -1 ? 'spectator' : player + 1;
    }

})();
