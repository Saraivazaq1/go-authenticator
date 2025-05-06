package main

import (
	"goauthenticator/database"
	"goauthenticator/models"
	"goauthenticator/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConectarDB()
	database.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	routes.ConfigurarRotas(router)
	router.Run()
}