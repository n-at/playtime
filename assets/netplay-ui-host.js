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

        //host can play as any player, auto set 1
        netplay.setClientPlayer(netplay.getClientId(), 1);

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

    function _displayPlayer(player) {
        return player === -1 ? 'spectator' : player + 1;
    }

})();
