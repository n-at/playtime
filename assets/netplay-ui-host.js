(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        document.getElementById('netplay-url').value = _buildNetplayUrl();
        document.getElementById('netplay-url-copy').addEventListener('click', _copyNetplayUrl);
        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);
        document.getElementById('netplay-quality').addEventListener('change', setQuality);

        _renderJoinQR();

        setInterval(clientCollectStats, 1000);
    });

    window.EJS_onGameStart = () => {
        netplay = NetplayClient({
            debug: window.NetplayDebug,
            gameCanvasEl: document.querySelector(`${EJS_player} canvas`),
            gameId: NetplayGameId,
            sessionId: NetplaySessionId,
            host: true,
            turnServerUrl: NetplayTurnServerUrl,
            turnServerUser: NetplayTurnServerUser,
            turnServerPassword: NetplayTurnServerPassword,

            onGreeting: wsGreeting,
            onSelfNameChanged: selfNameChanged,

            onClientCleanState: clientReset,
            onClientConnected: clientConnected,
            onClientDisconnected: clientDisconnected,
            onClientNameChanged: clientNameChanged,
            onClientPlayerChanged: clientPlayerChanged,

            onRTCControlChannelInput: controlInput,

            onClientError: errorHandler,
            onWSReconnecting: wsReconnecting,
            onRTCReconnecting: rtcReconnecting,
            onRTCControlChannelReconnecting: rtcControlReconnecting,
        });
        setTimeout(() => netplay.connect(), 1500);
    };

    ///////////////////////////////////////////////////////////////////////////
    // Self name and player
    ///////////////////////////////////////////////////////////////////////////

    function wsGreeting() {
        const netplayName = netplay.getName();
        const savedName = window.NetplayLoadClientName(netplayName);
        if (savedName && netplayName !== savedName) {
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

    const clientBytesSent = {};
    const clientBytesReceived = {};

    function clientReset() {
        document.getElementById('netplay-clients').innerHtml = '';
    }

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
            playerEl.classList.remove('text-end');
            playerEl.classList.add('text-center');
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

        controlUnpress(id, client.player);
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

    function clientPlayerChanged(id, newPlayer, oldPlayer) {
        const el = document.getElementById(`netplay-client-${id}-player-select`);
        if (el) {
            el.value = newPlayer.toString();
        }
        controlUnpress(id, oldPlayer);
    }

    function clientCollectStats() {
        if (!netplay) {
            return;
        }
        netplay
            .getClients()
            .forEach(client => {
                const id = client.client_id;
                if (id === netplay.getClientId()) {
                    return;
                }
                netplay
                    .getClientStats(id)
                    .then(stats => {
                        stats.forEach(report => {
                            if (report.type === 'transport') {
                                _clientTransportStats(id, report);
                            }
                        });
                    })
                    .catch(e => {
                        console.error(`client stats error`, client.client_id, e);
                    });
            });
    }

    /**
     * @param {string} clientId
     * @param {RTCTransportStats} stats
     * @private
     */
    function _clientTransportStats(clientId, stats) {
        const prevSent = clientBytesSent[clientId] ? clientBytesSent[clientId] : 0;
        const prevReceived = clientBytesReceived[clientId] ? clientBytesReceived[clientId] : 0;

        const speedUpEl = document.getElementById(`netplay-client-${clientId}-stat-speed-up`);
        if (speedUpEl) {
            speedUpEl.innerText = _formatSpeed(stats.bytesSent - prevSent);
        }

        const speedDownEl = document.getElementById(`netplay-client-${clientId}-stat-speed-down`);
        if (speedDownEl) {
            speedDownEl.innerText = _formatSpeed(stats.bytesReceived - prevReceived);
        }

        const totalUpEl = document.getElementById(`netplay-client-${clientId}-stat-total-up`);
        if (totalUpEl) {
            totalUpEl.innerText = _formatBytes(stats.bytesSent);
        }

        const totalDownEl = document.getElementById(`netplay-client-${clientId}-stat-total-down`);
        if (totalDownEl) {
            totalDownEl.innerText = _formatBytes(stats.bytesReceived);
        }

        clientBytesSent[clientId] = stats.bytesSent;
        clientBytesReceived[clientId] = stats.bytesReceived;
    }

    function _createClientNameEl(id, name) {
        const el = document.createElement('div');
        el.classList.add('col-6', 'col-md-9');

        const nameEl = document.createElement('div');
        nameEl.id = `netplay-client-${id}-name`;
        nameEl.classList.add('lead');
        nameEl.innerText = name;
        el.append(nameEl);

        if (id === netplay.getClientId()) {
            return el;
        }

        const statsEl = document.createElement('div');
        statsEl.classList.add('row');
        el.append(statsEl);

        statsEl.append(
            _createClientStatsEl(id, 'speed-up', 'bi-arrow-up', 'text-success', 'Upload speed'),
            _createClientStatsEl(id, 'speed-down', 'bi-arrow-down', 'text-danger', 'Download speed'),
            _createClientStatsEl(id, 'total-up', 'bi-arrow-up-square', 'text-success', 'Total data uploaded'),
            _createClientStatsEl(id, 'total-down', 'bi-arrow-down-square', 'text-danger', 'Total data downloaded'),
        );

        return el;
    }

    function _createClientStatsEl(id, stat, iconCls, iconStyleCls, title) {
        const containerEl = document.createElement('div')
        containerEl.classList.add('col-3');
        containerEl.title = title;

        const iconEl = document.createElement('i');
        iconEl.classList.add('bi', 'me-2', iconCls, iconStyleCls);
        containerEl.append(iconEl);

        const valueEl = document.createElement('span');
        valueEl.id = `netplay-client-${id}-stat-${stat}`;
        valueEl.innerText = '-';
        containerEl.append(valueEl);

        return containerEl
    }

    function _createClientPlayerEl(id) {
        const el = document.createElement('div');
        el.id = `netplay-client-${id}-player`;
        el.classList.add('text-end', 'col-6', 'col-md-3');
        return el;
    }

    function _createClientPlayerSelect(id, player) {
        const el = document.createElement('select');
        el.id = `netplay-client-${id}-player-select`;
        el.classList.add('form-select');

         [-1, 0, 1, 2, 3].forEach(playerId => {
            const playerTitle = window.NetplayPlayerDisplay(id, null, playerId);
            const option = document.createElement('option');
            option.value = playerId.toString();
            option.innerText = playerTitle;
            el.append(option);
         });

         el.value = player.toString();

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

    const clientControlPressed = {};

    function controlInput(clientId, player, code, value) {
        const client = netplay.getClient(clientId);
        if (!client || client.player !== player) {
            return;
        }

        if (!clientControlPressed[clientId]) {
            clientControlPressed[clientId] = {};
        }
        clientControlPressed[clientId][code] = !!value;

        EJS_emulator.gameManager.simulateInput(player, code, value);
    }

    function controlUnpress(clientId, player) {
        if (!clientControlPressed[clientId]) {
            return;
        }

        const pressed = clientControlPressed[clientId];
        for (let code in pressed) {
            if (pressed[code]) {
                EJS_emulator.gameManager.simulateInput(player, code, 0);
            }
        }

        delete clientControlPressed[clientId];
    }

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
            case'server':
                window.ShowToastMessage('danger', `Server error: ${e.message}`);
                break;
        }

        console.error('error', type, clientId, e);
    }

    function wsReconnecting() {
        window.ShowToastMessage('warning', `Reconnecting to server...`);
    }

    function rtcReconnecting(clientId) {
        const client = netplay.getClient(clientId);
        const clientName = client ? client.name : 'unknown client';
        window.ShowToastMessage('warning', `Reconnecting to ${clientName}...`);
    }

    function rtcControlReconnecting(clientId) {
        const client = netplay.getClient(clientId);
        const clientName = client ? client.name : 'unknown client';
        window.ShowToastMessage('warning', `Reconnecting to ${clientName}...`);
    }

    function setQuality() {
        const el = document.getElementById('netplay-quality');
        if (netplay) {
            netplay.setQuality(el.value);
        }
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

    function _formatSpeed(value) {
        if (value < 1024) {
            return `${value} B/s`;
        }

        value /= 1024.0;
        if (value < 1024) {
            return `${Math.round(value * 10) / 10} kB/s`;
        }

        value /= 1024.0;
        return `${Math.round(value * 10) / 10} MB/s`;
    }

    function _formatBytes(value) {
        if (value < 1024) {
            return `${value} B`;
        }

        value /= 1024.0;
        if (value < 1024) {
            return `${Math.round(value * 10) / 10} kB`;
        }

        value /= 1024.0;
        if (value < 1024) {
            return `${Math.round(value * 10) / 10} MB`;
        }

        value /= 1024.0;
        return `${Math.round(value * 10) / 10} GB`;
    }

})();
