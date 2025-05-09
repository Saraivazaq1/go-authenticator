package routes

import (
	"goauthenticator/controllers"
	"goauthenticator/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(r *gin.Engine) {

	// Rota padrão
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Servidor rodando"})
	})

	// Rota de registro
	r.POST("/registro", func(ctx *gin.Context)  {
		controllers.Registrar(ctx)
	})

	// Rota de login
	r.POST("/login", func(ctx *gin.Context)  {
		controllers.Login(ctx)
	})

	r.GET("/user", middleware.AuthMiddleware(), controllers.GetUser) // Rota de usuário
}