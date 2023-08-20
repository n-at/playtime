(() => {

    // default NetplayClient configuration
    // noinspection JSUnusedLocalSymbols
    const defaultConfiguration = {
        //game canvas element (for host - canvas of EmulatorJS, for client - display canvas)
        gameCanvasEl: null,

        //game id, required
        gameId: null,

        //current game session id, required
        sessionId: null,

        //is current client host for the game
        host: false,

        //URL of TURN/STUN/ICE server, required
        turnServerUrl: null,

        //login for TURN/STUN/ICE server (if required)
        turnServerUser: null,

        //password for TURN/STUN/ICE server (if required)
        turnServerPassword: null,

        //enable debug output
        debug: false,

        //handlers

        //when WebSocket is connected
        onWSConnected: () => {},

        //when WebSocket is disconnected
        onWSDisconnected: () => {},

        //when WebSocket encountered error
        //e Event https://developer.mozilla.org/en-US/docs/Web/API/Event
        onWSError: e => {},

        //when changed client self name
        onSelfNameChanged: name => {},

        //when changed client self player
        onSelfPlayerChanged: player => {},

        //when client (including self) is connected
        onClientConnected: (clientId, name, player) => {},

        //when client (including self) is disconnected
        onClientDisconnected: clientId => {},

        //when server sent error message (excluding WebSocket connection error)
        onClientError: message => {},

        //when client (including self) changed name
        onClientNameChanged: (clientId, name) => {},

        //when client (including self) changed player
        onClientPlayerChanged: (clientId, player) => {},
    };

    ///////////////////////////////////////////////////////////////////////////

    window.NetplayClient = function(config) {
        const configuration = Object.assign({}, defaultConfiguration, config);
        _validateConfiguration(configuration);

        return {
            configuration,

            ws: null,
            rtcHost: null,
            rtcClients: {},

            clientId: null,
            clientKey: null,
            hostId: null,
            name: null,
            player: null,
            clients: {},

            connect,
            getClientId,
            getClientKey,
            getHostId,
            getName,
            getPlayer,
            getClients,
            setName,
            setClientPlayer,
        };
    };

    ///////////////////////////////////////////////////////////////////////////

    const MessageTypeGreeting = 'greeting';
    const MessageTypeConnected = 'connected';
    const MessageTypeDisconnected = 'disconnected';
    const MessageTypeHeartbeat = 'heartbeat';
    const MessageTypeError = 'error';
    const MessageTypePlayerChange = 'player-change';
    const MessageTypePlayerChanged = 'player-changed';
    const MessageTypeClientNameChange = 'client-name-change';
    const MessageTypeClientNameChanged = 'client-name-changed';
    const MessageTypeSignallingOffer = 'signalling-offer';
    const MessageTypeSignallingAnswer = 'signalling-answer';
    const MessageTypeSignallingIceCandidate = 'signalling-ice-candidate';

    function connect() {
        const url = _buildWebSocketUrl(this.configuration.gameId, this.configuration.sessionId);
        this.ws = new WebSocket(url);
        this.ws.addEventListener('open', () => {
            _debug(this, 'WebSocket connected');
            this.configuration.onWSConnected();
        });
        this.ws.addEventListener('close', () => {
            _debug(this, 'WebSocket disconnected');
            this.configuration.onWSDisconnected();
        });
        this.ws.addEventListener('error', e => {
            _debug(this, 'WebSocket error', e);
            this.configuration.onWSError(e);
        });
        this.ws.addEventListener('message', e => {
            const message = JSON.parse(e.data);

            if (!message.type) {
                console.error('empty message type');
                return;
            }

            _debug(this, 'WebSocket incoming message', message.type, message);

            switch (message.type) {
                case MessageTypeGreeting:
                    wsMessageGreeting(this, message.greeting);
                    break;
                case MessageTypeConnected:
                    wsMessageConnected(this, message.connected);
                    break;
                case MessageTypeDisconnected:
                    wsMessageDisconnected(this, message.disconnected);
                    break;
                case MessageTypeHeartbeat:
                    wsMessageHeartbeat(this);
                    break;
                case MessageTypeError:
                    wsMessageError(this, message.error);
                    break;
                case MessageTypeClientNameChanged:
                    wsMessageNameChanged(this, message.name_changed);
                    break;
                case MessageTypePlayerChanged:
                    wsMessagePlayerChanged(this, message.player_changed);
                    break;
                case MessageTypeSignallingOffer:
                    wsMessageSignallingOffer(this, message.signalling);
                    break;
                case MessageTypeSignallingAnswer:
                    wsMessageSignallingAnswer(this, message.signalling);
                    break;
                case MessageTypeSignallingIceCandidate:
                    wsMessageSignallingIceCandidate(this, message.signalling);
                    break;
                default:
                    console.error(`unknown message type: ${message.type}`);
            }
        });
    }

    function wsSend(client, message) {
        if (!client || !client.ws) {
            return;
        }

        _debug(client, 'WebSocket send message', message.type, message);

        client.ws.send(JSON.stringify(message));
    }

    function wsMessageGreeting(client, message) {
        client.hostId = message.host_id;
        client.clientId = message.client_id;
        client.clientKey = message.client_key;
        client.clients = {};

        for (let connectedClientIdx in message.clients) {
            const connectedClient = message.clients[connectedClientIdx];
            clientConnected(client, connectedClient.client_id, connectedClient);
        }

        client.name = message.name;
        client.configuration.onSelfNameChanged(client.name);

        client.player = message.player;
        client.configuration.onSelfPlayerChanged(client.player);
    }

    function wsMessageConnected(client, message) {
        clientConnected(client, message.client_id, message);
    }

    function wsMessageDisconnected(client, message) {
        clientDisconnected(client, message.client_id);
    }

    function wsMessageHeartbeat(client) {
        wsSend(client, _messageHeartbeat());
    }

    function wsMessageError(client, message) {
        client.configuration.onClientError(message.message);
    }

    function wsMessageNameChanged(client, message) {
        const clientId = message.client_id;
        const name = message.name;

        if (!client.clients[clientId]) {
            return;
        }

        client.clients[clientId].name = name;
        client.configuration.onClientNameChanged(clientId, name);

        if (client.clientId === clientId) {
            client.name = name;
            client.configuration.onSelfNameChanged(name);
        }
    }

    function wsMessagePlayerChanged(client, message) {
        const clientId = message.clientId;
        const player = message.player;

        if (!client.clients[clientId]) {
            return;
        }

        client.clients[clientId].player = player;
        client.configuration.onClientPlayerChanged(clientId, player);

        if (client.clientId === clientId) {
            client.player = player;
            client.configuration.onSelfPlayerChanged(player);
        }
    }

    function wsMessageSignallingOffer(client, message) {
        //TODO RTC
    }

    function wsMessageSignallingAnswer(client, message) {
        //TODO RTC
    }

    function wsMessageSignallingIceCandidate(client, message) {
        //TODO RTC
    }

    ///////////////////////////////////////////////////////////////////////////

    function clientConnected(client, connectedClientId, clientData) {
        client.clients[connectedClientId] = clientData;
        client.configuration.onClientConnected(connectedClientId, clientData.name, clientData.player);

        if (client.clientId === connectedClientId) {
            return;
        }

        if (client.configuration.host && !client.rtcClients[connectedClientId]) {
            client.rtcClients[connectedClientId] = connectRTC(client, connectedClientId);
        }
        if (!client.configuration.host && !client.rtcHost) {
            client.rtcHost = connectRTC(client, null);
        }
    }

    function clientDisconnected(client, disconnectedClientId) {
        delete client.clients[disconnectedClientId];
        client.configuration.onClientDisconnected(disconnectedClientId);

        if (client.clientId === disconnectedClientId) {
            return;
        }

        if (client.configuration.host && client.rtcClients[disconnectedClientId]) {
            client.rtcClients[disconnectedClientId].close();
            delete client.rtcClients[disconnectedClientId];
        }
        if (!client.configuration.host && client.rtcHost) {
            client.rtcHost.close();
            client.rtcHost = null;
        }
    }

    ///////////////////////////////////////////////////////////////////////////

    function connectRTC(client, clientId) {
        //TODO RTC
        return null;
    }

    function connectRTCControl(client) {
        //TODO RTC
    }

    ///////////////////////////////////////////////////////////////////////////

    function getClientId() {
        return this.clientId;
    }

    function getClientKey() {
        return this.clientKey;
    }

    function getHostId() {
        return this.hostId;
    }

    function getName() {
        return this.name;
    }

    function getPlayer() {
        return this.player;
    }

    function getClients() {
        const clients = [];

        for (let clientId in this.clients) {
            clients.push(this.clients[clientId]);
        }

        return clients;
    }

    function setName(name) {
        wsSend(this, _messageNameChange(name));
    }

    function setClientPlayer(clientId, player) {
        if (!this.configuration.host) {
            console.error('setClientPlayer available only for host client');
            return;
        }
        wsSend(this, _messagePlayerChange(clientId, player));
    }

    ///////////////////////////////////////////////////////////////////////////

    function _debug(client, ...args) {
        if (!client.configuration.debug) {
            return;
        }
        if (console.debug) {
            console.debug(...args);
        } else {
            console.log(...args);
        }
    }

    function _validateConfiguration(config) {
        if (!config.gameCanvasEl) {
            throw new Error('gameCanvasEl required');
        }
        if (!config.gameId || typeof config.gameId != 'string') {
            throw new Error('gameId required');
        }
        if (!config.sessionId || typeof config.sessionId != 'string') {
            throw new Error('sessionId required');
        }
        if (!config.turnServerUrl || typeof config.turnServerUrl != 'string') {
            throw new Error('turnServerUrl required');
        }
    }

    function _buildWebSocketUrl(gameId, sessionId) {
        const protocol = location.protocol === 'https:' ? 'wss' : 'ws';
        return `${protocol}://${location.host}/netplay/${gameId}/${sessionId}/ws`;
    }

    function _messageNameChange(name) {
        return {
            type: MessageTypeClientNameChange,
            name_change: {
                name: name,
            },
        };
    }

    function _messagePlayerChange(clientId, player) {
        return {
            type: MessageTypePlayerChange,
            player_change: {
                client_id: clientId,
                player: player,
            },
        };
    }

    function _messageHeartbeat() {
        return {
            type: MessageTypeHeartbeat,
        };
    }

    function _messageSignalling(type, clientId, sdp) {
        return {
            type: type,
            signalling: {
                client_id: clientId,
                sdp: sdp,
            },
        }
    }

})();
