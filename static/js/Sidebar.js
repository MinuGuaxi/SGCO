document.addEventListener('DOMContentLoaded', () => {
    const configNavItems = document.querySelectorAll('.config-nav-item');
    const sections = document.querySelectorAll('.conteudo-section');

    // Adiciona o evento de clique para alternar as seções
    configNavItems.forEach(link => {
        link.addEventListener('click', (event) => {
            event.preventDefault();

            // Remove 'active' de todos os itens e seções
            configNavItems.forEach(item => item.classList.remove('active'));
            sections.forEach(section => section.classList.remove('active'));

            // Adiciona 'active' ao item clicado e à seção correspondente
            event.target.classList.add('active');
            const targetSectionId = event.target.getAttribute('data-config-nav-item');
            document.getElementById(targetSectionId).classList.add('active');
        });
    });
});
