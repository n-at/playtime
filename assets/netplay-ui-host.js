(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        document.getElementById('netplay-url').value = _buildNetplayUrl();
        document.getElementById('netplay-url-copy').addEventListener('click', _copyNetplayUrl);
        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);

        _renderJoinQR();
    });

    window.EJS_onGameStart = () => {
        netplay = NetplayClient({
            debug: window.NetplayDebug,
            gameCanvasEl: document.querySelector('#game canvas'),
            gameId: NetplayGameId,
            sessionId: NetplaySessionId,
            host: true,
            turnServerUrl: NetplayTurnServerUrl,
            turnServerUser: NetplayTurnServerUser,
            turnServerPassword: NetplayTurnServerPassword,

            onClientError: errorHandler,
            onGreeting: wsGreeting,
            onSelfNameChanged: selfNameChanged,
            onClientConnected: clientConnected,
            onClientDisconnected: clientDisconnected,
            onClientNameChanged: clientNameChanged,
            onClientPlayerChanged: clientPlayerChanged,
        });
        netplay.connect();
    };

    ///////////////////////////////////////////////////////////////////////////
    // Self name and player
    ///////////////////////////////////////////////////////////////////////////

    function wsGreeting() {
        const netplayName = netplay.getName();
        const savedName = window.NetplayLoadClientName(netplayName);
        if (netplayName !== savedName) {
            netplay.setName(savedName);
        }

        //host can play as any player, auto set 0
        netplay.setClientPlayer(netplay.getClientId(), 0);

        window.ShowToastMessage('success', 'Connected to server');
    }

    function selfNameChanged(name) {
        document.getElementById('netplay-name').value = name;
    }

    function changeSelfName() {
        const name = document.getElementById('netplay-name').value;

        if (name.trim().length === 0 || name.length > 32) {
            window.FlashButtonIcon('netplay-name-change', ['btn-outline-secondary'], ['bi-check'], ['btn-outline-danger'], ['bi-x']);
            return;
        }

        netplay.setName(name);
        window.NetplaySaveClientName(name);

        window.FlashButtonIcon('netplay-name-change', ['btn-outline-secondary'], [], ['btn-outline-success'], []);
    }

    ///////////////////////////////////////////////////////////////////////////
    // Game session clients list
    ///////////////////////////////////////////////////////////////////////////

    function clientConnected(id, name, player) {
        if (id !== netplay.getClientId()) {
            window.ShowToastMessage('success', `${name} (${NetplayPlayerDisplay(id, netplay.getHostId(), player)}) connected`);
        }

        const elId = `netplay-client-${id}`;
        if (document.getElementById(elId)) {
            clientNameChanged(name);
            clientPlayerChanged(player);
            return;
        }

        const containerEl = document.createElement('div');
        containerEl.id = `netplay-client-${id}`;
        containerEl.classList.add('list-group-item');

        const rowEl = document.createElement('div');
        rowEl.classList.add('row');
        containerEl.append(rowEl);

        const nameEl = _createClientNameEl(id, name);
        rowEl.append(nameEl);

        const playerEl = _createClientPlayerEl(id);
        rowEl.append(playerEl);

        if (id !== netplay.getClientId()) {
            playerEl.append(_createClientPlayerSelect(id, player));
        } else {
            playerEl.innerText = NetplayPlayerDisplay(id, netplay.getHostId(), player);
        }

        document.getElementById('netplay-clients').append(containerEl);
    }

    function clientDisconnected(id) {
        const client = netplay.getClient(id);
        if (client && id !== netplay.getClientId()) {
            window.ShowToastMessage('warning', `${client.name} (${NetplayPlayerDisplay(id, netplay.getHostId(), client.player)}) disconnected`);
        }

        const el = document.getElementById(`netplay-client-${id}`);
        if (el) {
            el.remove();
        }
    }

    function clientNameChanged(id, newName) {
        const client = netplay.getClient(id);
        if (client && id !== netplay.getClientId()) {
            window.ShowToastMessage('info', `${client.name} (${NetplayPlayerDisplay(id, netplay.getHostId(), client.player)}) is now ${name}`);
        }

        const el = document.getElementById(`netplay-client-${id}-name`);
        if (el) {
            el.innerText = newName;
        }
    }

    function clientPlayerChanged(id, newPlayer) {
        const el = document.getElementById(`netplay-client-${id}-player`);
        if (el) {
            el.value = newPlayer.toString();
        }
    }

    function _createClientNameEl(id, name) {
        const el = document.createElement('div');
        el.id = `netplay-client-${id}-name`;
        el.classList.add('lead', 'col-6', 'col-md-9');
        el.innerText = name;
        return el;
    }

    function _createClientPlayerEl(id) {
        const el = document.createElement('div');
        el.id = `netplay-client-${id}-player`;
        el.classList.add('text-end', 'col-6', 'col-md-3');
        return el;
    }

    function _createClientPlayerSelect(id, player) {
        const el = document.createElement('select');
        el.id = `netplay-client-${id}-player`;
        el.classList.add('form-select');

         [-1, 0, 1, 2, 3].forEach(playerId => {
            const playerTitle = window.NetplayPlayerDisplay(id, null, playerId);
            const option = document.createElement('option');
            option.value = playerId.toString();
            option.innerText = playerTitle;
            option.selected = (playerId === player);
            el.append(option);
         });

         el.addEventListener('change', () => {
            const newPlayer = parseInt(el.value);

            const client = netplay.getClient(id);
            if (!client) {
                return;
            }
            if (client.player === newPlayer) {
                return;
            }

            netplay.getClients().forEach(client => {
                if (newPlayer !== -1 && client.player === newPlayer) {
                    netplay.setClientPlayer(client.client_id, -1);
                }
            });

            netplay.setClientPlayer(id, newPlayer);
         });

        return el;
    }

    ///////////////////////////////////////////////////////////////////////////
    // Client controls
    ///////////////////////////////////////////////////////////////////////////

    //TODO

    ///////////////////////////////////////////////////////////////////////////
    // Connection status
    ///////////////////////////////////////////////////////////////////////////

    function errorHandler(type, clientId, e) {
        const client = netplay.getClient(clientId);
        const clientName = client ? client.name : 'unknown client';

        switch (type) {
            case 'web-socket':
                window.ShowToastMessage('danger', 'Server connection error');
                break;
            case 'rtc-offer-send':
            case 'rtc-answer-send':
            case 'rtc-ice-connection':
            case 'rtc-control-channel':
                window.ShowToastMessage('danger', `${clientName} connection error`);
                break;
            case 'rtc-answer-receive':
            case 'rtc-ice-candidate-accept':
                window.ShowToastMessage('warning', `${clientName} connection warning`);
                break;
            case 'rtc-connection':
                window.ShowToastMessage('danger', `${clientName} connection lost`);
                break;
        }

        console.error('error', type, clientId, e);
    }

    ///////////////////////////////////////////////////////////////////////////
    // Utils
    ///////////////////////////////////////////////////////////////////////////

    function _renderJoinQR() {
        const url = document.getElementById('netplay-url').value;
        new AwesomeQR.AwesomeQR({
            text: url,
            size: 250,
            margin: 5,
            colorDark: '#0d6efd',
            colorLight: '#fff',
            components: {
                data: { scale: 0.5 },
                timing: { scale: 0.5, protectors: false },
                alignment: { scale: 0.5, protectors: false },
                cornerAlignment: { scale: 0.5, protectors: true },
            }
        }).draw().then(dataUrl => {
            document.getElementById('netplay-qr').src = dataUrl;
        });
    }

    function _buildNetplayUrl() {
        return `${location.protocol}//${location.host}/netplay/${window.NetplayGameId}/${window.NetplaySessionId}`;
    }

    async function _copyNetplayUrl() {
        if (!navigator.clipboard) {
            return;
        }
        const url = document.getElementById('netplay-url').value;
        await navigator.clipboard.writeText(url);
        window.FlashButtonIcon('netplay-url-copy', ['btn-outline-secondary'], ['bi-clipboard'], ['btn-outline-success'], ['bi-clipboard-check']);
    }

})();
