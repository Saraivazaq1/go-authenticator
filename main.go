package main

import (
	"goauthenticator/backend/database"
	"goauthenticator/backend/models"
	"goauthenticator/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Se conecta ao banco de dados
	database.ConectarDB()
	database.DB.AutoMigrate(&models.User{}) // Cria a tabela users

	gin.SetMode(gin.ReleaseMode)

	// Instancia o servidor
	router := gin.Default()

	// Renderizando o HTML e conectando os arquivos estáticos
	router.LoadHTMLGlob("frontend/templates/*")
	router.Static("/styles", "./frontend/styles")
	router.Static("/scripts", "./frontend/scripts")

	routes.ConfigurarRotas(router)
	router.Run()
}