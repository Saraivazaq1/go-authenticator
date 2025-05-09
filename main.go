package main

import (
	"goauthenticator/database"
	"goauthenticator/models"
	"goauthenticator/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Se conecta ao banco de dados
	database.ConectarDB()
	database.DB.AutoMigrate(&models.User{}) // Cria a tabela users

	gin.SetMode(gin.ReleaseMode)

	// Instancia o servidor
	router := gin.Default()
	routes.ConfigurarRotas(router)
	router.Run()
}