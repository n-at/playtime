(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        document.getElementById('netplay-url').value = _buildNetplayUrl();
        document.getElementById('netplay-url-copy').addEventListener('click', _copyNetplayUrl);

        _renderJoinQR();
    });

    window.EJS_onGameStart = () => {
        netplay = NetplayClient({
            debug: true,
            gameCanvasEl: document.querySelector('#game canvas'),
            gameId: NetplayGameId,
            sessionId: NetplaySessionId,
            host: true,
            turnServerUrl: NetplayTurnServerUrl,
            turnServerUser: NetplayTurnServerUser,
            turnServerPassword: NetplayTurnServerPassword,

            onClientError: errorHandler,
        });
        netplay.connect();
    };

    ///////////////////////////////////////////////////////////////////////////
    // Self name and player
    ///////////////////////////////////////////////////////////////////////////

    //TODO

    ///////////////////////////////////////////////////////////////////////////
    // Game session clients list
    ///////////////////////////////////////////////////////////////////////////

    //TODO

    ///////////////////////////////////////////////////////////////////////////
    // Connection status
    ///////////////////////////////////////////////////////////////////////////

    function errorHandler(type, clientId, e) {
        //TODO
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
