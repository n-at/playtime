(() => {

    EJS_onSaveState = e => {
        uploadState(e.state, e.screenshot)
    };

    EJS_onGameStart = () => {
        console.log('game started');
    };

    window.addEventListener('load', () => {
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

})();
