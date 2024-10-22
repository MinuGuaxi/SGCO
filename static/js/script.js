document.addEventListener('DOMContentLoaded', function () {
    const sendButton = document.getElementById('sendButton');
    const messageInput = document.getElementById('messageInput');
    const conversationLinks = document.querySelectorAll('.content-conversas-header');
    const defaultMessage = document.querySelector('.conversation-default');
    const conversationAreas = document.querySelectorAll('.conversation');

    // Função para trocar de conversa
    conversationLinks.forEach(link => {
        link.addEventListener('click', function (e) {
            e.preventDefault();

            // Esconder mensagem padrão
            if (defaultMessage) {
                defaultMessage.style.display = 'none';
            }

            // Identificar a conversa a ser mostrada
            const conversationId = link.getAttribute('data-conversation');
            const conversationToShow = document.querySelector(conversationId);

            // Esconder todas as outras conversas
            conversationAreas.forEach(area => {
                area.style.display = 'none';  // Esconder a conversa
            });

            // Mostrar a conversa selecionada
            if (conversationToShow) {
                conversationToShow.style.display = 'block';  // Mostrar a conversa selecionada
            } else {
                console.error(`Conversa com ID ${conversationId} não encontrada.`);
            }
        });
    });

    // Função para adicionar mensagens à conversa ativa
    function addMessage(content, isSender = true) {
        const activeConversation = document.querySelector('.conversation[style*="display: block"] .conversation-wrapper');

        if (!activeConversation) {
            console.error("Nenhuma conversa ativa.");
            return;
        }

        const messageWrapper = document.createElement('li');
        messageWrapper.classList.add('conversation-item');
        if (isSender) messageWrapper.classList.add('me');

        const messageContent = `
            <div class="conversation-item-side">
                <img src="sender-avatar-url.jpg" alt="User Avatar">
            </div>
            <div class="conversation-item-content">
                <div class="conversation-item-box">
                    <div class="conversation-item-text"><p>${content}</p></div>
                    <div class="conversation-item-time">Agora</div>
                </div>
            </div>
        `;

        messageWrapper.innerHTML = messageContent;
        activeConversation.appendChild(messageWrapper);
        scrollToBottom(activeConversation);
    }

    // Função para rolar até o final da conversa
    function scrollToBottom(conversationWrapper) {
        conversationWrapper.scrollTop = conversationWrapper.scrollHeight;
    }

    // Evento de clique no botão de enviar mensagem
    sendButton.addEventListener('click', function () {
        const message = messageInput.value.trim();
        if (message) {
            addMessage(message);
            messageInput.value = ''; // Limpar campo de mensagem
        }
    });

    // Enviar mensagem com a tecla "Enter"
    messageInput.addEventListener('keydown', function (e) {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            sendButton.click();  // Dispara o clique do botão de enviar
        }
    });
});
