import Tags from '/assets/node_modules/bootstrap5-tags/tags.js'

(() => {

    let searchField = null;
    let tagsField = null;

    window.addEventListener('load', () => {
        searchField = document.getElementById('games-search');
        searchField.addEventListener('keyup', doSearch);

        Tags.init('#games-tags');
        const tagsEl = document.getElementById('games-tags');
        tagsField = Tags.getInstance(tagsEl);
        tagsField.setConfig('onSelectItem', doSearch);
        tagsField.setConfig('onClearItem', doSearch);

        document.querySelectorAll('span.game-tag').forEach(tagEl => {
            const tagValue = tagEl.innerText.trim();
            tagEl.addEventListener('click', () => {
                if (!tagsField.getSelectedValues().includes(tagValue)) {
                    tagsField.addItem(tagValue);
                    doSearch();
                }
            });
        });

        loadSearchParams();
    });

    function doSearch() {
        const searchText = searchField.value;
        const searchTags = tagsField.getSelectedValues();

        search(searchText, searchTags);

        if (window.sessionStorage) {
            sessionStorage._games_search_text = searchText;
            sessionStorage._games_search_tags = JSON.stringify(searchTags);
        }
    }

    function loadSearchParams() {
        if (!window.sessionStorage) {
            return;
        }
        if (window.sessionStorage._games_search_text) {
            searchField.value = window.sessionStorage._games_search_text;
        }
        if (window.sessionStorage._games_search_tags) {
            JSON.parse(window.sessionStorage._games_search_tags).forEach(tag => tagsField.addItem(tag));
        }
        doSearch();
    }

    function search(text, tags) {
        const searchText = escapeRegExpChars(text.trim());
        const re = new RegExp(searchText, 'ig');

        let gameFound = false;

        document.querySelectorAll('div.game').forEach(gameEl => {
            const gameName = gameEl.querySelector('a.game-name').innerText;
            const textMatched = !text || gameName.match(re);

            let tagsCount = 0;
            gameEl.querySelectorAll('span.game-tag').forEach(tagEl => {
                const tag = tagEl.innerText.trim();
                if (tags.includes(tag)) {
                    tagsCount++;
                }
            });
            let tagsMatched = !tags || tags.length === 0 || tagsCount === tags.length;

            if (textMatched && tagsMatched) {
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
