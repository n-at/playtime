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

        let gameFound = false;

        document.querySelectorAll('div.game').forEach(gameEl => {
            const gameName = gameEl.querySelector('a.game-name').innerText;
            if (!text || gameName.match(re)) {
                gameEl.classList.remove('d-none');
                gameFound = true;
            } else {
                gameEl.classList.add('d-none');
            }
        });

        notFound(!gameFound);
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
