(() => {

    let netplay = null;

    window.addEventListener('load', () => {
        window.addEventListener('resize', setupGameDisplaySize);
        document.getElementById('netplay-name-change').addEventListener('click', changeSelfName);
        document.getElementById('netplay-control-scheme-save').addEventListener('click', saveControlScheme);
        document.getElementById('netplay-control-scheme-reset').addEventListener('click', resetControlScheme);

        const gameEl = document.getElementById('game');
        gameEl.addEventListener('keydown', controlsButtonDown);
        gameEl.addEventListener('keyup', controlsButtonUp);
        setInterval(controlsPollGamepad, 1000 / 60);

        loadControlScheme();
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
