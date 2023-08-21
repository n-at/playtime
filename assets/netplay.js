(() => {

    // default NetplayClient configuration
    // noinspection JSUnusedLocalSymbols
    const defaultConfiguration = {
        //game <canvas> element, required for host
        gameCanvasEl: null,

        //game <video> element, required for client
        gameVideoEl: null,

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

        //when received video track from host
        onVideoTrackAdded: () => {},

        //when received audio track from host
        onAudioTrackAdded: () => {},

        //when creating offer for client failed
        onRTCOfferError: clientId => {},

        //when creating answer for client failed
        onRTCAnswerError: clientId => {},

        //when connection state changed
        //clientId - connected client
        //state - new, connecting, connected, disconnected, failed, closed
        onRTCConnectionStateChanged: (clientId, state) => {},

        //when signalling state changed
        //clientId - connected client
        //state - stable, have-local-offer, have-remote-offer, have-local-pranswer, have-remote-pranswer, closed
        onRTCSignallingStateChanged: (clientId, state) => {},

        //when ICE state changed
        //clientId - connected client
        //state - new, checking, connected, completed, failed, disconnected, closed
        onRTCIceStateChanged: (clientId, state) => {},

        //when ICE gathering state changed
        //clientId - connected client
        //state - new, gathering, complete
        onRTCIceGatheringStateChanged: (clientId, state) => {},

        //when ICE error
        onRTCIceError: clientId => {},

        //when control data channel opened
        onRTCControlChannelOpen: clientId => {},

        //when control data channel error
        onRTCControlChannelError: clientId => {},

        //when controller button pressed
        onRTCControlChannelInput: (clientId, player, control) => {},

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
            //configuration
            configuration,

            //connections
            ws: null,
            rtcHost: null,
            rtcClients: {},
            rtcHostControlChannel: {},
            rtcControlChannels: {},

            //client data
            clientId: null,
            clientKey: null,
            host: configuration.host,
            hostId: null,
            name: null,
            player: null,
            clients: {},

            //media stream tracks (host)
            videoTrack: null,
            audioTrack: null,

            //instance methods
            connect,
            getClientId,
            getClientKey,
            getHostId,
            getName,
            getPlayer,
            getClients,
            setName,
            setClientPlayer,
            sendControlInput,
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
        const clientId = message.from_id;
        const sdp = JSON.parse(message.sdp);

        if (client.rtcClients[clientId]) {
            rtcSendAnswer(client, clientId, client.rtcClients[clientId], sdp);
        }
        if (client.hostId === clientId && client.rtcHost) {
            rtcSendAnswer(client, clientId, client.rtcHost, sdp);
        }
    }

    function wsMessageSignallingAnswer(client, message) {
        const clientId = message.from_id;
        const sdp = JSON.parse(message.sdp);

        if (client.rtcClients[clientId]) {
            rtcHandleAnswer(client, clientId, client.rtcClients[clientId], sdp);
        }
        if (client.hostId === clientId && client.rtcHost) {
            rtcHandleAnswer(client, clientId, client.rtcHost, sdp);
        }
    }

    function wsMessageSignallingIceCandidate(client, message) {
        const clientId = message.from_id;
        const sdp = JSON.parse(message.sdp);

        if (client.rtcClients[clientId]) {
            rtcHandleIceCandidate(client, clientId, client.rtcClients[clientId], sdp);
        }
        if (client.hostId === clientId && client.rtcHost) {
            rtcHandleIceCandidate(client, clientId, client.rtcHost, sdp);
        }
    }

    ///////////////////////////////////////////////////////////////////////////

    function clientConnected(client, connectedClientId, clientData) {
        client.clients[connectedClientId] = clientData;
        client.configuration.onClientConnected(connectedClientId, clientData.name, clientData.player);

        if (client.clientId === connectedClientId) {
            return;
        }

        if (client.host && !client.rtcClients[connectedClientId]) {
            client.rtcClients[connectedClientId] = connectRTC(client, connectedClientId);
        }
        if (!client.host && !client.rtcHost) {
            client.rtcHost = connectRTC(client, null);
        }
    }

    function clientDisconnected(client, disconnectedClientId) {
        delete client.clients[disconnectedClientId];
        client.configuration.onClientDisconnected(disconnectedClientId);

        if (client.clientId === disconnectedClientId) {
            return;
        }

        if (client.rtcClients[disconnectedClientId]) {
            client.rtcClients[disconnectedClientId].close();
            delete client.rtcClients[disconnectedClientId];
        }
        if (client.rtcControlChannels[disconnectedClientId]) {
            client.rtcControlChannels[disconnectedClientId].close();
            delete client.rtcControlChannels[disconnectedClientId];
        }

        if (client.host === disconnectedClientId && client.rtcHost) {
            client.rtcHost.close();
            client.rtcHost = null;
        }
        if (client.host === disconnectedClientId && client.rtcHostControlChannel) {
            client.rtcHostControlChannel.close();
            client.rtcHostControlChannel = null;
        }
    }

    function connectRTC(client, destinationClientId) {
        const connection = new RTCPeerConnection({
            iceServers: [
                _buildIceServerConfiguration(
                    client.configuration.turnServerUrl,
                    client.configuration.turnServerUser,
                    client.configuration.turnServerPassword
                ),
            ],
        });
        connection.addEventListener('connectionstatechange', () => {
            _debug(client, 'RTC connection state changed', destinationClientId, connection.connectionState);
           client.configuration.onRTCConnectionStateChanged(destinationClientId, connection.connectionState);
        });
        connection.addEventListener('datachannel', e => {
            rtcClientControlDataChannel(client, destinationClientId, e.channel);
        });
        connection.addEventListener('icecandidate', e => {
            rtcSendIceCandidate(client, destinationClientId, connection, e);
        });
        connection.addEventListener('icecandidateerror', e => {
           console.error(`RTC ICE candidate for ${destinationClientId} error`, e);
           client.configuration.onRTCIceError(destinationClientId);
        });
        connection.addEventListener('iceconnectionstatechange', () => {
            _debug(client, 'RTC ICE connection state changed', destinationClientId, connection.iceConnectionState);
            client.configuration.onRTCIceStateChanged(destinationClientId, connection.iceConnectionState);
        });
        connection.addEventListener('icegatheringstatechange', () => {
            _debug(client, 'RTC ICE gathering state changed', destinationClientId, connection.iceGatheringState);
            client.configuration.onRTCIceGatheringStateChanged(destinationClientId, connection.iceConnectionState);
        });
        connection.addEventListener('negotiationneeded', () => {
           rtcSendOffer(client, destinationClientId, connection);
        });
        connection.addEventListener('signalingstatechange', () => {
            _debug(client, 'RTC signalling state changed', destinationClientId, connection.signalingState);
           client.configuration.onRTCSignallingStateChanged(destinationClientId, connection.signalingState);
        });
        connection.addEventListener('track', e => {
            rtcTrack(client, e);
        });

        if (client.host) {
            const mediaStream = new MediaStream();
            collectMediaTracks(client).forEach(track => mediaStream.addTrack(track));
            mediaStream.getTracks().forEach(track => connection.addTrack(track, mediaStream));
        }
        if (!client.host) {
            rtcHostControlDataChannel(client, connection.createDataChannel('controls'));
        }

        return connection;
    }

    function collectMediaTracks(client) {
        if (!client.videoTrack || client.videoTrack.readyState !== 'live') {
            const videoTracks = client.configuration.gameCanvasEl.captureStream().getVideoTracks();
            if (videoTracks.length !== 0) {
                client.videoTrack = videoTracks[0];
            } else {
                client.videoTrack = null;
                console.error('Unable to capture video stream');
            }
        }

        if (!client.audioTrack || client.audioTrack.readyState !== 'live') {
            if (window.AL && window.AL.currentCtx && window.AL.currentCtx.audioCtx) {
                const destination = window.AL.currentCtx.audioCtx.createMediaStreamDestination();
                const audioTracks = destination.stream.getAudioTracks();
                if (audioTracks.length !== 0) {
                    client.audioTrack = audioTracks[0];
                } else {
                    client.audioTrack = null;
                    console.error('Unable to capture audio stream');
                }
            } else {
                console.error('Unable to capture audio stream (AL not available)');
                client.audioTrack = null;
            }
        }

        const tracks = [];
        if (client.videoTrack && client.videoTrack.readyState === 'live') {
            tracks.push(client.videoTrack);
        }
        if (client.audioTrack && client.audioTrack.readyState === 'live') {
            tracks.push(client.audioTrack);
        }
        return tracks;
    }

    /**
     * @param client Object
     * @param destinationClientId string
     * @param connection RTCPeerConnection
     */
    function rtcSendOffer(client, destinationClientId, connection) {
        if (!client.host && destinationClientId === null) {
            destinationClientId = client.hostId;
        }
        connection
            .createOffer()
            .then(offer => connection.setLocalDescription(offer))
            .then(() => {
                const message = _messageSignalling(MessageTypeSignallingOffer, destinationClientId, connection.localDescription);
                wsSend(client, message);
            })
            .catch(e => {
                console.error(`RTC create offer for ${destinationClientId} error`, e);
                client.configuration.onRTCOfferError(destinationClientId);
            });
    }

    /**
     * @param client Object
     * @param destinationClientId string
     * @param connection RTCPeerConnection
     * @param sdp string
     */
    function rtcSendAnswer(client, destinationClientId, connection, sdp) {
        const description = new RTCSessionDescription(sdp);

        connection
            .setRemoteDescription(description)
            .then(() => connection.createAnswer())
            .then(answer => connection.setLocalDescription(answer))
            .then(() => {
                const message = _messageSignalling(MessageTypeSignallingAnswer, destinationClientId, connection.localDescription);
                wsSend(client, message);
            })
            .catch(e => {
                console.error(`RTC create answer for ${destinationClientId} error`, e);
                client.configuration.onRTCAnswerError(destinationClientId);
            });
    }

    /**
     * @param client Object
     * @param destinationClientId string
     * @param connection RTCPeerConnection
     * @param sdp string
     */
    function rtcHandleAnswer(client, destinationClientId, connection, sdp) {
        const description = new RTCSessionDescription(sdp);

        connection
            .setRemoteDescription(description)
            .catch(e => {
                console.error(`RTC handle answer from ${destinationClientId} error`, e);
            });
    }

    /**
     * @param client Object
     * @param destinationClientId string
     * @param connection RTCPeerConnection
     * @param e RTCPeerConnectionIceEvent
     */
    function rtcSendIceCandidate(client, destinationClientId, connection, e) {
        if (!e.candidate) {
            return;
        }

        const message = _messageSignalling(MessageTypeSignallingIceCandidate, destinationClientId, e.candidate);
        wsSend(client, message);
    }

    /**
     * @param client Object
     * @param destinationClientId string
     * @param connection RTCPeerConnection
     * @param sdp string
     */
    function rtcHandleIceCandidate(client, destinationClientId, connection, sdp) {
        const candidate = new RTCIceCandidate(sdp);

        connection
            .addIceCandidate(candidate)
            .catch(e => {
                console.error(`RTC handle ICE candidate from ${destinationClientId} error`, e);
                client.configuration.onRTCIceError(destinationClientId);
            });
    }

    /**
     * @param client Object
     * @param e RTCTrackEvent
     */
    function rtcTrack(client, e) {
        if (client.host) {
            return;
        }
        if (!e.streams || e.streams.length === 0) {
            console.error('No media streams received');
            return;
        }

        client.configuration.gameVideoEl.srcObject = e.streams[0];

        if (e.track.kind === 'video') {
            _debug(client, 'RTC video track received');
            client.configuration.onVideoTrackAdded();
        } else if (e.track.kind === 'audio') {
            _debug(client, 'RTC audio track received');
            client.configuration.onAudioTrackAdded();
        }
    }

    /**
     * @param client Object
     * @param dataChannel RTCDataChannel
     */
    function rtcHostControlDataChannel(client, dataChannel) {
        client.rtcHostControlChannel = dataChannel;

        dataChannel.addEventListener('open', () => {
            client.configuration.onRTCControlChannelOpen(client.hostId);
        });

        //host doesn't send any data to clients (yet)

        dataChannel.addEventListener('error', e => {
            console.error('RTC host control data channel error', client.hostId, e);
            client.configuration.onRTCControlChannelError(client.hostId);
        });
    }

    /**
     * @param client Object
     * @param destinationClientId string
     * @param dataChannel RTCDataChannel
     */
    function rtcClientControlDataChannel(client, destinationClientId, dataChannel) {
        client.rtcControlChannels[destinationClientId] = dataChannel;

        dataChannel.addEventListener('open', () => {
            client.configuration.onRTCControlChannelOpen(destinationClientId);
        });

        dataChannel.addEventListener('error', e => {
            console.error('RTC client control data channel error', destinationClientId, e);
            client.configuration.onRTCControlChannelError(destinationClientId);
        });

        dataChannel.addEventListener('message', e => {
            if (!client.clients[destinationClientId]) {
                return;
            }

            const destinationClient = client.clients[destinationClientId];
            if (destinationClient.player === -1) {
                return;
            }

            const input = JSON.parse(e.data);
            if (!input) {
                return;
            }

            _debug(client, 'RTC DC control', destinationClientId, input);

            client.configuration.onRTCControlChannelInput(destinationClientId, destinationClient.player, input.code);
        });
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
        if (!this.host) {
            console.error('setClientPlayer available only for host client');
            return;
        }
        wsSend(this, _messagePlayerChange(clientId, player));
    }

    function sendControlInput(inputCode) {
        if (!this.rtcHostControlChannel || this.rtcHostControlChannel.readyState !== 'open') {
            return;
        }
        this.rtcHostControlChannel.send(JSON.stringify({code: inputCode}));
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
        if (config.host && !config.gameCanvasEl) {
            throw new Error('gameCanvasEl required');
        }
        if (!config.host && !config.gameVideoEl) {
            throw new Error('gameAudioEl required');
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

    function _buildIceServerConfiguration(url, user, password) {
        const ice = {
            urls: url,
        };
        if (user) {
            ice.username = user;
        }
        if (password) {
            ice.credential = password;
        }
        return ice;
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
                sdp: JSON.stringify(sdp),
            },
        }
    }

})();
