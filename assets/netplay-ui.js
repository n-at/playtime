(() => {

    window.addEventListener('load', () => {
        const netplayUrl = `${location.protocol}//${location.host}/netplay/${window.NetplayGameId}/${window.NetplaySessionId}`;
        document.getElementById('netplay-url').value = netplayUrl;
        document.getElementById('netplay-url-copy').addEventListener('click', async () => {
            await navigator.clipboard.writeText(netplayUrl);
            window.FlashButtonIcon(
                'netplay-url-copy',
                ['btn-outline-secondary'],
                ['bi-clipboard'],
                ['btn-outline-success'],
                ['bi-clipboard-check'],
            );
        });

        new AwesomeQR.AwesomeQR({
            text: netplayUrl,
            size: 300,
            margin: 5,
            colorDark: '#0d6efd',
            colorLight: '#fff',
            components: {
                data: {
                    scale: 0.5,
                },
                timing: {
                    scale: 0.5,
                    protectors: false,
                },
                alignment: {
                    scale: 0.5,
                    protectors: false,
                },
                cornerAlignment: {
                    scale: 0.5,
                    protectors: true,
                },
            }
        }).draw().then(dataUrl => {
            document.getElementById('netplay-qr').src = dataUrl;
        });
    });

})();
