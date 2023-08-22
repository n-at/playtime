(() => {

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

        window.addEventListener('keydown', e => {
            document.getElementById('keyboard-control-button').innerText = e.key.toLowerCase();
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
                        value = buttonIdx;
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
                    value = `${axisName}:${Math.round(axisValue)}`;
                });

                if (active) {
                    document.getElementById('gamepad-control-name').innerText = gamepad.id;
                    document.getElementById('gamepad-control-button').innerText = value;
                }
            });
        }, 1000 / 60);
    });

})();
