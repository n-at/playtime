(() => {

    window.addEventListener('load', () => {
        const searchField = document.getElementById('games-search');
        searchField.addEventListener('keyup', () => {
            search(searchField.value);
        });
    });

    function search(text) {
        const searchText = escapeRegExpChars(text.trim());
        const re = new RegExp(searchText, 'ig');

        let platformFound = false;

        document.querySelectorAll('div.game-list-platform').forEach(gameListEl => {
            let gameFound = false;
            gameListEl.querySelectorAll('div.game').forEach(gameEl => {
                const gameName = gameEl.querySelector('span.game-name').innerText;
                if (!text || gameName.match(re)) {
                    gameEl.classList.remove('d-none');
                    gameFound = true;
                } else {
                    gameEl.classList.add('d-none');
                }
            });
            if (!gameFound) {
                gameListEl.classList.add('d-none');
            } else {
                gameListEl.classList.remove('d-none');
                platformFound = true;
            }
        });

        notFound(!platformFound);
    }

    function notFound(value) {
        const el = document.getElementById('games-not-found');
        if (value) {
            el.classList.remove('d-none');
        } else {
            el.classList.add('d-none');
        }
    }

    function escapeRegExpChars(str) {
        if (!str) {
            return '';
        }
        const escaped = str.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
        return escaped.replace(/\s+/g, '\\s+');
    }

})();
