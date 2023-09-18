(() => {

    const ControlSchemeMapping = {
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
        //24: 'quick-save-state', (in PlaytimeControls)
        //25: 'quick-load-state', (in PlaytimeControls)
        //26: 'change-state-slot', (unused)
        27: 'fast-forward',
        28: 'rewind',
        29: 'slow-motion',
    };

    const ControlSchemeMappingState = {
        'save': 'quick-save-state',
        'load': 'quick-load-state',
    };

    window.addEventListener('load', () => {
        if (window.ControlSchemeVariant === 'host') {
            document.getElementById('play-control-scheme-save').addEventListener('click', saveHostControlScheme);
            document.getElementById('play-control-scheme-reset').addEventListener('click', resetHostControlScheme);
            loadHostControlScheme();
            resetHostControlScheme();
        } else if (window.ControlSchemeVariant === 'client') {
            document.getElementById('netplay-control-scheme-save').addEventListener('click', saveClientControlScheme);
            document.getElementById('netplay-control-scheme-reset').addEventListener('click', resetClientControlScheme);
            loadClientControlScheme();
            resetClientControlScheme();
        }
    });

    ///////////////////////////////////////////////////////////////////////////

    function loadHostControlScheme() {
        //nothing here - already loaded
    }

    function saveHostControlScheme() {
        window.EJS_emulator.controls = {};
        window.PlaytimeControls = {};

        ['0', '1', '2' , '3'].forEach(player => {
            console.log(gatherInputs(player, ControlSchemeMapping));
            window.EJS_emulator.controls[player.toString()] = gatherInputs(player, ControlSchemeMapping);
            window.PlaytimeControls[player.toString()] = gatherInputs(player, ControlSchemeMappingState);
        });

        uploadHostControlScheme();
    }

    function resetHostControlScheme() {
        ['0', '1', '2' , '3'].forEach(player => {
            if (window.EJS_emulator && window.EJS_emulator.controls) {
                assignInputs(player, window.EJS_emulator.controls[player], ControlSchemeMapping);
            } else {
                assignInputs(player, window.EJS_defaultControls[player], ControlSchemeMapping);
            }
            assignInputs(player, window.PlaytimeControls[player], ControlSchemeMappingState);
        });
    }

    function uploadHostControlScheme() {
        const formData = new FormData();
        const inputs = gatherInputsToUpload();
        for (let key in inputs) {
            formData.append(key, inputs[key]);
        }
        formData.append('_playtime_csrf', window._csrf);

        fetch(`/games/controls/${GameId}/save`, {
            method: 'post',
            body: formData,
        })
            .then(() => {
                window.FlashButtonIcon('btn-control-scheme', ['btn-outline-secondary'], ['bi-controller'], ['btn-outline-success'], ['bi-check']);
            })
            .catch(e => {
                window.FlashButtonIcon('btn-control-scheme', ['btn-outline-secondary'], ['bi-controller'], ['btn-outline-danger'], ['bi-x']);
                console.error('game controls save error', e);
            });
    }

    ///////////////////////////////////////////////////////////////////////////

    function loadClientControlScheme() {
        if (!window.localStorage || !window.localStorage.playtimeNetplayControls) {
            return;
        }
        try {
            let controls = JSON.parse(window.localStorage.playtimeNetplayControls)[window.GamePlatform];
            if (controls) {
                window.ControlScheme = controls;
            }
        } catch (e) {
            console.error('Unable to load controls', e);
        }
    }

    function saveClientControlScheme() {
        window.ControlScheme = gatherInputs('0', ControlSchemeMapping);

        if (!window.localStorage) {
            return;
        }

        let controls;
        try {
            controls = JSON.parse(window.localStorage.playtimeNetplayControls);
        } catch (e) {
            controls = {};
        }
        controls[window.GamePlatform] = window.ControlScheme;
        window.localStorage.playtimeNetplayControls = JSON.stringify(controls);
    }

    function resetClientControlScheme() {
        assignInputs('0', window.ControlScheme, ControlSchemeMapping);
    }

    ///////////////////////////////////////////////////////////////////////////

    /**
     * Assign controls values to form inputs
     *
     * @param {string} player
     * @param {Object} controls
     * @param {Object} mapping
     */
    function assignInputs(player, controls, mapping) {
        for (let buttonId in controls) {
            const buttonName = mapping[buttonId];
            PlatformSettingsControls.setValue('keyboard', player, buttonName, controls[buttonId].keycode ? controls[buttonId].keycode : '');
            PlatformSettingsControls.setValue('gamepad', player, buttonName, controls[buttonId].value2 ? controls[buttonId].value2 : '');
        }
    }

    /**
     * Gather values from form inputs
     *
     * @param {string} player
     * @param {Object} mapping
     * @returns {{}}
     */
    function gatherInputs(player, mapping) {
        const controls = {};

        for (let buttonId in mapping) {
            const buttonName = mapping[buttonId];

            if (!controls[buttonId]) {
                controls[buttonId] = {
                    keycode: null,
                    value2: null,
                };
            }

            controls[buttonId].keycode = parseInt(PlatformSettingsControls.getValue('keyboard', player, buttonName));
            if (isNaN(controls[buttonId].keycode)) {
                controls[buttonId].keycode = undefined;
            }

            controls[buttonId].value2 = PlatformSettingsControls.getValue('gamepad', player, buttonName);
        }

        return controls;
    }

    function gatherInputsToUpload() {
        const inputs = {};

        document.querySelectorAll('input[type="hidden"][data-input="keyboard"], input[type="hidden"][data-input="gamepad"]').forEach(input => {
            inputs[input.name] = input.value;
        });

        return inputs;
    }

})();
