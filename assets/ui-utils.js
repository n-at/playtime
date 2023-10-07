(() => {

    window.FlashButtonIcon = (btnId, oldBtnCls, oldIconCls, newBtnCls, newIconCls) => {
        const btn = document.getElementById(btnId);
        if (!btn) {
            return;
        }
        const icon = btn.firstElementChild;
        if (!icon) {
            return;
        }
        if (oldBtnCls) {
            btn.classList.remove(...oldBtnCls);
        }
        if (newBtnCls) {
            btn.classList.add(...newBtnCls)
        }
        if (oldIconCls) {
            icon.classList.remove(...oldIconCls);
        }
        if (newIconCls) {
            icon.classList.add(...newIconCls);
        }
        setTimeout(() => {
            if (newBtnCls) {
                btn.classList.remove(...newBtnCls);
            }
            if (oldBtnCls) {
                btn.classList.add(...oldBtnCls);
            }
            if (newIconCls) {
                icon.classList.remove(...newIconCls);
            }
            if (oldIconCls) {
                icon.classList.add(...oldIconCls);
            }
        }, 1000);
    };

    ///////////////////////////////////////////////////////////////////////////
    // Toast messages
    ///////////////////////////////////////////////////////////////////////////

    /**
     * @param {string} cls
     * @param {string} text
     * @param {number} delay
     * @constructor
     */
    window.ShowToastMessage = function(cls, text, delay = 1000) {
        const el = document.createElement('div');
        el.classList.add('toast', 'align-items-center', 'border-0', 'opacity-50', 'mb-1', `text-bg-${cls}`);
        el.role = 'alert';

        const flexEl = document.createElement('div');
        flexEl.classList.add('d-flex');
        el.append(flexEl);

        const bodyEl = document.createElement('div');
        bodyEl.classList.add('toast-body', 'p-2');
        bodyEl.innerText = text;
        flexEl.append(bodyEl);

        const btnEl = document.createElement('button');
        btnEl.type = 'button';
        btnEl.classList.add('btn-close', 'btn-close-white', 'me-2', 'm-auto');
        btnEl.setAttribute('data-bs-dismiss', 'toast');
        btnEl.ariaLabel = 'Close';
        flexEl.append(btnEl);

        if (!delay) {
            delay = 10000;
        }

        document.getElementById('notifications').append(el);
        bootstrap.Toast.getOrCreateInstance(el, {delay: delay}).show();
    };

    ///////////////////////////////////////////////////////////////////////////
    // Display utils
    ///////////////////////////////////////////////////////////////////////////

    window.NetplayPlayerDisplay = (clientId, hostId, player) => {
        if (clientId === hostId) {
            return 'host';
        }
        if (player === -1) {
            return 'spectator'
        }
        return `player ${player + 1}`;
    };

    window.NetplayLoadClientName = defaultName => {
        if (!window.localStorage) {
            return;
        }
        if (window.localStorage.playtimeNetplayName) {
            return window.localStorage.playtimeNetplayName;
        } else {
            NetplaySaveClientName(defaultName);
            return defaultName;
        }
    };

    window.NetplaySaveClientName = name => {
        if (!window.localStorage) {
            return;
        }
        window.localStorage.playtimeNetplayName = name;
    };

    window.NetplayLoadClientPlayer = (gameId, sessionId, defaultPlayer) => {
        if (!window.localStorage || !window.localStorage.playtimeNetplayPlayer) {
            return defaultPlayer;
        }
        try {
            const players = JSON.parse(window.localStorage.playtimeNetplayPlayer);
            if (players[gameId] !== undefined && players[gameId][sessionId] !== undefined) {
                return players[gameId][sessionId];
            }
        } catch (e) {
            console.error(e);
        }
        window.NetplaySaveClientPlayer(gameId, sessionId, defaultPlayer);
        return defaultPlayer;
    };

    window.NetplaySaveClientPlayer = (gameId, sessionId, player) => {
        if (!window.localStorage) {
            return;
        }
        let players = {};
        if (window.localStorage.playtimeNetplayPlayer) {
            try {
                players = JSON.parse(window.localStorage.playtimeNetplayPlayer);
            } catch (e) {
                console.error(e);
            }
        }
        if (players[gameId] === undefined) {
            players[gameId] = {};
        }
        players[gameId][sessionId] = player;
        window.localStorage.playtimeNetplayPlayer = JSON.stringify(players);
    };

    //https://stackoverflow.com/questions/11381673/detecting-a-mobile-browser
    window.isMobile = () => {
        let check = false;
        (function(a){if(/(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino/i.test(a)||/1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i.test(a.substr(0,4))) check = true;})(navigator.userAgent||navigator.vendor||window.opera);
        return check;
    }

    ///////////////////////////////////////////////////////////////////////////
    // Controls utils
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

    const KeyCodes = {
        0: '',
        8: 'backspace',
        9: 'tab',
        13: 'enter',
        16: 'shift',
        17: 'ctrl',
        18: 'alt',
        19: 'pause/break',
        20: 'caps lock',
        27: 'escape',
        32: 'space',
        33: 'page up',
        34: 'page down',
        35: 'end',
        36: 'home',
        37: 'left arrow',
        38: 'up arrow',
        39: 'right arrow',
        40: 'down arrow',
        45: 'insert',
        46: 'delete',
        48: '0',
        49: '1',
        50: '2',
        51: '3',
        52: '4',
        53: '5',
        54: '6',
        55: '7',
        56: '8',
        57: '9',
        65: 'a',
        66: 'b',
        67: 'c',
        68: 'd',
        69: 'e',
        70: 'f',
        71: 'g',
        72: 'h',
        73: 'i',
        74: 'j',
        75: 'k',
        76: 'l',
        77: 'm',
        78: 'n',
        79: 'o',
        80: 'p',
        81: 'q',
        82: 'r',
        83: 's',
        84: 't',
        85: 'u',
        86: 'v',
        87: 'w',
        88: 'x',
        89: 'y',
        90: 'z',
        91: 'left window key',
        92: 'right window key',
        93: 'select key',
        96: 'numpad 0',
        97: 'numpad 1',
        98: 'numpad 2',
        99: 'numpad 3',
        100: 'numpad 4',
        101: 'numpad 5',
        102: 'numpad 6',
        103: 'numpad 7',
        104: 'numpad 8',
        105: 'numpad 9',
        106: 'multiply',
        107: 'add',
        109: 'subtract',
        110: 'decimal point',
        111: 'divide',
        112: 'f1',
        113: 'f2',
        114: 'f3',
        115: 'f4',
        116: 'f5',
        117: 'f6',
        118: 'f7',
        119: 'f8',
        120: 'f9',
        121: 'f10',
        122: 'f11',
        123: 'f12',
        144: 'num lock',
        145: 'scroll lock',
        186: 'semi-colon',
        187: 'equal sign',
        188: 'comma',
        189: 'dash',
        190: 'period',
        191: 'forward slash',
        192: 'grave accent',
        219: 'open bracket',
        220: 'back slash',
        221: 'close braket',
        222: 'single quote',
    };

    window.ControlsTransformKeyboardCode = code => {
        const label = KeyCodes[code];
        if (label !== undefined) {
            return label;
        } else if (code || code !== '') {
            return code.toString();
        } else {
            return '';
        }
    };

    window.ControlsTransformKeyboardValue = value => {
        for (let keyCode in KeyCodes) {
            if (KeyCodes[keyCode] === value) {
                return parseInt(keyCode);
            }
        }
        const code = parseInt(value);
        if (isNaN(code)) {
            return -1;
        } else {
            return code;
        }
    };

    window.ControlsTransformGamepadCode = code => {
        if (ButtonLabels[code] !== undefined) {
            return ButtonLabels[code];
        } else {
            return `GAMEPAD_${code}`;
        }
    };

})();
