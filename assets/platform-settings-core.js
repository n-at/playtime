(() => {

    const coresWithThreads = [
        'a5200',
        'beetle_vb',
        'desmume2015',
        'desmume',
        'fbalpha2012_cps1',
        'fbalpha2012_cps2',
        'fbneo',
        'fceumm',
        'gambatte',
        'gearcoleco',
        'handy',
        'mame2003',
        'mame2003_plus',
        'mednafen_ngp',
        'mednafen_pce',
        'mednafen_pcfx',
        'mednafen_psx_hw',
        'mednafen_wswan',
        'melonds',
        'mgba',
        'mupen64plus_next',
        'nestopia',
        'opera',
        'parallel_n64',
        'pcsx_rearmed',
        'picodrive',
        'prosystem',
        'snes9x',
        'stella2014',
        'vice_x64',
        'vice_x64sc',
        'vice_x128',
        'vice_xpet',
        'vice_xplus4',
        'vice_xvic',
        'virtualjaguar',
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

    function showCoreOptions() {
        const value = document.getElementById('core').value;
        document.querySelectorAll('.core-options').forEach(it => it.classList.add('d-none'));
        document.getElementById(`core-options-${value}`).classList.remove('d-none');
    }

    window.addEventListener('load', () => {
        document.getElementById('core').addEventListener('change', () => {
            setupThreadsControl();
            showCoreOptions();
        });
        setupThreadsControl();
        showCoreOptions();
    });

})();
