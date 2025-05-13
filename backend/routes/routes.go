package routes

import (
	"goauthenticator/backend/controllers"
	"goauthenticator/backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(r *gin.Engine) {

	// Rota padrão
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{"status": http.StatusOK})
	})

	r.GET("/:page", func(ctx *gin.Context) {
		page := ctx.Param("page")
		ctx.HTML(http.StatusOK, page+".html", nil)
	})

	// Rota de registro
	r.POST("/registro", func(ctx *gin.Context) {
		controllers.Registrar(ctx)
	})

	// Rota de login
	r.POST("/login", func(ctx *gin.Context) {
		controllers.Login(ctx)
	})

	r.GET("/user", middleware.AuthMiddleware(), controllers.GetUser) // Rota de usuário
}
