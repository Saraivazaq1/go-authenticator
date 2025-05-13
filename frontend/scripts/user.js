document.addEventListener("DOMContentLoaded", async () => {
	const token = localStorage.getItem("token"); // Pega o token de autenticação no localStorage

	// Tratamento de erros
	if (!token) {
		alert("Você precisa estar logado.");
		window.location.href = "/loginPage";
		return;
	}

	try {

		// Se o token estiver correto o usuário será autenticado e irá para a página de usuário
		const response = await fetch("/user", {
			method: "GET",
			headers: {
				"Authorization": `Bearer ${token}`,
			},
		});

		// Tratamento de erros
		if (!response.ok) {
			alert("Sessão expirada ou token inválido.");
			localStorage.removeItem("token");
			window.location.href = "/loginPage";
			return;
		}

		const data = await response.json();

		// Preenche os campos com os dados do usuário
		document.getElementById("user-id").textContent = `ID: ${data.id}`;
		document.getElementById("user-username").textContent = `Username: ${data.username}`;
		document.getElementById("user-email").textContent = `Email: ${data.email}`;

	} catch (error) {

		// Tratamento de erros
		console.error("Erro ao buscar dados do usuário:", error);
		alert("Erro ao carregar dados do usuário.");
		localStorage.removeItem("token");
		window.location.href = "/loginPage";
	}
});

// Função de logout
document.getElementById("logout-button").addEventListener("click", () => {
	localStorage.removeItem("token");
	window.location.href = "/";
});
