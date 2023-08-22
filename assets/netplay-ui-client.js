(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        window.addEventListener('resize', setupGameDisplaySize);
        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);
        document.getElementById('netplay-control-scheme-save').addEventListener('click', saveControlScheme);
        document.getElementById('netplay-control-scheme-reset').addEventListener('click', resetControlScheme);
        document.getElementById('netplay-virtual-gamepad-toggle').addEventListener('click', virtualGamepadToggle);

        const gameEl = document.getElementById('game');
        gameEl.addEventListener('keydown', controlsButtonDown);
        gameEl.addEventListener('keyup', controlsButtonUp);
        setInterval(controlsPollGamepad, 1000 / 60);

        loadControlScheme();
        virtualGamepadInit();
        virtualGamepadLoad();
        setupGameDisplaySize();
        connectionScreen('Connecting to server');

        netplay = NetplayClient({
            debug: true,
            gameVideoEl: gameEl,
            gameId: window.NetplayGameId,
            sessionId: window.NetplaySessionId,
            host: false,
            turnServerUrl: window.NetplayTurnServerUrl,
            turnServerUser: window.NetplayTurnServerUser,
            turnServerPassword: window.NetplayTurnServerPassword,

            onClientError: errorHandler,
            onWSConnected: wsConnected,
            onRTCConnectionStateChanged: rtcConnectionStateChanged,

            onGreeting: wsGreeting,
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

    function selfNameChanged() {
        document.getElementById('netplay-name').value = netplay.getName();
        document.getElementById('netplay-player').innerText = `${netplay.getName()}: ${_displayPlayer(netplay.getPlayer())}`;
    }

    function selfPlayerChanged() {
        document.getElementById('netplay-player').innerText = `${netplay.getName()}: ${_displayPlayer(netplay.getPlayer())}`;
    }

    function changeSelfName() {
        const name = document.getElementById('netplay-name').value;

        if (name.trim().length === 0 || name.length > 32) {
            window.FlashButtonIcon('netplay-name-change', ['btn-outline-secondary'], ['bi-check'], ['btn-outline-danger'], ['bi-x']);
            return;
        }

        netplay.setName(name);
        saveClientName(name);

        window.FlashButtonIcon('netplay-name-change', ['btn-outline-secondary'], [], ['btn-outline-success'], []);
    }

    function wsGreeting() {
        loadClientName();
    }

    function saveClientName(name) {
        if (!window.localStorage) {
            return;
        }
        window.localStorage.playtimeNetplayName = name;
    }

    function loadClientName() {
        if (!window.localStorage) {
            return;
        }
        if (window.localStorage.playtimeNetplayName) {
            netplay.setName(window.localStorage.playtimeNetplayName);
        } else {
            saveClientName(netplay.getName());
        }
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
        const el = document.getElementById(`netplay-client-${id}`);
        if (el) {
            el.remove();
        }
    }

    function clientNameChanged(id, name) {
        const el = document.getElementById(`netplay-client-${id}-name`);
        if (el) {
            el.innerText = name;
        }
    }

    function clientPlayerChanged(id, player) {
        const el = document.getElementById(`netplay-client-${id}-player`);
        if (el) {
            el.innerText = _displayPlayer(player);
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

            {region: 'right', top: 55, left: -50,  input: 1,  id:'a', text: 'A', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 0,   input: 0,  id:'b', text: 'B', cls: 'btn-outline-secondary', circle: true},
            {region: 'right', top: 55, left: 50, input: 8,  id:'c', text: 'C', cls: 'btn-outline-secondary', circle: true},
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
            //TODO is mobile
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
        connectionScreen('Connected to server');
    }

    function rtcConnectionStateChanged(clientId, state) {
        switch (state) {
            case 'connecting':
                connectionScreen('Connecting to game host');
                break;
            case 'connected':
                connectionScreen(false);
                break;
        }
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
    // Error screen
    ///////////////////////////////////////////////////////////////////////////

    function errorHandler(type, clientId, e) {
        if (['web-socket', 'rtc-offer-send', 'rtc-answer-send', 'rtc-connection', 'rtc-ice-connection', 'rtc-control-channel'].includes(type)) {
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
                window.ShowToastMessage('Game host connection lost');
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
