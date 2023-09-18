(() => {

    let modalKeyboardEl;
    let modalKeyboard, modalGamepad;

    window.addEventListener('load', () => {
        modalKeyboardEl = document.getElementById('modal-keyboard-control');
        modalKeyboard = new bootstrap.Modal(modalKeyboardEl, {
            keyboard: false,
        });

        modalGamepad = new bootstrap.Modal(document.getElementById('modal-gamepad-control'), {
            keyboard: false,
        });

        ///////////////////////////////////////////////////////////////////////

        document.querySelectorAll('input.keyboard').forEach(input => {
            const title = input.getAttribute('data-title');
            const player = input.getAttribute('data-player');
            const button = input.getAttribute('data-btn');

            updateKeyDisplay('keyboard', button);

            input.addEventListener('click', () => {
                document.getElementById('keyboard-control-player').innerText = (parseInt(player)+1).toString();
                document.getElementById('keyboard-control-title').innerText = title;
                document.getElementById('keyboard-control-display').innerText = input.value;
                document.getElementById('keyboard-control-code').value = getValue('keyboard', button);
                document.getElementById('keyboard-control-button').value = button;
                modalKeyboard.show();
            });
        });

        document.getElementById('keyboard-control-clear').addEventListener('click', () => {
            document.getElementById('keyboard-control-display').innerText = '';
            document.getElementById('keyboard-control-code').value = '';
        });

        document.getElementById('keyboard-control-save').addEventListener('click', () => {
            const button = document.getElementById('keyboard-control-button').value;
            setValue('keyboard', button, document.getElementById('keyboard-control-code').value);
            modalKeyboard.hide();
        });

        document.getElementById('modal-keyboard-control').addEventListener('keydown', e => {
            document.getElementById('keyboard-control-code').value = e.keyCode;
            document.getElementById('keyboard-control-display').innerText = keyboardDisplay(e.keyCode);
            e.preventDefault();
        });

        ///////////////////////////////////////////////////////////////////////

        document.querySelectorAll('input.gamepad').forEach(input => {
            const title = input.getAttribute('data-title');
            const player = input.getAttribute('data-player');
            const button = input.getAttribute('data-btn');

            updateKeyDisplay('gamepad', button);

            input.addEventListener('click', () => {
                document.getElementById('gamepad-control-player').innerText = (parseInt(player)+1).toString();
                document.getElementById('gamepad-control-title').innerText = title;
                document.getElementById('gamepad-control-display').innerText = input.value;
                document.getElementById('gamepad-control-input').value = getValue('gamepad', button);
                document.getElementById('gamepad-control-button').value = button;
                modalGamepad.show();
            });
        });

        document.getElementById('gamepad-control-clear').addEventListener('click', () => {
            document.getElementById('gamepad-control-display').innerText = '';
            document.getElementById('gamepad-control-input').value = '';
        });

        document.getElementById('gamepad-control-save').addEventListener('click', () => {
            const button = document.getElementById('gamepad-control-button').value;
            setValue('gamepad', button, document.getElementById('gamepad-control-input').value);
            modalGamepad.hide();
        });

        ///////////////////////////////////////////////////////////////////////

        setInterval(() => {
            if (!navigator.getGamepads) {
                return;
            }

            navigator.getGamepads().forEach(gamepad => {
                if (!gamepad) {
                    return;
                }

                let active = false;
                let value = null;

                gamepad.buttons.forEach((button, buttonIdx) => {
                    if (button.pressed) {
                        if (ButtonLabels[buttonIdx] !== undefined) {
                            value = ButtonLabels[buttonIdx];
                        } else {
                            value = `GAMEPAD_${buttonIdx}`;
                        }
                        active = true;
                    }
                });
                gamepad.axes.forEach((axisValue, axisIdx) => {
                    const axisName = ['LEFT_STICK_X', 'LEFT_STICK_Y', 'RIGHT_STICK_X', 'RIGHT_STICK_Y'][axisIdx];
                    if (!axisName) {
                        return;
                    }

                    axisValue = Math.round(axisValue);
                    if (axisValue === 0) {
                        return;
                    }

                    active = true;
                    value = `${axisName}:${axisValue > 0 ? '+1' : '-1'}`;
                });

                if (active) {
                    document.getElementById('gamepad-control-name').innerText = gamepad.id;
                    document.getElementById('gamepad-control-display').innerText = value;
                    document.getElementById('gamepad-control-input').value = value;
                }
            });
        }, 1000 / 60);
    });

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
        Tab            : "9",
        NumpadEqual    : "12",
        Enter          : "13",
        Shift          : "16",
        Control        : "17",
        Alt            : "18",
        Pause          : "19",
        CapsLock       : "20",
        Escape         : "27",
        Space          : "32",
        PageUp         : "33",
        PageDown       : "34",
        End            : "35",
        Home           : "36",
        ArrowLeft      : "37",
        ArrowUp        : "38",
        ArrowRight     : "39",
        ArrowDown      : "40",
        PrintScreen    : "44",
        Insert         : "45",
        Delete         : "46",
        Digit0         : "48",
        Digit1         : "49",
        Digit2         : "50",
        Digit3         : "51",
        Digit4         : "52",
        Digit5         : "53",
        Digit6         : "54",
        Digit7         : "55",
        Digit8         : "56",
        Digit9         : "57",
        A              : "65",
        B              : "66",
        C              : "67",
        D              : "68",
        E              : "69",
        F              : "70",
        G              : "71",
        H              : "72",
        I              : "73",
        J              : "74",
        K              : "75",
        L              : "76",
        M              : "77",
        N              : "78",
        O              : "79",
        P              : "80",
        Q              : "81",
        R              : "82",
        S              : "83",
        T              : "84",
        U              : "85",
        V              : "86",
        W              : "87",
        X              : "88",
        Y              : "89",
        Z              : "90",
        MetaLeft       : "91",
        MetaRight      : "92",
        ContextMenu    : "93",
        Numpad0        : "96",
        Numpad1        : "97",
        Numpad2        : "98",
        Numpad3        : "99",
        Numpad4        : "100",
        Numpad5        : "101",
        Numpad6        : "102",
        Numpad7        : "103",
        Numpad8        : "104",
        Numpad9        : "105",
        NumpadMultiply : "106",
        NumpadAdd      : "107",
        NumpadSubtract : "109",
        Decimal        : "110",
        NumpadDivide   : "111",
        F1             : "112",
        F2             : "113",
        F3             : "114",
        F4             : "115",
        F5             : "116",
        F6             : "117",
        F7             : "118",
        F8             : "119",
        F9             : "120",
        F10            : "121",
        F11            : "122",
        F12            : "123",
        F13            : "124",
        F14            : "125",
        F15            : "126",
        F16            : "127",
        F17            : "128",
        F18            : "129",
        F19            : "130",
        F20            : "131",
        F21            : "132",
        F22            : "133",
        F23            : "134",
        F24            : "135",
        NumLock        : "144",
        ScrollLock     : "145",
        Semicolon      : "186",
        Equal          : "187",
        Comma          : "188",
        Minus          : "189",
        Period         : "190",
        Backquote      : "192",
        IntlRo         : "193",
        NumpadComma    : "194",
        BracketLeft    : "219",
        Backslash      : "220",
        BracketRight   : "221",
        Quote          : "222",
        IntlYen        : "255",
    };

    const KeyLabels = {};
    for (let key in KeyCodes) {
        KeyLabels[KeyCodes[key]] = key;
    }

    function updateKeyDisplay(input, btn) {
        const displayEl = document.querySelector(`input[type="text"][data-input="${input}"][data-btn="${btn}"]`);
        const valueEl = document.querySelector(`input[type="hidden"][data-input="${input}"][data-btn="${btn}"]`);

        if (!displayEl || !valueEl) {
            return;
        }

        if (input === 'gamepad') {
            displayEl.value = valueEl.value;
        }

        if (input === 'keyboard') {
            displayEl.value = keyboardDisplay(valueEl.value);
        }
    }

    function keyboardDisplay(code) {
        const label = KeyLabels[code];
        if (label !== undefined) {
            return `${label} (${code})`;
        } else if (code) {
            return `Key #${code}`;
        } else {
            return '';
        }
    }

    function getValue(input, button) {
        const el = document.querySelector(`input[type="hidden"][data-input="${input}"][data-btn="${button}"]`);
        if (el) {
            return el.value;
        } else {
            return '';
        }
    }

    function setValue(input, button, value) {
        const el = document.querySelector(`input[type="hidden"][data-input="${input}"][data-btn="${button}"]`);
        if (el) {
            el.value = value;
        }
        updateKeyDisplay(input, button);
    }

})();
