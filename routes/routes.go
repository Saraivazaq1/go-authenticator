package routes

import "github.com/gin-gonic/gin"

func ConfigurarRotas(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Servidor rodando"})
	})
}