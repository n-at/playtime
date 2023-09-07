(() => {

    window.addEventListener('load', () => {
        const loadStateModal = new bootstrap.Modal(document.getElementById('modal-load-state'));

        document.getElementById('btn-save-state').addEventListener('click', saveState);
        document.getElementById('btn-load-state-latest').addEventListener('click', loadLatestState);
        document.getElementById('btn-load-state').addEventListener('click', async () => {
            const states = await listStates();
            renderSaveStates(states, loadStateModal);
            loadStateModal.show();
        });

        document.addEventListener('keydown', keyboardStateControl);
        document.getElementById('game').addEventListener('keydown', keyboardStateControl);
        setInterval(gamepadStateControl, 1000/60);
    });

    ///////////////////////////////////////////////////////////////////////////

    async function keyboardStateControl(e) {
        if (e.repeat) {
            return;
        }
        const key = e.key.toLowerCase();
        for (let playerIdx = 0; playerIdx < 4; playerIdx++) {
            if (key === PlaytimeControls[playerIdx].load.value) {
                e.preventDefault();
                await loadLatestState();
                return;
            }
            if (key === PlaytimeControls[playerIdx].save.value) {
                e.preventDefault();
                await saveState();
                return;
            }
        }
    }

    async function gamepadStateControl() {
        if (!navigator.getGamepads) {
            return;
        }
        navigator.getGamepads().forEach(gamepad => {
            if (!gamepad) {
                return;
            }
            let value = null;
            gamepad.buttons.forEach((button, buttonIdx) => {
                if (button.pressed) {
                    value = buttonIdx;
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
                value = `${axisName}:${axisValue}`;
            });
            if (value === null) {
                return;
            }
            for (let playerIdx = 0; playerIdx < 4; playerIdx++) {
                if (value === PlaytimeControls[playerIdx].load.value2) {
                    loadLatestState();
                    return;
                }
                if (value === PlaytimeControls[playerIdx].save.value2) {
                    saveState();
                    return;
                }
            }
        });
    }

    ///////////////////////////////////////////////////////////////////////////

    async function saveState() {
        const state = await EJS_emulator.gameManager.getState();
        const screenshot = EJS_emulator.gameManager.screenshot();
        const response = await uploadState(state, screenshot);
        if (response) {
            window.LatestStateUrl = response.StateFileDownloadLink
        }
    }

    async function uploadState(state, screenshot) {
        const formData = new FormData();
        formData.append('state', new Blob([state]));
        formData.append('screenshot', new Blob([screenshot]));
        formData.append('_playtime_csrf', window._csrf);

        try {
            const url = `/games/save-states/${GameId}/upload`;
            const response = await fetch(url, {
                method: 'post',
                body: formData,
            });
            const json = await response.json();

            stateSaveSuccess();

            return json;
        } catch (e) {
            console.log('save state upload error', e);
            stateSaveError();
            return null;
        }
    }

    async function loadLatestState() {
        if (!window.LatestStateUrl) {
            return;
        }
        await loadState(window.LatestStateUrl);
    }

    async function loadState(stateUrl) {
        try {
            const result = await fetch(stateUrl);
            const data = await result.arrayBuffer();
            EJS_emulator.gameManager.loadState(new Uint8Array(data));
            stateLoadSuccess();
        } catch (e) {
            console.error("save state load error", e);
            stateLoadError();
        }
    }

    async function listStates() {
        try {
            const url = `/games/save-states/${GameId}/list`;
            const response = await fetch(url);
            return await response.json();
        } catch (e) {
            console.error("list states error", e);
            return [];
        }
    }

    ///////////////////////////////////////////////////////////////////////////

    function stateSaveSuccess() {
        window.FlashButtonIcon(
            'btn-save-state',
            ['btn-outline-secondary'],
            ['bi-box-arrow-down'],
            ['btn-outline-success'],
            ['bi-check']
        );
    }

    function stateSaveError() {
        window.FlashButtonIcon(
            'btn-save-state',
            ['btn-outline-secondary'],
            ['bi-box-arrow-down'],
            ['btn-outline-danger'],
            ['bi-x']
        );
    }

    function stateLoadSuccess() {
        window.FlashButtonIcon(
            'btn-load-state-latest',
            ['btn-outline-secondary'],
            ['bi-box-arrow-up'],
            ['btn-outline-success'],
            ['bi-check']
        );
    }

    function stateLoadError() {
        window.FlashButtonIcon(
            'btn-load-state-latest',
            ['btn-outline-secondary'],
            ['bi-box-arrow-up'],
            ['btn-outline-danger'],
            ['bi-x']
        );
    }

    ///////////////////////////////////////////////////////////////////////////

    function renderSaveStates(states, modal) {
        const stateList = document.getElementById('load-state-list');
        stateList.innerHTML = '';

        const emptyState = document.getElementById('load-state-empty');

        if (!states || states.length === 0) {
            emptyState.classList.remove('d-none');
            return;
        } else {
            emptyState.classList.add('d-none');
        }

        states.forEach(state => {
            const container = document.createElement('div');
            container.classList.add('col-12', 'col-sm-6', 'col-lg-4');

            const card = document.createElement('div');
            card.classList.add('card', 'mb-3');

            const img = document.createElement('img');
            img.classList.add('card-img-top')
            img.src = state.ScreenshotDownloadLink;
            img.alt = 'Screenshot';

            const body = document.createElement('div');
            body.classList.add('card-body', 'text-center');

            const button = document.createElement('button');
            button.classList.add('btn', 'btn-sm', 'btn-outline-secondary');
            button.type = 'button';
            button.innerText = new Date(state.Created).toLocaleString();
            button.onclick = () => {
                modal.hide();
                loadState(state.StateFileDownloadLink).then();
            };

            body.append(button);
            card.append(img, body);
            container.append(card);
            stateList.append(container);
        });
    }

})();
