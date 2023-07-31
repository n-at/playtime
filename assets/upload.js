(() => {
    window.addEventListener('load', () => {
        const dropOverlay = document.getElementById('drop-overlay');
        const form = document.getElementById('upload');
        const input = document.getElementById('games');

        window.addEventListener('dragover', e => {
            e.preventDefault();
            dropOverlay.classList.remove('d-none');
        });
        dropOverlay.addEventListener('dragleave', () => {
            dropOverlay.classList.add('d-none');
        });

        window.addEventListener('drop', e => {
            dropOverlay.classList.add('d-none');

            if (!e.dataTransfer || !e.dataTransfer.items || e.dataTransfer.items.length === 0) {
                return;
            }

            e.preventDefault();

            input.files = e.dataTransfer.files;
            form.submit();
        });

        document.getElementById('btn-upload').addEventListener('click', () => {
            input.click();
        });

        input.addEventListener('change', () => {
            form.submit();
        })
    });
})();
