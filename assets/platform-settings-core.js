(() => {

    const coresWithThreads = [
        'mednafen_psx_hw',
        'melonds',
        'mgba',
        'mupen64plus_next',
        'opera',
        'parallel_n64',
        'pcsx_rearmed',
        'yabause',
    ];

    function setupThreadsControl() {
        const fieldThreads = document.getElementById('threads');
        const fieldCores = document.getElementById('core');

        if (coresWithThreads.includes(fieldCores.value)) {
            fieldThreads.parentElement.classList.remove('d-none');
        } else {
            fieldThreads.parentElement.classList.add('d-none');
            fieldThreads.checked = null;
        }
    }

    window.addEventListener('load', () => {
        document.getElementById('core').addEventListener('change', setupThreadsControl);
        setupThreadsControl();
    });

})();
