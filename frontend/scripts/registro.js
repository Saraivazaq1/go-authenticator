const form = document.querySelector('form');

form.addEventListener('submit', async (event) => {
    event.preventDefault();

    // Pega os valores do input para criar o body do response
    const username = form.elements['username'].value;
    const email = form.elements['email'].value;
    const password = form.elements['password'].value;

    try {

        // Construção do response com o body
        const response = await fetch('/registro', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, email, password })
        });

        const data = await response.json();

        // Tratamento de erros
        if (!response.ok) {
            alert(data.error || 'Erro ao registrar');
            return;
        }

        // Salva o token no localStorage
        localStorage.setItem('token', data.token);

        // Redireciona manualmente
        window.location.href = '/userPage';

    } catch (error) {

        // Tratamento de erros
        console.error('Erro:', error);
        alert('Erro na requisição. Tente novamente.');
    }
});
