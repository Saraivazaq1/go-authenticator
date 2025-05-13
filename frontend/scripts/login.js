const form = document.getElementById('login-form');

form.addEventListener('submit', async (event) => {
	event.preventDefault();

	// Construção do body com o input da página
	const username = document.getElementById('username').value;
	const password = document.getElementById('password').value;

	try {

		const response = await fetch('/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ username, password })
		});

		const data = await response.json();

		if (response.ok) {
			// Salva o token no localStorage
			localStorage.setItem('token', data.token);

			// Redireciona para a página protegida
			window.location.href = '/userPage';
		} else {
			alert(data.error || 'Erro ao fazer login');
		}

	} catch (error) {

		// Tratamento de erros
		console.error(error);
		alert('Erro na requisição. Tente novamente.');
	}
});