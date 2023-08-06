(() => {

    //https://gist.github.com/nateps/1172490

    const ua = navigator.userAgent;
    const iphone = ~ua.indexOf('iPhone') || ~ua.indexOf('iPod');
    const ipad = ~ua.indexOf('iPad');
    const ios = iphone || ipad;
    const fullscreen = !!window.navigator.standalone;
    const android = ~ua.indexOf('Android');
    let lastWidth = 0

    const body = document.getElementsByTagName('body')[0];

    if (android) {
        window.addEventListener('scroll', () => {
            body.style.height = window.innerHeight + 'px';
        });
    }

    window.addEventListener('load', setupScroll);
    window.addEventListener('resize', () => {
        let pageWidth = body.offsetWidth;
        if (lastWidth === pageWidth) {
            return;
        }
        lastWidth = pageWidth;
        setupScroll();
    });

    function setupScroll() {
        if (ios) {
            let height = document.documentElement.clientHeight;
            if (iphone && !fullscreen) {
                height += 60;
            }
            body.style.height = height + 'px';
        } else if (android) {
            body.style.height = (window.innerHeight + 56) + 'px';
        }
        setTimeout(scrollTo, 0, 0, 1);
    }

})();
