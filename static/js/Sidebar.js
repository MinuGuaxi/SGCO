const ConfigNavItem = document.querySelectorAll('.config-nav-item');
const sections = document.querySelectorAll('section[class^="conteudo-1"]');

// Função para alternar as seções de acordo com o item de navegação clicado
ConfigNavItem.forEach(link => {
    link.addEventListener('click', (event) => {
        // Remove a classe 'active' de todos os itens de navegação
        ConfigNavItem.forEach(link => link.classList.remove('active'));

        // Adiciona a classe 'active' ao item de navegação clicado
        event.target.classList.add('active');

        // Esconde todas as seções
        sections.forEach(section => {
            section.classList.remove('active');
        });

        // Mostra a seção correspondente ao item de navegação clicado
        const targetSection = event.target.dataset.configNavItem;
        document.getElementById(targetSection).classList.add('active');
    });
});
