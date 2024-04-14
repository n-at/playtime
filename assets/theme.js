(() => {

    const ThemeSystemDefault = 'system';
    const ThemeLight = 'light';
    const ThemeDark = 'dark';

    let currentTheme = ThemeSystemDefault;

    function applyTheme(theme) {
        const html = document.getElementsByTagName('html')[0];
        html.setAttribute('data-bs-theme', theme);
    }

    function setCurrentTheme(theme) {
        currentTheme = theme;
        window.localStorage.__theme = theme;
        if (theme === ThemeSystemDefault) {
            applyTheme(isSystemDarkMode() ? ThemeDark : ThemeLight);
        } else {
            applyTheme(theme);
        }

        const themeSwitcherIconEl = document.querySelector('#theme i.bi');
        themeSwitcherIconEl.classList.remove('bi-circle-half', 'bi-sun', 'bi-moon-stars');
        switch (theme) {
            case ThemeSystemDefault:
                themeSwitcherIconEl.classList.add('bi-circle-half');
                break;
            case ThemeLight:
                themeSwitcherIconEl.classList.add('bi-sun');
                break;
            case ThemeDark:
                themeSwitcherIconEl.classList.add('bi-moon-stars');
                break;
        }
    }

    function getCurrentTheme() {
        const theme = window.localStorage.__theme;
        return theme ? theme : ThemeSystemDefault;
    }

    function isSystemDarkMode() {
        return (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches);
    }

    window.addEventListener('load', () => {
        document.getElementById('theme-system-default').addEventListener('click', () => {
            setCurrentTheme(ThemeSystemDefault);
        });
        document.getElementById('theme-light').addEventListener('click', () => {
            setCurrentTheme(ThemeLight);
        });
        document.getElementById('theme-dark').addEventListener('click', () => {
            setCurrentTheme(ThemeDark);
        });

        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', event => {
            const newTheme = event.matches ? ThemeDark : ThemeLight;
            if (currentTheme === ThemeSystemDefault) {
                applyTheme(newTheme);
            }
        });
    });

    setCurrentTheme(getCurrentTheme());
})();
