# Sistema de autenticação em Go

Um sistema feito para o registro e login de usuários utilizando um sistema de autenticação via tokens JWT.

---
## Importante
Após clonar o projeto é necessário renomear o arquivo `.env.example` para `.env`.

```bash
cp .env.example .env
```
Após isso é necessário preencher os valores das variáveis de ambiente com as informações dos banco de dados e do token de autenticação.

## Execução
1- Atualize dependências
```go
go mod tidy
```

2- Execute o projeto
```go
gin run main.go
```

3- Abra o arquivo (por padrão está na porta 3000, a home page está na rota padrão "/")
```bash
localhost:3000/
```
