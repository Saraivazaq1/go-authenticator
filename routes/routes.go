package routes

import (
	"goauthenticator/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Servidor rodando"})
	})

	r.POST("/registro", func(ctx *gin.Context)  {
		controllers.Registrar(ctx)
	})

	r.POST("/login", func(ctx *gin.Context)  {
		controllers.Login(ctx)
	})
}