(() => {

    EJS_onSaveState = e => {
        uploadState(e.state, e.screenshot)
    };

    EJS_onGameStart = () => {
        console.log('game started');
    };

    window.addEventListener('load', () => {
        document.getElementById('btn-save-state').addEventListener('click', saveState);
    });

    ///////////////////////////////////////////////////////////////////////////

    async function saveState() {
        const state = await EJS_emulator.gameManager.getState();
        const screenshot = EJS_emulator.gameManager.screenshot();
        uploadState(state, screenshot);
    }

    function uploadState(state, screenshot) {
        const formData = new FormData();
        formData.append('state', new Blob([state]));
        formData.append('screenshot', new Blob([screenshot]));

        const url = `/games/save-states/${GameId}/upload`;

        fetch(url, {
            method: 'post',
            body: formData,
        }).then(result => {
            result.text().then(stateId => {
                console.log('save state result', stateId);
            });
        }).catch(e => {
            console.error('save state upload error', e);
        })
    }

})();
