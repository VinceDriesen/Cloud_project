window.onload = function() {
    const elements = document.querySelectorAll('#home > h1, #home > h2, #home > a');
    elements.forEach((element, index) => {
        setTimeout(
            () => {
                // element.classList.add('animate-ease-in-down');
                element.classList.remove('invisible');
                element.classList.add('animate-ease-in-down')
            },
            (index + 1) * 150
        );
    });
};
