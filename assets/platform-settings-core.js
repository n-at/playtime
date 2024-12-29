(() => {

    function showCoreOptions() {
        const value = document.getElementById('core').value;
        document.querySelectorAll('.core-options').forEach(it => it.classList.add('d-none'));
        document.getElementById(`core-options-${value}`).classList.remove('d-none');
    }

    window.addEventListener('load', () => {
        document.getElementById('core').addEventListener('change', () => {
            showCoreOptions();
        });
        showCoreOptions();
    });

})();
