(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        window.addEventListener('resize', setupGameDisplaySize);
        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);
        document.getElementById('netplay-virtual-gamepad-toggle').addEventListener('click', virtualGamepadToggle);
        document.getElementById('netplay-play').addEventListener('click', play);

        if (fullscreenEnabled()) {
            const fullscreenButton = document.getElementById('netplay-fullscreen');
            fullscreenButton.addEventListener('click', fullscreen);
            fullscreenButton.classList.remove('d-none');
        }

        const gameEl = document.getElementById('game');
        gameEl.addEventListener('play', () => playScreen(false));
        gameEl.addEventListener('pause', () => playScreen(true));
        gameEl.disablePictureInPicture = true;

        const gameOverlayEl = document.getElementById('game-overlay');
        gameOverlayEl.addEventListener('keydown', controlsButtonDown);
        gameOverlayEl.addEventListener('keyup', controlsButtonUp);
        setInterval(controlsPollGamepad, 1000 / 60);

        gameOverlayEl.focus();

        virtualGamepadInit();
        virtualGamepadLoad();
        setupGameDisplaySize();

        netplay = NetplayClient({
            debug: window.NetplayDebug,
            gameVideoEl: gameEl,
            gameId: window.NetplayGameId,
            sessionId: window.NetplaySessionId,
            host: false,
            turnServerUrl: window.NetplayTurnServerUrl,
            turnServerUser: window.NetplayTurnServerUser,
            turnServerPassword: window.NetplayTurnServerPassword,

            onWSConnected: wsConnected,
            onRTCConnectionStateChanged: rtcConnectionStateChanged,
            onRTCControlChannelOpen: controlDataChannelOpen,

            onGreeting: wsGreeting,
            onSelfNameChanged: selfNameChanged,
            onSelfPlayerChanged: selfPlayerChanged,

            onClientCleanState: clientReset,
            onClientConnected: clientConnected,
            onClientDisconnected: clientDisconnected,
            onClientNameChanged: clientNameChanged,
            onClientPlayerChanged: clientPlayerChanged,

            onClientError: errorHandler,
            onWSReconnecting: wsReconnecting,
            onRTCReconnecting: rtcReconnecting,
            onRTCControlChannelReconnecting: rtcDCReconnecting,
        });

        connectionScreen('Connecting to server');
        netplay.connect();
    });

    ///////////////////////////////////////////////////////////////////////////
    // Controls
    ///////////////////////////////////////////////////////////////////////////

    const ButtonLabels = {
        0: 'BUTTON_1',
        1: 'BUTTON_2',
        2: 'BUTTON_3',
        3: 'BUTTON_4',
        4: 'LEFT_TOP_SHOULDER',
        5: 'RIGHT_TOP_SHOULDER',
        6: 'LEFT_BOTTOM_SHOULDER',
        7: 'RIGHT_BOTTOM_SHOULDER',
        8: 'SELECT',
        9: 'START',
        10: 'LEFT_STICK',
        11: 'RIGHT_STICK',
        12: 'DPAD_UP',
        13: 'DPAD_DOWN',
        14: 'DPAD_LEFT',
        15: 'DPAD_RIGHT',
    };

    let gamepadId = null;
    let gamepadPrevState = {};

    function controlsButtonDown(e) {
        if (netplay.getPlayer() === -1 || e.repeat) {
            return;
        }
        controlMapButtons(e.keyCode)
            .forEach(button => netplay.sendControlInput(button, 1.0));
        e.preventDefault();
    }

    function controlsButtonUp(e) {
        if (netplay.getPlayer() === -1) {
            return;
        }
        controlMapButtons(e.keyCode)
            .forEach(button => netplay.sendControlInput(button, 0.0));
        e.preventDefault();
    }

    function controlMapButtons(value) {
        const buttons = [];
        for (let code in window.ControlScheme) {
            if (window.ControlScheme[code].keycode === value) {
                buttons.push(code);
            }
        }
        return buttons;
    }

    function controlsPollGamepad() {
        if (netplay.getPlayer() === -1 || !navigator.getGamepads) {
            return;
        }

        let gamepad = null;

        //select current gamepad

        navigator.getGamepads().forEach(g => {
            if (!g) {
                return;
            }
            if (g.id === gamepadId) {
                gamepad = g;
            }
        });
        if (gamepad === null) {
            const gamepads = navigator.getGamepads();
            if (gamepads.length === 0) {
                return;
            }
            gamepad = gamepads[0];
            if (gamepad === null) {
                return;
            }
            gamepadId = gamepad.id;
            gamepadPrevState = {};
        }

        //collect pressed buttons

        const pressedButtons = [];
        gamepad.buttons.forEach((button, buttonIdx) => {
            if (button.pressed) {
                if (ButtonLabels[buttonIdx] !== undefined) {
                    pressedButtons.push(ButtonLabels[buttonIdx]);
                } else {
                    pressedButtons.push(`GAMEPAD_${buttonIdx}`);
                }
            }
        });
        gamepad.axes.forEach((axisValue, axisIdx) => {
            const name = ['LEFT_STICK_X', 'LEFT_STICK_Y', 'RIGHT_STICK_X', 'RIGHT_STICK_Y'][axisIdx];
            if (!name) {
                return;
            }
            const value = Math.round(axisValue);
            if (value === 0) {
                return;
            }
            pressedButtons.push(`${name}:${value < 0 ? '-1' : '+1'}`);
        });

        //map state

        const gamepadCurrentState = {};
        pressedButtons.forEach(button => {
            for (let code in window.ControlScheme) {
                if (window.ControlScheme[code].value2 === button) {
                    gamepadCurrentState[code] = true;
                }
            }
        });

        //compare states

        for (let code in gamepadCurrentState) {
            if (gamepadCurrentState[code] && !gamepadPrevState[code]) {
                netplay.sendControlInput(code, 1.0);
            }
        }
        for (let code in gamepadPrevState) {
            if (gamepadPrevState[code] && !gamepadCurrentState[code]) {
                netplay.sendControlInput(code, 0.0);
            }
        }

        gamepadPrevState = gamepadCurrentState;
    }

    function controlDataChannelOpen() {
        setInterval(() => netplay.sendControlHeartbeat(), 5000);
    }

    ///////////////////////////////////////////////////////////////////////////
    // Self name and player
    ///////////////////////////////////////////////////////////////////////////

    function selfNameChanged(name) {
        document.getElementById('netplay-name').value = name;
        document.getElementById('netplay-player').innerText = `${name}: ${NetplayPlayerDisplay(netplay.getClientId(), netplay.getHostId(), netplay.getPlayer())}`;
    }

    function selfPlayerChanged(player) {
        const playerDisplay = NetplayPlayerDisplay(netplay.getClientId(), netplay.getHostId(), player);
        window.ShowToastMessage('primary', `You now play as ${playerDisplay}`);
        document.getElementById('netplay-player').innerText = `${netplay.getName()}: ${playerDisplay}`;
        virtualGamepadShow(player);
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

    function wsGreeting() {
        //send previously saved client name
        const netplayName = netplay.getName();
        const savedName = window.NetplayLoadClientName(netplayName);
        if (savedName && netplayName !== savedName) {
            netplay.setName(savedName);
        }

        //send previously saved client player
        const netplayPlayer = netplay.getPlayer();
        const savedPlayer = window.NetplayLoadClientPlayer(window.NetplayGameId, window.NetplaySessionId, netplayPlayer);
        if (savedPlayer !== netplayPlayer) {
            netplay.setPlayer(savedPlayer);
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Game session clients list
    ///////////////////////////////////////////////////////////////////////////

    function clientReset() {
        document.getElementById('netplay-clients').innerHTML = '';
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

        const el = document.createElement('div');
        el.id = `netplay-client-${id}`;
        el.classList.add('list-group-item');

        const containerEl = document.createElement('div');
        containerEl.classList.add('row');
        el.append(containerEl);

        const nameEl = document.createElement('div');
        nameEl.id = `netplay-client-${id}-name`;
        nameEl.classList.add('lead', 'col-6', 'col-md-9');
        nameEl.innerText = name;
        containerEl.append(nameEl);

        const playerEl = document.createElement('div');
        playerEl.id = `netplay-client-${id}-player`;
        playerEl.classList.add('text-end', 'col-6', 'col-md-3');
        playerEl.innerText = NetplayPlayerDisplay(id, netplay.getHostId(), player);
        containerEl.append(playerEl);

        document.getElementById('netplay-clients').append(el);
    }

    function clientDisconnected(id) {
        const client = netplay.getClient(id);
        if (client && id !== netplay.getClientId()) {
            window.ShowToastMessage('warning', `${client.name} (${NetplayPlayerDisplay(id, netplay.getHostId(), client.player)}) disconnected`);
        }
        if (id === netplay.getHostId()) {
            connectionScreen('Awaiting game host');
            playScreen(false);
        }

        const el = document.getElementById(`netplay-client-${id}`);
        if (el) {
            el.remove();
        }
    }

    function clientNameChanged(id, name) {
        const client = netplay.getClient(id);
        if (client && id !== netplay.getClientId()) {
            window.ShowToastMessage('info', `${client.name} (${NetplayPlayerDisplay(id, netplay.getHostId(), client.player)}) is now ${name}`);
        }

        const el = document.getElementById(`netplay-client-${id}-name`);
        if (el) {
            el.innerText = name;
        }
    }

    function clientPlayerChanged(id, newPlayer, oldPlayer) {
        const client = netplay.getClient(id);
        if (client && id !== netplay.getClientId()) {
            const oldPlayerDisplay = window.NetplayPlayerDisplay(id, netplay.getHostId(), oldPlayer);
            const newPlayerDisplay = window.NetplayPlayerDisplay(id, netplay.getHostId(), newPlayer);
            window.ShowToastMessage('info', `${client.name} (${oldPlayerDisplay}) is now ${newPlayerDisplay}`);
        }

        const el = document.getElementById(`netplay-client-${id}-player`);
        if (el) {
            el.innerText = window.NetplayPlayerDisplay(id, netplay.getHostId(), newPlayer);
        }

        if (id === netplay.getClientId()) {
            window.NetplaySaveClientPlayer(window.NetplayGameId, window.NetplaySessionId, newPlayer);
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Virtual gamepad
    ///////////////////////////////////////////////////////////////////////////

    let virtualGamepadVisible = false;

    /*
     * Gamepad button definition fields:
     * - region - string (left, center, right)
     * - square - boolean
     * - circle - boolean
     * - large - boolean
     * - small - boolean
     * - text - string (alternative to icon)
     * - icon - string (one of bi-* classes)
     * - top - int - top relative position, px
     * - bottom - int - bottom relative position, px
     * - left - int - left relative position, px
     * - right - int - right relative position, px
     * - input - int - control input value
     */

    const virtualGamepads = {
        'nes': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start',  text: 'start',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
        ],
        'gba': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start',  text: 'start',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -80, left:  10, input: 10, id:'l', text: 'L', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -80, right: 10, input: 11, id:'r', text: 'R', cls: 'btn-outline-secondary', circle: true},
        ],
        'snes': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105,  left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105,  right: -15, input: 3, id:'start',  text: 'start',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 10,  right: 55,  input: 9, id:'x', text: 'X', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 100, right: 55,  input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55,  right: 100, input: 1, id:'y', text: 'Y', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55,  right: 10,  input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -80, left:  10, input: 10, id:'l', text: 'L', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -80, right: 10, input: 11, id:'r', text: 'R', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaMD': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'mode',  text: 'mode',  cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start', text: 'start', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 75,   left: -80, input: 1,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 75,   left: -15, input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 75,   left: 50,  input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: -80, input: 10, id:'x', text: 'X', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: -15, input: 9,  id:'y', text: 'Y', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: 50,  input: 11, id:'z', text: 'Z', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaMS': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'1', text: '1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'2', text: '2', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaGG': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, right: 20, input: 3, id:'start', text: 'start', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'1', text: '1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'2', text: '2', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaSaturn': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, right: 20, input: 3, id:'start', text: 'start', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 75,   left: -80, input: 1,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 75,   left: -15, input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 75,   left: 50,  input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: -80, input: 9,  id:'x', text: 'X', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: -15, input: 10, id:'y', text: 'Y', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: 50,  input: 11, id:'z', text: 'Z', cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -90, left:  10, input: 12, id:'l', text: 'L', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -90, right: 10, input: 13, id:'r', text: 'R', cls: 'btn-outline-secondary', circle: true},
        ],
        'atari2600': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 15,  left:  20, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 60,  right: 20, input: 3, id:'pause',  text: 'pause',  cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: 20, input: 9, id:'reset',  text: 'reset',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: 30, input: 0, id:'fire', icon: 'bi-fire', cls: 'btn-outline-danger', circle: true},
        ],
        'atari7800': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 15,  left:  20, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 60,  right: 20, input: 3, id:'pause',  text: 'pause',  cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: 20, input: 9, id:'reset',  text: 'reset',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'1', text: '1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'2', text: '2', cls: 'btn-outline-secondary', circle: true},
        ],
        'lynx': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 15,  left:  20, input: 10, id:'option1', text: 'opt 1', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 60,  right: 20, input: 3,  id:'start',   text: 'start', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: 20, input: 11, id:'option2', text: 'opt 2', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
        ],
        'jaguar': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'option', text: 'option', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'pause',  text: 'pause',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -80, input: 8,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: -15, input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 50,  input: 1,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
        ],
        '3do': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'X', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start',  text: 'P',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -80, input: 1,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: -15, input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 50,  input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
        ],
        'pce': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'run',    text: 'run',    cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'ii', text: 'II', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'i',  text: 'I',  cls: 'btn-outline-secondary', circle: true},
        ],
        'pcfx': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'run',    text: 'run',    cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 75,   left: 50,  input: 8,  id:'i',   text: 'I',   cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 75,   left: -15, input: 0,  id:'ii',  text: 'II',  cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 75,   left: -80, input: 9,  id:'iii', text: 'III', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: 50,  input: 1,  id:'iv',  text: 'IV',  cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: -15, input: 10, id:'v',   text: 'V',   cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -10,  left: -80, input: 11, id:'vi',  text: 'VI',  cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -90, left:  10, input: 12, id:'m1', text: 'M1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -90, right: 10, input: 13, id:'m2', text: 'M2', cls: 'btn-outline-secondary', circle: true},
        ],
        'ngp': [
            {region: 'left', top: 0,   left: 55,  input: 4, id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 110, left: 55,  input: 5, id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 0,   input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 110, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, right: -15, input: 3, id:'option',  text: 'option',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -20, input: 0, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 0,  input: 8, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
        ],
        'gb':      'nes',
        'nds':     'snes',
        'sega32x': 'segaMD',
        'segaCD':  'segaMD',
    };

    function virtualGamepadInit() {
        if (!virtualGamepads[window.GamePlatform]) {
            return;
        }

        let gamepad = virtualGamepads[window.GamePlatform];
        if (typeof gamepad === 'string') {
            gamepad = virtualGamepads[gamepad];
        }

        gamepad.forEach(buttonDefinition => {
            const button = document.createElement('button');
            button.type = 'button';
            button.classList.add('btn');

            if (buttonDefinition.cls) {
                button.classList.add(buttonDefinition.cls);
            }
            if (buttonDefinition.square) {
                if (buttonDefinition.large) {
                    button.classList.add('btn-square-lg');
                } else {
                    button.classList.add('btn-square');
                }
            } else if (buttonDefinition.circle) {
                if (buttonDefinition.large) {
                    button.classList.add('btn-circle-lg');
                } else {
                    button.classList.add('btn-circle');
                }
            } else{
                button.classList.add('btn-default');
                if (buttonDefinition.large) {
                    button.classList.add('btn-lg');
                } else if (buttonDefinition.small) {
                    button.classList.add('btn-sm');
                }
            }

            if (buttonDefinition.top !== undefined) {
                button.style.top = buttonDefinition.top + 'px';
            }
            if (buttonDefinition.bottom !== undefined) {
                button.style.bottom = buttonDefinition.bottom + 'px';
            }
            if (buttonDefinition.left !== undefined) {
                button.style.left = buttonDefinition.left + 'px';
            }
            if (buttonDefinition.right !== undefined) {
                button.style.right = buttonDefinition.right + 'px';
            }

            if (buttonDefinition.icon) {
                const icon = document.createElement('i');
                icon.classList.add('bi');
                icon.classList.add(buttonDefinition.icon);
                button.append(icon);
            } else if (buttonDefinition.text) {
                button.innerText = buttonDefinition.text;
                button.style.fontWeight = 'bold';
            }

            button.addEventListener('mousedown', e => {
                e.preventDefault();
                button.classList.add('active');
                virtualGamepadPress(buttonDefinition.input);
            });
            button.addEventListener('mouseup', e => {
                e.preventDefault();
                button.classList.remove('active');
                virtualGamepadRelease(buttonDefinition.input);
            });
            button.addEventListener('touchstart', e => {
                e.preventDefault();
                button.classList.add('active');
                virtualGamepadPress(buttonDefinition.input);
            });
            button.addEventListener('touchend', e => {
                e.preventDefault();
                button.classList.remove('active');
                virtualGamepadRelease(buttonDefinition.input);
            });
            button.addEventListener('touchcancel', e => {
                e.preventDefault();
                button.classList.remove('active');
                virtualGamepadRelease(buttonDefinition.input);
            });

            switch (buttonDefinition.region) {
                case 'left':
                    document.querySelector('#virtual-gamepad .virtual-gamepad-area-left').append(button);
                    break;
                case 'center':
                    document.querySelector('#virtual-gamepad .virtual-gamepad-area-center').append(button);
                    break;
                case 'right':
                    document.querySelector('#virtual-gamepad .virtual-gamepad-area-right').append(button);
                    break;
            }
        });
    }

    function virtualGamepadPress(input) {
        if (netplay.getPlayer() !== -1) {
            netplay.sendControlInput(input, 1.0);
        }
    }

    function virtualGamepadRelease(input) {
        if (netplay.getPlayer() !== -1) {
            netplay.sendControlInput(input, 0.0);
        }
    }

    function virtualGamepadLoad() {
        if (window.localStorage && window.localStorage.playtimeVirtualGamepad) {
            virtualGamepadVisible = (window.localStorage.playtimeVirtualGamepad === 'true');
        } else {
            virtualGamepadVisible = window.isMobile();
        }
    }

    function virtualGamepadToggle() {
        virtualGamepadVisible = !virtualGamepadVisible;
        if (window.localStorage) {
            window.localStorage.playtimeVirtualGamepad = (virtualGamepadVisible ? 'true' : 'false');
        }
        virtualGamepadShow(netplay.getPlayer());
    }

    function virtualGamepadShow(player) {
        const virtualGamepadControl = document.getElementById('netplay-virtual-gamepad-control');
        if (player !== -1) {
            virtualGamepadControl.classList.remove('d-none');
        } else {
            virtualGamepadControl.classList.add('d-none');
        }

        const virtualGamepadContainer = document.getElementById('virtual-gamepad');
        if (player !== -1 && virtualGamepadVisible) {
            virtualGamepadContainer.classList.remove('d-none');
        } else {
            virtualGamepadContainer.classList.add('d-none');
        }

        const btn = document.getElementById('netplay-virtual-gamepad-toggle');
        if (virtualGamepadVisible) {
            btn.classList.remove('btn-outline-secondary');
            btn.classList.add('btn-outline-success');
        } else {
            btn.classList.remove('btn-outline-success');
            btn.classList.add('btn-outline-secondary');
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Connection screen
    ///////////////////////////////////////////////////////////////////////////

    function wsConnected() {
        connectionScreen('Connected to server. Awaiting game host');
        errorScreen(false);
    }

    function wsReconnecting() {
        window.ShowToastMessage('warning', 'Reconnecting to server...', 2000);
    }

    function rtcConnectionStateChanged(clientId, state) {
        switch (state) {
            case 'connecting':
                connectionScreen('Connecting to game host');
                break;
            case 'connected':
                connectionScreen(false);
                errorScreen(false);
                playScreen(true);
                break;
        }
    }

    function rtcReconnecting() {
        window.ShowToastMessage('warning', 'Reconnecting to game host...', 2000);
    }

    function rtcDCReconnecting() {
        window.ShowToastMessage('warning', 'Reconnecting to game host...', 2000);
    }

    function connectionScreen(display) {
        const el = document.getElementById('connection-screen');
        if (display) {
            el.classList.remove('d-none');
            document.getElementById('connection-screen-status').innerText = display;
        } else {
            el.classList.add('d-none');
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Play screen
    ///////////////////////////////////////////////////////////////////////////

    function playScreen(display) {
        const el = document.getElementById('play-screen');
        if (display) {
            el.classList.remove('d-none');
        } else {
            el.classList.add('d-none');
        }
    }

    function play() {
        playScreen(false);
        document.getElementById('game').play();
    }

    ///////////////////////////////////////////////////////////////////////////
    // Error screen
    ///////////////////////////////////////////////////////////////////////////

    function errorHandler(type) {
        if (type === 'retry-limit') {
            playScreen(false);
            connectionScreen(false);
            errorScreen(true);
        }
        switch (type) {
            case 'web-socket':
                window.ShowToastMessage('danger', 'Server connection error');
                break;
            case 'rtc-offer-send':
            case 'rtc-answer-send':
            case 'rtc-ice-connection':
            case 'rtc-control-channel':
                window.ShowToastMessage('danger', 'Game host connection error');
                break;
            case 'rtc-answer-receive':
            case 'rtc-ice-candidate-accept':
                window.ShowToastMessage('warning', 'Game host connection warning');
                break;
            case 'rtc-connection':
                window.ShowToastMessage('danger', 'Game host connection lost');
                break;
            case'server':
                window.ShowToastMessage('danger', `Server error: ${e.message}`);
                break;
        }
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
    // Fullscreen
    ///////////////////////////////////////////////////////////////////////////

    function fullscreenEnabled() {
        return document.fullscreenEnabled || document.webkitFullscreenEnabled;
    }

    function fullscreen() {
        if (!fullscreenEnabled()) {
            return;
        }
        if (document.fullscreenElement || document.webkitFullscreenElement) {
            return;
        }

        const el = document.getElementById('game-container');

        if (el.requestFullscreen) {
            el.requestFullscreen();
        } else if (el.webkitRequestFullscreen) {
            el.webkitRequestFullscreen();
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Utils
    ///////////////////////////////////////////////////////////////////////////

    function setupGameDisplaySize() {
        let headerOffset = 50;
        if (document.fullscreenElement || document.webkitFullscreenElement) {
            headerOffset = 0;
        }

        const game = document.getElementById('game');
        game.width = window.innerWidth;
        game.height = window.innerHeight - headerOffset;
    }

})();
