(() => {

    const darkModeSwitch = document.getElementById('dark-mode');
    darkModeSwitch.checked = getCurrentColorScheme() === 'dark';
    darkModeSwitch.onchange = () => {
        const theme = darkModeSwitch.checked ? 'dark' : 'light';
        applyDarkTheme(theme);
    };
    applyDarkTheme(getCurrentColorScheme());

    function applyDarkTheme(theme) {
        const html = document.getElementsByTagName('html')[0];
        html.setAttribute('data-bs-theme', theme);

        window.localStorage.__dark_theme = theme;
    }

    function getCurrentColorScheme() {
        const theme = window.localStorage.__dark_theme;
        if (!theme) {
            return isSystemDarkMode() ? 'dark' : 'light';
        }
        return theme;
    }

    function isSystemDarkMode() {
        return (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches);
    }
})();
