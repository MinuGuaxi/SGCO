const ConfigNavItem = document.querySelectorAll('.config-nav-item');

ConfigNavItem.forEach(link => {
    link.addEventListener('click', (event) =>   {
        ConfigNavItem.forEach(link => link.classList.remove('active'));

        event.target.classList.add('active');
       // event.preventDefault();

        const ConfigItem = event.target.dataset.configItem;
        console.log('O item:', ConfigItem);
    })
})