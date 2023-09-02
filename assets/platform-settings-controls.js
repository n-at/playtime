(() => {

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

    let modalKeyboard, modalGamepad;

    window.addEventListener('load', () => {
        modalKeyboard = new bootstrap.Modal(document.getElementById('modal-keyboard-control'), {
            keyboard: false,
        });

        modalGamepad = new bootstrap.Modal(document.getElementById('modal-gamepad-control'), {
            keyboard: false,
        });

        ///////////////////////////////////////////////////////////////////////

        document.querySelectorAll('input.keyboard').forEach(input => {
            const title = input.getAttribute('data-title');
            const player = input.getAttribute('data-player');
            input.addEventListener('click', () => {
                document.getElementById('keyboard-control-player').innerText = (parseInt(player)+1).toString();
                document.getElementById('keyboard-control-title').innerText = title;
                document.getElementById('keyboard-control-input').value = input.id;
                document.getElementById('keyboard-control-button').innerText = input.value;
                modalKeyboard.show();
            });
        });

        document.getElementById('keyboard-control-clear').addEventListener('click', () => {
            document.getElementById('keyboard-control-button').innerText = '';
        });

        document.getElementById('keyboard-control-save').addEventListener('click', () => {
            const id = document.getElementById('keyboard-control-input').value;
            document.getElementById(id).value = document.getElementById('keyboard-control-button').innerText;
            modalKeyboard.hide();
        });

        ///////////////////////////////////////////////////////////////////////

        document.querySelectorAll('input.gamepad').forEach(input => {
            input.addEventListener('click', () => {
                const title = input.getAttribute('data-title');
                const player = input.getAttribute('data-player');
                input.addEventListener('click', () => {
                    document.getElementById('gamepad-control-player').innerText = (parseInt(player)+1).toString();
                    document.getElementById('gamepad-control-title').innerText = title;
                    document.getElementById('gamepad-control-input').value = input.id;
                    document.getElementById('gamepad-control-button').innerText = input.value;
                    modalGamepad.show();
                });
            });
        });

        document.getElementById('gamepad-control-clear').addEventListener('click', () => {
            document.getElementById('gamepad-control-button').innerText = '';
        });

        document.getElementById('gamepad-control-save').addEventListener('click', () => {
            const id = document.getElementById('gamepad-control-input').value;
            document.getElementById(id).value = document.getElementById('gamepad-control-button').innerText;
            modalGamepad.hide();
        });

        ///////////////////////////////////////////////////////////////////////

        document.getElementById('modal-keyboard-control').addEventListener('keydown', e => {
            document.getElementById('keyboard-control-button').innerText = e.key.toLowerCase();
            e.preventDefault();
        });

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
                    document.getElementById('gamepad-control-button').innerText = value;
                }
            });
        }, 1000 / 60);
    });

})();
