(() => {

    window.addEventListener('load', () => {
        const loadStateModal = new bootstrap.Modal(document.getElementById('modal-load-state'));

        document.getElementById('btn-save-state').addEventListener('click', async () => {
            const state = await saveState();
            if (state !== null) {
                window.LatestStateUrl = state.StateFileDownloadLink;
            }
        });

        document.getElementById('btn-load-state-latest').addEventListener('click', async () => {
            if (!window.LatestStateUrl) {
                return;
            }
            await loadState(window.LatestStateUrl);
        });

        document.getElementById('btn-load-state').addEventListener('click', async () => {
            const states = await listStates();
            renderSaveStates(states, loadStateModal);
            loadStateModal.show();
        });
    });

    ///////////////////////////////////////////////////////////////////////////

    async function saveState() {
        const state = await EJS_emulator.gameManager.getState();
        const screenshot = EJS_emulator.gameManager.screenshot();
        return await uploadState(state, screenshot);
    }

    async function uploadState(state, screenshot) {
        const formData = new FormData();
        formData.append('state', new Blob([state]));
        formData.append('screenshot', new Blob([screenshot]));

        try {
            const url = `/games/save-states/${GameId}/upload`;
            const response = await fetch(url, {
                method: 'post',
                body: formData,
            });
            return await response.json();
        } catch (e) {
            console.log('save state upload error', e)
            return null;
        }
    }

    async function loadState(stateUrl) {
        try {
            const result = await fetch(stateUrl);
            const data = await result.arrayBuffer();
            EJS_emulator.gameManager.loadState(new Uint8Array(data));
        } catch (e) {
            console.error("save state load error", e);
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
            container.classList.add('col-4');

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
