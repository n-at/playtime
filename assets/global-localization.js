(() => {

    window.addEventListener('load', () => {
        document.querySelectorAll('button.localization').forEach(el => {
            const code = el.getAttribute('data-lang');

            el.addEventListener('click', () => {
                const expires = new Date(Date.now() + 30 * 24 * 60 * 60 * 1000);
                document.cookie = `playtime-l10n=${code}; path=/; expires=${expires.toUTCString()}`;
                location.reload();
            });
        });
    });

})();
