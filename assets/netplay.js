(() => {

    const MessageType = {
        Greeting: 'greeting',
        Connected: 'connected',
        Disconnected: 'disconnected',
        Heartbeat: 'heartbeat',
        Error: 'error',
        PlayerChange: 'player-change',
        PlayerChanged: 'player-changed',
        ClientNameChange: 'client-name-change',
        ClientNameChanged: 'client-name-changed',
        SignallingOffer: 'signalling-offer',
        SignallingAnswer: 'signalling-answer',
        SignallingIceCandidate: 'signalling-ice-candidate',
        DCInput: 'input',
        DCHeartbeat: 'heartbeat',
        DCPlayer: 'player',
    };

    const ClientErrorType = {
        WebSocket: 'web-socket',
        RtcLocalDescription: 'rtc-offer-send',
        RtcRemoteDescription: 'rtc-answer-send',
        RtcAnswerReceive: 'rtc-answer-receive',
        RtcConnection: 'rtc-connection',
        RtcIceCandidate: 'rtc-ice-candidate',
        RtcIceCandidateAccept: 'rtc-ice-candidate-accept',
        RtcIceConnection: 'rtc-ice-connection',
        RtcControlChannel: 'rtc-control-channel',
        Server: 'server',
        RetryLimit: 'retry-limit',
    };

    const TrackType = {
        Video: 'video',
        Audio: 'audio',
    };

    const ConnectionState = {
        Connecting: 'connecting',
        Connected: 'connected',
        Disconnected: 'disconnected',
    };

    const RetryTimeout = 1000;
    const RetryLimit = 10;

    ///////////////////////////////////////////////////////////////////////////
    // Default configuration
    ///////////////////////////////////////////////////////////////////////////

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

        //when any error occurred
        //type - value from ClientErrorType
        //clientId - associated client
        //e - Error https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error
        //    Event https://developer.mozilla.org/en-US/docs/Web/API/Event
        onClientError: (type, clientId, e) => {},

        //when client resets its state (e.g. after connection failure)
        onClientCleanState: () => {},

        //when WebSocket is connected
        onWSConnected: () => {},

        //when WebSocket is disconnected
        onWSDisconnected: () => {},

        //when WebSocket is disconnected and started new connect attempt
        onWSReconnecting: () => {},

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

        //when RTC is reconnecting
        onRTCReconnecting: clientId => {},

        //when control data channel opened
        onRTCControlChannelOpen: clientId => {},

        //when control data channel closed
        onRTCControlChannelClose: clientId => {},

        //when control data channel is reconnecting
        onRTCControlChannelReconnecting: clientId => {},

        //when controller button pressed
        onRTCControlChannelInput: (clientId, player, code, value) => {},

        //when received media track from host
        onRTCTrack: (type, tracks) => {},

        //when server sent greeting message
        onGreeting: () => {},

        //when changed client self name
        onSelfNameChanged: (newName, oldName) => {},

        //when changed client self player
        onSelfPlayerChanged: (newPlayer, oldPlayer) => {},

        //when client (including self) is connected
        onClientConnected: (clientId, name, player) => {},

        //when client (including self) is disconnected
        onClientDisconnected: clientId => {},

        //when client (including self) changed name
        onClientNameChanged: (clientId, newName, oldName) => {},

        //when client (including self) changed player
        onClientPlayerChanged: (clientId, newPlayer, oldPlayer) => {},
    };

    ///////////////////////////////////////////////////////////////////////////
    // Client definition
    ///////////////////////////////////////////////////////////////////////////

    window.NetplayClient = function(config) {
        const configuration = Object.assign({}, defaultConfiguration, config);
        _validateConfiguration(configuration);

        return {
            configuration,

            //connections
            connectionState: ConnectionState.Disconnected,
            retryConnection: false,
            retries: {},
            ws: null,
            rtcClients: {},
            rtcControlChannels: {},
            rtcMakingOffer: {},
            rtcIgnoreOffer: {},

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
            connect() {
                this.retryConnection = true;
                wsConnect(this);
            },

            disconnect() {
                this.retryConnection = false;
                this.connectionState = ConnectionState.Disconnected;
                cleanClientState(this);
            },

            getClientId() {
                return this.clientId;
            },

            getClientKey() {
                return this.clientKey;
            },

            getHostId() {
                return this.hostId;
            },

            getName() {
                return this.name;
            },

            getPlayer() {
                return this.player;
            },

            /**
             * @param {string} clientId
             * @returns {Object|null}
             */
            getClient(clientId) {
                if (this.clients[clientId]) {
                    return Object.assign({}, this.clients[clientId]);
                } else {
                    return null;
                }
            },

            /**
             * @returns {*[]}
             */
            getClients() {
                const clients = [];
                for (let clientId in this.clients) {
                    clients.push(Object.assign({}, this.clients[clientId]));
                }
                return clients;
            },

            /**
             * @param {string} clientId
             * @returns {Promise<RTCStatsReport>|Promise<never>}
             */
            getClientStats(clientId) {
                if (!this.rtcClients[clientId]) {
                    return Promise.reject(new Error('Client not found'));
                }
                return this.rtcClients[clientId].getStats();
            },

            /**
             * @param {string} name
             */
            setName(name) {
                wsSend(this, _messageNameChange(name));
            },

            /**
             * @param {string} clientId
             * @param {number} player
             */
            setClientPlayer(clientId, player) {
                if (!this.host) {
                    console.error('setClientPlayer available only for host client');
                    return;
                }
                wsSend(this, _messagePlayerChange(clientId, player));
            },

            /**
             * @param {number} inputCode
             * @param {number} inputValue
             */
            sendControlInput(inputCode, inputValue) {
                rtcDCSend(this, this.hostId, _messageDCInput(inputCode, inputValue));
            },

            sendControlHeartbeat() {
                rtcDCSend(this, this.hostId, _messageDCHeartbeat());
            },

            /**
             * @param {number} player
             */
            sendControlPlayer(player) {
                rtcDCSend(this, this.hostId, _messageDCPlayer(player));
            },
        };
    };

    ///////////////////////////////////////////////////////////////////////////
    // Client methods
    ///////////////////////////////////////////////////////////////////////////

    /**
     * @param {Object} client
     */
    function cleanClientState(client) {
        if (client.ws) {
            client.ws.close();
            client.ws = null;
        }

        for (let clientId in client.clients) {
            clientDisconnected(client, clientId);
        }

        client.clientId = null;
        client.clientKey = null;
        client.hostId = null;
        client.player = null;
        client.clients = {};
        client.rtcClients = {};
        client.rtcControlChannels = {};
        client.rtcMakingOffer = {};
        client.rtcIgnoreOffer = {};

        client.configuration.onClientCleanState();
    }

    /**
     * @param {Object} client
     * @param {string} disconnectedClientId
     */
    function clientDisconnected(client, disconnectedClientId) {
        client.configuration.onClientDisconnected(disconnectedClientId);
        delete client.clients[disconnectedClientId];
        delete client.retries[disconnectedClientId];
        delete client.rtcMakingOffer[disconnectedClientId];
        delete client.rtcIgnoreOffer[disconnectedClientId];

        if (client.rtcControlChannels[disconnectedClientId]) {
            const conn = client.rtcControlChannels[disconnectedClientId];
            delete client.rtcControlChannels[disconnectedClientId];
            conn.close();
        }

        if (client.rtcClients[disconnectedClientId]) {
            const conn = client.rtcClients[disconnectedClientId];
            delete client.rtcClients[disconnectedClientId];
            conn.close();
        }
    }

    /**
     * @param {Object} client
     * @param {string} connectedClientId
     * @param {Object} connectedClientData
     * @param {string} connectedClientData.client_id
     * @param {string} connectedClientData.name
     * @param {number} connectedClientData.player
     */
    function clientConnected(client, connectedClientId, connectedClientData) {
        client.clients[connectedClientId] = Object.assign({}, connectedClientData);
        client.configuration.onClientConnected(connectedClientId, connectedClientData.name, connectedClientData.player);

        if (client.clientId === connectedClientId) {
            return;
        }

        //host connecting to all clients (except self)
        //client connecting only to host

        if ((client.host && connectedClientId !== client.hostId) || (!client.host && connectedClientId === client.hostId)) {
            if (!client.rtcClients[connectedClientId] || !['new', 'connecting', 'connected'].includes(client.rtcClients[connectedClientId].connectionState)) {
                rtcConnect(client, connectedClientId);
            }
        }
    }

    /**
     * @param {Object} client
     * @returns {MediaStream}
     */
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
                const alContext = window.AL.currentCtx;
                const audioContext = alContext.audioCtx;

                const gainNodes = [];
                for (let sourceIdx in alContext.sources) {
                    gainNodes.push(alContext.sources[sourceIdx].gain);
                }

                const merger = audioContext.createChannelMerger(gainNodes.length);
                gainNodes.forEach(node => node.connect(merger));

                const destination = audioContext.createMediaStreamDestination();
                merger.connect(destination);

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

        const stream = new MediaStream();
        if (client.videoTrack && client.videoTrack.readyState === 'live') {
            stream.addTrack(client.videoTrack);
        }
        if (client.audioTrack && client.audioTrack.readyState === 'live') {
            stream.addTrack(client.audioTrack);
        }
        return stream;
    }

    ///////////////////////////////////////////////////////////////////////////
    // WebSocket
    ///////////////////////////////////////////////////////////////////////////

    /**
     * @param {Object} client
     */
    function wsConnect(client) {
        if (client.connectionState !== ConnectionState.Disconnected) {
            return;
        }

        client.connectionState = ConnectionState.Connecting;

        cleanClientState(client);

        const url = _buildWebSocketUrl(client.configuration.gameId, client.configuration.sessionId);
        client.ws = new WebSocket(url);
        client.ws.addEventListener('open', () => {
            _debug(client, 'WebSocket connected');
            client.connectionState = ConnectionState.Connected;
            client.configuration.onWSConnected();
            client.retries[null] = 0;
        });
        client.ws.addEventListener('close', () => {
            _debug(client, 'WebSocket disconnected');
            client.connectionState = ConnectionState.Disconnected;
            client.configuration.onWSDisconnected();
            if (client.retryConnection) {
                client.configuration.onWSReconnecting();
                _retry(client, null, () => wsConnect(client));
            }
        });
        client.ws.addEventListener('error', e => {
            _debug(client, 'WebSocket error', e);
            client.configuration.onClientError(ClientErrorType.WebSocket, null, e);
        });
        client.ws.addEventListener('message', e => {
            wsMessage(client, e.data);
        });
    }

    /**
     * @param {Object} client
     * @param {string} data
     */
    function wsMessage(client, data) {
        const message = JSON.parse(data);

        if (!message.type) {
            _debug(client, 'WebSocket incoming message empty type');
            return;
        }

        _debug(client, 'WebSocket incoming message', message.type, message);

        switch (message.type) {
            case MessageType.Greeting:
                wsMessageGreeting(client, message.greeting);
                break;
            case MessageType.Connected:
                wsMessageConnected(client, message.connected);
                break;
            case MessageType.Disconnected:
                wsMessageDisconnected(client, message.disconnected);
                break;
            case MessageType.Heartbeat:
                wsMessageHeartbeat(client);
                break;
            case MessageType.Error:
                wsMessageError(client, message.error);
                break;
            case MessageType.ClientNameChanged:
                wsMessageNameChanged(client, message.name_changed);
                break;
            case MessageType.PlayerChanged:
                wsMessagePlayerChanged(client, message.player_changed);
                break;
            case MessageType.SignallingOffer:
                wsMessageSignallingOffer(client, message.signalling);
                break;
            case MessageType.SignallingAnswer:
                wsMessageSignallingAnswer(client, message.signalling);
                break;
            case MessageType.SignallingIceCandidate:
                wsMessageSignallingIceCandidate(client, message.signalling);
                break;
            default:
                _debug(client, 'WebSocket incoming message unknown type', message.type);
        }
    }

    /**
     * @param {Object} client
     * @param {Object} message JSON message
     */
    function wsSend(client, message) {
        if (!client.ws) {
            return;
        }

        _debug(client, 'WebSocket send message', message.type, message);

        client.ws.send(JSON.stringify(message));
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.host_id
     * @param {string} message.client_id
     * @param {string} message.client_key
     * @param {string} message.name
     * @param {number} message.player
     * @param {[]} message.clients
     */
    function wsMessageGreeting(client, message) {
        client.hostId = message.host_id;
        client.clientId = message.client_id;
        client.clientKey = message.client_key;
        client.clients = {};

        for (let connectedClientIdx in message.clients) {
            const connectedClient = message.clients[connectedClientIdx];
            clientConnected(client, connectedClient.client_id, connectedClient);
        }

        client.configuration.onSelfNameChanged(message.name, client.name);
        client.name = message.name;

        client.configuration.onSelfPlayerChanged(message.player, client.player);
        client.player = message.player;

        client.configuration.onGreeting();
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.client_id
     * @param {string} message.name
     * @param {number} message.player
     */
    function wsMessageConnected(client, message) {
        clientConnected(client, message.client_id, message);
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.client_id
     */
    function wsMessageDisconnected(client, message) {
        clientDisconnected(client, message.client_id);
    }

    /**
     * @param {Object} client
     */
    function wsMessageHeartbeat(client) {
        wsSend(client, _messageHeartbeat());
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.message
     */
    function wsMessageError(client, message) {
        client.configuration.onClientError(ClientErrorType.Server, null, new Error(message.message));
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.client_id
     * @param {string} message.name
     */
    function wsMessageNameChanged(client, message) {
        const clientId = message.client_id;
        const name = message.name;

        if (!client.clients[clientId]) {
            return;
        }

        client.configuration.onClientNameChanged(clientId, name, client.clients[clientId].name);
        client.clients[clientId].name = name;

        if (client.clientId === clientId) {
            client.configuration.onSelfNameChanged(name, client.name);
            client.name = name;
        }
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.client_id
     * @param {number} message.player
     */
    function wsMessagePlayerChanged(client, message) {
        const clientId = message.client_id;
        const player = message.player;

        if (!client.clients[clientId]) {
            return;
        }

        client.configuration.onClientPlayerChanged(clientId, player, client.clients[clientId].player);
        client.clients[clientId].player = player;

        if (client.clientId === clientId) {
            client.configuration.onSelfPlayerChanged(player, client.player);
            client.player = player;
        }
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.from_id
     * @param {string} message.sdp
     */
    function wsMessageSignallingOffer(client, message) {
        const clientId = message.from_id;
        const sdp = JSON.parse(atob(message.sdp));

        if (client.rtcClients[clientId]) {
            rtcHandleRemoteDescription(client, clientId, client.rtcClients[clientId], sdp);
        }
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.from_id
     * @param {string} message.sdp
     */
    function wsMessageSignallingAnswer(client, message) {
        const clientId = message.from_id;
        const sdp = JSON.parse(atob(message.sdp));

        if (client.rtcClients[clientId]) {
            rtcHandleRemoteDescription(client, clientId, client.rtcClients[clientId], sdp);
        }
    }

    /**
     * @param {Object} client
     * @param {Object} message
     * @param {string} message.from_id
     * @param {string} message.sdp
     */
    function wsMessageSignallingIceCandidate(client, message) {
        const clientId = message.from_id;
        const sdp = JSON.parse(atob(message.sdp));

        if (client.rtcClients[clientId]) {
            rtcHandleIceCandidate(client, clientId, client.rtcClients[clientId], sdp);
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // RTC
    ///////////////////////////////////////////////////////////////////////////

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     */
    function rtcConnect(client, destinationClientId) {
        if (client.clientId === destinationClientId || !client.clients[destinationClientId]) {
            //not allowed to connect to self or if destination client not exists
            return;
        }

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
            rtcConnectionStateChanged(client, destinationClientId, connection);
        });
        connection.addEventListener('datachannel', e => {
            rtcDCConnected(client, destinationClientId, e.channel);
        });
        connection.addEventListener('icecandidate', e => {
            rtcSendIceCandidate(client, destinationClientId, connection, e);
        });
        connection.addEventListener('icecandidateerror', e => {
            rtcIceCandidateError(client, destinationClientId, connection, e);
        });
        connection.addEventListener('iceconnectionstatechange', () => {
            rtcIceConnectionStateChanged(client, destinationClientId, connection);
        });
        connection.addEventListener('icegatheringstatechange', () => {
            _debug(client, 'RTC ICE gathering state changed', destinationClientId, connection.iceGatheringState);
            client.configuration.onRTCIceGatheringStateChanged(destinationClientId, connection.iceConnectionState);
        });
        connection.addEventListener('negotiationneeded', () => {
           rtcSendDescription(client, destinationClientId, connection);
        });
        connection.addEventListener('signalingstatechange', () => {
            _debug(client, 'RTC signalling state changed', destinationClientId, connection.signalingState);
           client.configuration.onRTCSignallingStateChanged(destinationClientId, connection.signalingState);
        });
        connection.addEventListener('track', e => {
            rtcTrack(client, destinationClientId, e);
        });

        client.rtcClients[destinationClientId] = connection;

        if (client.host) {
            const mediaStream = collectMediaTracks(client);
            mediaStream.getTracks().forEach(track => connection.addTrack(track, mediaStream));

            rtcDCConnect(client, destinationClientId);
        }
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     */
    function rtcConnectionStateChanged(client, destinationClientId, connection) {
        _debug(client, 'RTC connection state changed', destinationClientId, connection.connectionState);

        client.configuration.onRTCConnectionStateChanged(destinationClientId, connection.connectionState);

        if (connection.connectionState === 'connected') {
            client.retries[destinationClientId] = 0;
        }
        if (connection.connectionState === 'failed') {
            client.configuration.onClientError(ClientErrorType.RtcConnection, destinationClientId, new Error('RTC connection failed'));
        }
        if (connection.connectionState === 'failed' || connection.connectionState === 'closed') {
            if (client.retryConnection && client.clients[destinationClientId]) {
                client.configuration.onRTCReconnecting(destinationClientId);
                _retry(client, destinationClientId, () => rtcConnect(client, destinationClientId));
            }
        }
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     * @param {Event} e
     */
    function rtcIceCandidateError(client, destinationClientId, connection, e) {
        _debug(client, 'RTC ICE candidate for error', destinationClientId, e);
        client.configuration.onClientError(ClientErrorType.RtcIceCandidate, destinationClientId, e);
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     */
    function rtcIceConnectionStateChanged(client, destinationClientId, connection) {
        _debug(client, 'RTC ICE connection state changed', destinationClientId, connection.iceConnectionState);

        client.configuration.onRTCIceStateChanged(destinationClientId, connection.iceConnectionState);

        if (connection.iceConnectionState === 'failed') {
            client.configuration.onClientError(ClientErrorType.RtcIceConnection, destinationClientId, new Error('RTC ICE connection failed'));
            if (client.retryConnection && client.clients[destinationClientId]) {
                connection.restartIce();
            }
        }
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     */
    function rtcSendDescription(client, destinationClientId, connection) {
        client.rtcMakingOffer[destinationClientId] = true;

        connection
            .setLocalDescription()
            .then(() => {
                wsSend(client, _messageSignalling(MessageType.SignallingOffer, destinationClientId, connection.localDescription));
                client.rtcMakingOffer[destinationClientId] = false;
            })
            .catch(e => {
                _debug(client, 'RTC send offer error', destinationClientId, e);
                client.rtcMakingOffer[destinationClientId] = false;
                client.configuration.onClientError(ClientErrorType.RtcLocalDescription, destinationClientId, e);
            });
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     * @param {RTCSessionDescriptionInit} sdp
     */
    function rtcHandleRemoteDescription(client, destinationClientId, connection, sdp) {
        const description = new RTCSessionDescription(sdp);

        const offerCollision = (description.type === 'offer' && (client.rtcMakingOffer[destinationClientId] || connection.signalingState !== 'stable'));

        client.rtcIgnoreOffer[destinationClientId] = (!rtcPolite(client) && offerCollision);

        if (client.rtcIgnoreOffer[destinationClientId]) {
            return;
        }

        connection
            .setRemoteDescription(description)
            .then(() => {
                if (description.type !== 'offer') {
                    return;
                }
                connection
                    .setLocalDescription()
                    .then(() => {
                        wsSend(client, _messageSignalling(MessageType.SignallingAnswer, destinationClientId, connection.localDescription));
                    })
                    .catch(e => {
                        _debug(client, 'RTC send answer error', destinationClientId, e);
                        client.configuration.onClientError(ClientErrorType.RtcLocalDescription, destinationClientId, e);
                    });
            })
            .catch(e => {
                _debug(client, 'RTC set remote description error', destinationClientId, e);
                client.configuration.onClientError(ClientErrorType.RtcRemoteDescription, destinationClientId, e);
            });
    }

    /**
     * @param {Object} client
     * @returns {boolean}
     */
    function rtcPolite(client) {
        return !client.host;
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     * @param {RTCPeerConnectionIceEvent} e
     */
    function rtcSendIceCandidate(client, destinationClientId, connection, e) {
        if (!e.candidate) {
            return;
        }
        wsSend(client, _messageSignalling(MessageType.SignallingIceCandidate, destinationClientId, e.candidate));
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCPeerConnection} connection
     * @param {RTCIceCandidateInit} sdp
     */
    function rtcHandleIceCandidate(client, destinationClientId, connection, sdp) {
        const candidate = new RTCIceCandidate(sdp);

        connection
            .addIceCandidate(candidate)
            .catch(e => {
                _debug(client, 'RTC handle ICE candidate from error', destinationClientId, e);
                client.configuration.onClientError(ClientErrorType.RtcIceCandidateAccept, destinationClientId, e);
            });
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCTrackEvent} e
     */
    function rtcTrack(client, destinationClientId, e) {
        if (client.host) {
            return;
        }
        if (!e.track || !e.streams || e.streams.length === 0) {
            console.error('No media streams received');
            return;
        }

        if (e.track.kind === 'video') {
            _debug(client, 'RTC video track received', destinationClientId);
            client.configuration.onRTCTrack(TrackType.Video, e.streams[0].getVideoTracks());
        } else if (e.track.kind === 'audio') {
            _debug(client, 'RTC audio track received', destinationClientId);
            client.configuration.onRTCTrack(TrackType.Audio, e.streams[0].getAudioTracks());
        }

        e.track.addEventListener('unmute', () => {
            client.configuration.gameVideoEl.srcObject = e.streams[0];
        }, {once: true});
    }

    ///////////////////////////////////////////////////////////////////////////
    // RTC Data Channel (for control input)
    ///////////////////////////////////////////////////////////////////////////

    /**
     * (for host)
     * @param client Object
     * @param destinationClientId string
     */
    function rtcDCConnect(client, destinationClientId) {
        if (client.clientId === destinationClientId || !client.clients[destinationClientId] || !client.rtcClients[destinationClientId]) {
            //not allowed to connect to self or of destination client not exists or destination connection not exists
            return;
        }

        const channelLabel = _buildDataChannelLabel();
        _debug(client, 'RTC DC to client', destinationClientId, channelLabel);

        const dataChannel = client.rtcClients[destinationClientId].createDataChannel(channelLabel);

        dataChannel.addEventListener('open', () => {
            _debug(client, 'RTC DC to client open', destinationClientId, channelLabel);
            client.retries[destinationClientId] = 0;
            client.configuration.onRTCControlChannelOpen(destinationClientId);
        });
        dataChannel.addEventListener('close', () => {
            _debug(client, 'RTC DC to client close', destinationClientId, channelLabel);
            client.configuration.onRTCControlChannelClose(destinationClientId);
            if (client.retryConnection && client.clients[destinationClientId] && client.rtcClients[destinationClientId]) {
                client.configuration.onRTCControlChannelReconnecting(destinationClientId);
                _retry(client, destinationClientId, () => rtcDCConnect(client, destinationClientId));
            }
        });
        dataChannel.addEventListener('error', e => {
            _debug(client, 'RTC DC to client error', destinationClientId, channelLabel, e);
            client.configuration.onClientError(ClientErrorType.RtcControlChannel, destinationClientId, e);
        });
        dataChannel.addEventListener('message', e => {
            rtcDCReceive(client, destinationClientId, channelLabel, e.data);
        });

        client.rtcControlChannels[destinationClientId] = dataChannel;
    }

    /**
     * (for client)
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {RTCDataChannel} dataChannel
     */
    function rtcDCConnected(client, destinationClientId, dataChannel) {
        if (!client.clients[destinationClientId] || destinationClientId !== client.hostId) {
            dataChannel.close();
            return;
        }

        client.rtcControlChannels[destinationClientId] = dataChannel;

        dataChannel.addEventListener('open', () => {
            _debug(client, 'RTC DC to host open', destinationClientId, dataChannel.label);
            client.configuration.onRTCControlChannelOpen(client.hostId);
        });

        //host doesn't send any data to clients (yet)

        dataChannel.addEventListener('close', () => {
            _debug(client, 'RTC DC to host closed', destinationClientId, dataChannel.label);
            client.configuration.onRTCControlChannelClose(destinationClientId);
        });
        dataChannel.addEventListener('error', e => {
            _debug(client, 'RTC DC to host error', destinationClientId, dataChannel.label, e);
            client.configuration.onClientError(ClientErrorType.RtcControlChannel, client.hostId, e);
        });
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {string} channelLabel
     * @param {string} data
     */
    function rtcDCReceive(client, destinationClientId, channelLabel, data) {
        if (!client.clients[destinationClientId]) {
            return;
        }

        _debug(client, 'RTC DC from client message', destinationClientId, channelLabel, data);

        const message = JSON.parse(data);
        if (!message) {
            return;
        }

        switch (message.type) {
            case MessageType.DCInput:
                const destinationClient = client.clients[destinationClientId];
                if (destinationClient.player !== -1) {
                    client.configuration.onRTCControlChannelInput(destinationClientId, destinationClient.player, message.code, message.value);
                }
                break;

            case MessageType.DCHeartbeat:
                break;

            case MessageType.DCPlayer:
                wsSend(this, _messagePlayerChange(destinationClientId, message.player));
                break;

            default:
                console.log('Unknown control message type', message.type, message);
        }
    }

    /**
     * @param {Object} client
     * @param {string} destinationClientId
     * @param {Object} message
     */
    function rtcDCSend(client, destinationClientId, message) {
        if (!client.rtcControlChannels[destinationClientId]) {
            return;
        }

        const channel = client.rtcControlChannels[destinationClientId];
        if (channel.readyState !== 'open') {
            return;
        }

        channel.send(JSON.stringify(message));
    }

    ///////////////////////////////////////////////////////////////////////////
    // Utils
    ///////////////////////////////////////////////////////////////////////////

    /**
     * @param {Object} client
     * @param {*} args
     * @private
     */
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

    /**
     * @param {Object} client
     * @param {string|null} destinationClientId
     * @param {function} action
     * @private
     */
    function _retry(client, destinationClientId, action) {
        if (client.retries[destinationClientId] === undefined) {
            client.retries[destinationClientId] = 0;
        }
        if (client.retries[destinationClientId] >= RetryLimit) {
            client.configuration.onClientError(ClientErrorType.RetryLimit, destinationClientId, new Error('Retry limit exceeded'));
            if (!client.host || !destinationClientId) {
                client.disconnect();
            } else if (client.host && destinationClientId) {
                clientDisconnected(client, destinationClientId);
            }
            return;
        }

        client.retries[destinationClientId]++;

        setTimeout(action, RetryTimeout + Math.round(Math.random() * 500 - 250));
    }

    /**
     * @param {Object} config
     * @private
     */
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

    /**
     * @param {string} gameId
     * @param {string} sessionId
     * @returns {string}
     * @private
     */
    function _buildWebSocketUrl(gameId, sessionId) {
        const protocol = location.protocol === 'https:' ? 'wss' : 'ws';
        return `${protocol}://${location.host}/netplay/${gameId}/${sessionId}/ws`;
    }

    /**
     * @param {string} url
     * @param {string} user
     * @param {string} password
     * @returns {{urls}}
     * @private
     */
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

    /**
     * @returns {string}
     * @private
     */
    function _buildDataChannelLabel() {
        const id = Math.round(Math.random() * 1000000);
        return `control-${id}`;
    }

    /**
     * @param {string} name
     * @returns {{type: string, name_change: {name}}}
     * @private
     */
    function _messageNameChange(name) {
        return {
            type: MessageType.ClientNameChange,
            name_change: {
                name: name,
            },
        };
    }

    /**
     * @param {string} clientId
     * @param {number} player
     * @returns {{player_change: {client_id, player}, type: string}}
     * @private
     */
    function _messagePlayerChange(clientId, player) {
        return {
            type: MessageType.PlayerChange,
            player_change: {
                client_id: clientId,
                player: player,
            },
        };
    }

    /**
     * @returns {{type: string}}
     * @private
     */
    function _messageHeartbeat() {
        return {
            type: MessageType.Heartbeat,
        };
    }

    /**
     * @param {string} type
     * @param {string} clientId
     * @param {Object} sdp
     * @returns {{type, signalling: {client_id, sdp: string}}}
     * @private
     */
    function _messageSignalling(type, clientId, sdp) {
        return {
            type: type,
            signalling: {
                client_id: clientId,
                sdp: btoa(JSON.stringify(sdp)),
            },
        }
    }

    /**
     * @returns {{type: string}}
     * @private
     */
    function _messageDCHeartbeat() {
        return {
            type: MessageType.DCHeartbeat,
        };
    }

    /**
     * @param {number} inputCode
     * @param {number} inputValue
     * @returns {{code, type: string, value}}
     * @private
     */
    function _messageDCInput(inputCode, inputValue) {
        return {
            type: MessageType.DCInput,
            code: inputCode,
            value: inputValue,
        }
    }

    /**
     * @param {number} player
     * @returns {{type: string, player}}
     * @private
     */
    function _messageDCPlayer(player) {
        return {
            type: MessageType.DCPlayer,
            player: player,
        };
    }

})();
