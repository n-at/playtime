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

            updateKeyDisplay('keyboard', player, button);

            input.addEventListener('click', () => {
                document.getElementById('keyboard-control-player-title').innerText = (parseInt(player)+1).toString();
                document.getElementById('keyboard-control-title').innerText = title;
                document.getElementById('keyboard-control-display').innerText = input.value ? input.value : 'Not set';
                document.getElementById('keyboard-control-code').value = getValue('keyboard', player, button);
                document.getElementById('keyboard-control-button').value = button;
                document.getElementById('keyboard-control-player').value = player;
                modalKeyboard.show();
            });
        });

        document.getElementById('keyboard-control-clear').addEventListener('click', () => {
            document.getElementById('keyboard-control-display').innerText = 'Not set';
            document.getElementById('keyboard-control-code').value = '';
        });

        document.getElementById('keyboard-control-save').addEventListener('click', () => {
            const button = document.getElementById('keyboard-control-button').value;
            const player = document.getElementById('keyboard-control-player').value;
            setValue('keyboard', player, button, document.getElementById('keyboard-control-code').value);
            modalKeyboard.hide();
        });

        document.getElementById('modal-keyboard-control').addEventListener('keydown', e => {
            const value = window.ControlsTransformKeyboardCode(e.keyCode);
            document.getElementById('keyboard-control-code').value = value;
            document.getElementById('keyboard-control-display').innerText = value;
            e.preventDefault();
        });

        ///////////////////////////////////////////////////////////////////////

        document.querySelectorAll('input.gamepad').forEach(input => {
            const title = input.getAttribute('data-title');
            const player = input.getAttribute('data-player');
            const button = input.getAttribute('data-btn');

            updateKeyDisplay('gamepad', player, button);

            input.addEventListener('click', () => {
                document.getElementById('gamepad-control-player-title').innerText = (parseInt(player)+1).toString();
                document.getElementById('gamepad-control-title').innerText = title;
                document.getElementById('gamepad-control-display').innerText = input.value ? input.value : 'Not set';
                document.getElementById('gamepad-control-input').value = getValue('gamepad', player, button);
                document.getElementById('gamepad-control-button').value = button;
                document.getElementById('gamepad-control-player').value = player;

                modalGamepad.show();
            });
        });

        document.getElementById('gamepad-control-clear').addEventListener('click', () => {
            document.getElementById('gamepad-control-display').innerText = 'Not set';
            document.getElementById('gamepad-control-input').value = '';
        });

        document.getElementById('gamepad-control-save').addEventListener('click', () => {
            const button = document.getElementById('gamepad-control-button').value;
            const player = document.getElementById('gamepad-control-player').value;
            setValue('gamepad', player, button, document.getElementById('gamepad-control-input').value);
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
                        value = window.ControlsTransformGamepadCode(buttonIdx);
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

    function updateKeyDisplay(input, player, btn) {
        const displayEl = document.querySelector(`input[type="text"][data-input="${input}"][data-player="${player}"][data-btn="${btn}"]`);
        const valueEl = document.querySelector(`input[type="hidden"][data-input="${input}"][data-player="${player}"][data-btn="${btn}"]`);

        if (!displayEl || !valueEl) {
            return;
        }

        if (input === 'gamepad') {
            displayEl.value = valueEl.value;
        }
        if (input === 'keyboard') {
            displayEl.value = valueEl.value;
        }
    }

    function getValue(input, player, button) {
        const el = document.querySelector(`input[type="hidden"][data-input="${input}"][data-player="${player}"][data-btn="${button}"]`);
        if (el) {
            return el.value;
        } else {
            return '';
        }
    }

    function setValue(input, player, button, value) {
        const el = document.querySelector(`input[type="hidden"][data-input="${input}"][data-player="${player}"][data-btn="${button}"]`);
        if (el) {
            el.value = value;
        }
        updateKeyDisplay(input, player, button);
    }

    window.PlatformSettingsControls = {
        getValue,
        setValue,
    };

})();
