(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        window.addEventListener('resize', setupGameDisplaySize);
        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);
        document.getElementById('netplay-control-scheme-save').addEventListener('click', saveControlScheme);
        document.getElementById('netplay-control-scheme-reset').addEventListener('click', resetControlScheme);
        document.getElementById('netplay-virtual-gamepad-toggle').addEventListener('click', virtualGamepadToggle);
        document.getElementById('netplay-play').addEventListener('click', play);
        document.getElementById('netplay-fullscreen').addEventListener('click', fullscreen);
        document.getElementById('game').addEventListener('play', () => playScreen(false));
        document.getElementById('game').addEventListener('pause', () => playScreen(true));

        window.addEventListener('keydown', controlsButtonDown);
        window.addEventListener('keyup', controlsButtonUp);
        setInterval(controlsPollGamepad, 1000 / 60);

        loadControlScheme();
        virtualGamepadInit();
        virtualGamepadLoad();
        setupGameDisplaySize();

        netplay = NetplayClient({
            debug: window.NetplayDebug,
            gameVideoEl: document.getElementById('game'),
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

    let gamepadId = null;
    let gamepadPrevState = {};

    function controlsButtonDown(e) {
        if (netplay.getPlayer() === -1 || e.repeat) {
            return;
        }
        controlMapButtons(e.key.toLowerCase())
            .forEach(button => netplay.sendControlInput(button, 1.0));
        e.preventDefault();
    }

    function controlsButtonUp(e) {
        if (netplay.getPlayer() === -1) {
            return;
        }
        controlMapButtons(e.key.toLowerCase())
            .forEach(button => netplay.sendControlInput(button, 0.0));
        e.preventDefault();
    }

    function controlMapButtons(value) {
        const buttons = [];
        for (let code in window.ControlScheme) {
            if (window.ControlScheme[code].value === value) {
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
                pressedButtons.push(buttonIdx.toString());
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
            pressedButtons.push(`${name}:${value}`);
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
        if (latestPlayer !== null) {
            netplay.sendControlPlayer(latestPlayer);
        }
        setInterval(() => netplay.sendControlHeartbeat(), 5000);
    }

    ///////////////////////////////////////////////////////////////////////////
    // Control scheme
    ///////////////////////////////////////////////////////////////////////////

    const controlSchemeMapping = {
        0: 'b',
        1: 'y',
        2: 'select',
        3: 'start',
        4: 'up',
        5: 'down',
        6: 'left',
        7: 'right',
        8: 'a',
        9: 'x',
        10: 'l',
        11: 'r',
        12: 'l2',
        13: 'r2',
        14: 'l3',
        15: 'r3',
        16: 'l-stick-right',
        17: 'l-stick-left',
        18: 'l-stick-down',
        19: 'l-stick-up',
        20: 'r-stick-right',
        21: 'r-stick-left',
        22: 'r-stick-down',
        23: 'r-stick-up',
    };

    let defaultControlScheme = {};

    function loadControlScheme() {
        defaultControlScheme = Object.assign({}, window.ControlScheme);

        if (!window.localStorage || !window.localStorage.playtimeNetplayControls) {
            return;
        }

        let controls = {};

        try {
            controls = JSON.parse(window.localStorage.playtimeNetplayControls);
            controls = controls[window.GamePlatform];
        } catch (e) {
            console.error('Unable to load controls', e);
            return;
        }
        if (!controls) {
            return;
        }
        window.ControlScheme = controls;
        renderFormControls();
    }

    function saveControlScheme() {
        for (let buttonId in window.ControlScheme) {
            const buttonName = controlSchemeMapping[buttonId];

            const keyboardInput = document.querySelector(`input.keyboard[data-btn="${buttonName}"]`);
            if (keyboardInput) {
                window.ControlScheme[buttonId].value = keyboardInput.value;
            }

            const gamepadInput = document.querySelector(`input.gamepad[data-btn="${buttonName}"]`);
            if (gamepadInput) {
                window.ControlScheme[buttonId].value2 = gamepadInput.value;
            }
        }

        if (window.localStorage) {
            let controls;
            try {
                controls = JSON.parse(window.localStorage.playtimeNetplayControls);
            } catch (e) {
                controls = {};
            }
            controls[window.GamePlatform] = window.ControlScheme;
            window.localStorage.playtimeNetplayControls = JSON.stringify(controls);
        }
    }

    function resetControlScheme() {
        window.ControlScheme = Object.assign({}, defaultControlScheme);
        renderFormControls();
    }

    function renderFormControls() {
        for (let buttonId in window.ControlScheme) {
            const buttonName = controlSchemeMapping[buttonId];

            const keyboardInput = document.querySelector(`input.keyboard[data-btn="${buttonName}"]`);
            if (keyboardInput) {
                keyboardInput.value = window.ControlScheme[buttonId].value;
            }

            const gamepadInput = document.querySelector(`input.gamepad[data-btn="${buttonName}"]`);
            if (gamepadInput) {
                gamepadInput.value = window.ControlScheme[buttonId].value2;
            }
        }
    }

    ///////////////////////////////////////////////////////////////////////////
    // Self name and player
    ///////////////////////////////////////////////////////////////////////////

    let latestPlayer = null;

    function selfNameChanged(name) {
        document.getElementById('netplay-name').value = name;
        document.getElementById('netplay-player').innerText = `${name}: ${NetplayPlayerDisplay(netplay.getClientId(), netplay.getHostId(), netplay.getPlayer())}`;
    }

    function selfPlayerChanged(player) {
        const playerDisplay = NetplayPlayerDisplay(netplay.getClientId(), netplay.getHostId(), player);
        window.ShowToastMessage('primary', `You now play as ${playerDisplay}`);
        document.getElementById('netplay-player').innerText = `${netplay.getName()}: ${playerDisplay}`;
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
        if (netplayName !== savedName) {
            netplay.setName(savedName);
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
            latestPlayer = newPlayer;
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
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start',  text: 'start',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -10, input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 10, input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
        ],
        'gba': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start',  text: 'start',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -10, input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 10, input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -50, left:  10, input: 10, id:'l', text: 'L', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -50, right: 10, input: 11, id:'r', text: 'R', cls: 'btn-outline-secondary', circle: true},
        ],
        'snes': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105,  left:  -15, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105,  right: -15, input: 3, id:'start',  text: 'start',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 10,  right: 55,  input: 9, id:'x', text: 'X', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 100, right: 55,  input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55,  right: 100, input: 1, id:'y', text: 'Y', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55,  right: 10,  input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -50, left:  10, input: 10, id:'l', text: 'L', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -50, right: 10, input: 11, id:'r', text: 'R', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaMD': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'mode',  text: 'mode',  cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start', text: 'start', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -50, input: 9,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 0,   input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 50,  input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 0,  left: -50, input: 10, id:'x', text: 'X', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 0,  left: 0,   input: 9,  id:'y', text: 'Y', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 0,  left: 50,  input: 11, id:'z', text: 'Z', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaMS': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'right', top: 55, left: -10, input: 0, id:'1', text: '1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 10, input: 8, id:'2', text: '2', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaGG': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, right: 20, input: 3, id:'start', text: 'start', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -10, input: 0, id:'1', text: '1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 10, input: 8, id:'2', text: '2', cls: 'btn-outline-secondary', circle: true},
        ],
        'segaSaturn': [
            {region: 'left', top: 10,  left: 55, input: 4,  id:'up',    icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5,  id:'down',  icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, right: 20, input: 3, id:'start', text: 'start', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -50, input: 1,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 0,   input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 50,  input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 0,  left: -50, input: 9,  id:'x', text: 'X', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 0,  left: 0,   input: 10, id:'y', text: 'Y', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 0,  left: 50,  input: 11, id:'z', text: 'Z', cls: 'btn-outline-secondary', circle: true},

            {region: 'left',  top: -60, left:  10, input: 12, id:'l', text: 'L', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: -60, right: 10, input: 13, id:'r', text: 'R', cls: 'btn-outline-secondary', circle: true},
        ],
        'atari2600': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 15, left:  20, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 60, right: 20, input: 3, id:'pause',  text: 'pause',  cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: 20, input: 9, id:'reset',  text: 'reset',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: 30, input: 0, id:'fire', icon: 'bi-fire', cls: 'btn-outline-danger', circle: true},
        ],
        'atari7800': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 15,  left:  20, input: 2, id:'select', text: 'select', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 60,  right: 20, input: 3, id:'pause',  text: 'pause',  cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: 20, input: 9, id:'reset',  text: 'reset',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -10, input: 0, id:'1', text: '1', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 10, input: 8, id:'2', text: '2', cls: 'btn-outline-secondary', circle: true},
        ],
        'lynx': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 15,  left:  20, input: 10, id:'option1', text: 'opt 1', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 60,  right: 20, input: 3,  id:'start',   text: 'start', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: 20, input: 11, id:'option2', text: 'opt 2', cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -10, input: 0, id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, right: 10, input: 8, id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
        ],
        'jaguar': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'option', text: 'option', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'pause',  text: 'pause',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: 50,  input: 8,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 0,   input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: -50, input: 1,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
        ],
        '3do': [
            {region: 'left', top: 10,  left: 55, input: 4, id:'up',     icon: 'bi-caret-up-fill',    cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 100, left: 55, input: 5, id:'down',   icon: 'bi-caret-down-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 10,  input: 6, id:'left',  icon: 'bi-caret-left-fill',  cls: 'btn-outline-secondary', square: true},
            {region: 'left', top: 55,  left: 100, input: 7, id:'right', icon: 'bi-caret-right-fill', cls: 'btn-outline-secondary', square: true},

            {region: 'center', top: 105, left:  -15, input: 2, id:'select', text: 'X', cls: 'btn-outline-secondary', small: true},
            {region: 'center', top: 105, right: -15, input: 3, id:'start',  text: 'P',  cls: 'btn-outline-secondary', small: true},

            {region: 'right', top: 55, left: -50, input: 1,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 0,   input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 50,  input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
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
        if (window.localStorage && window.localStorage.playtimeNetplayVirtualGamepad) {
            virtualGamepadVisible = (window.localStorage.playtimeNetplayVirtualGamepad === 'true');
        } else {
            virtualGamepadVisible = _isMobile();
        }
        virtualGamepadShow();
    }

    function virtualGamepadToggle() {
        virtualGamepadVisible = !virtualGamepadVisible;
        if (window.localStorage) {
            window.localStorage.playtimeNetplayVirtualGamepad = (virtualGamepadVisible ? 'true' : 'false');
        }
        virtualGamepadShow();
    }

    function virtualGamepadShow() {
        const virtualGamepadContainer = document.getElementById('virtual-gamepad');
        if (virtualGamepadVisible) {
            virtualGamepadContainer.classList.remove('d-none');
        } else {
            virtualGamepadContainer.classList.add('d-none');
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

    function fullscreen() {
        if (!document.fullscreenEnabled && !document.webkitFullscreenEnabled) {
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

    //https://stackoverflow.com/questions/11381673/detecting-a-mobile-browser
    function _isMobile() {
        let check = false;
        (function(a){if(/(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino/i.test(a)||/1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i.test(a.substr(0,4))) check = true;})(navigator.userAgent||navigator.vendor||window.opera);
        return check;
    }

})();
