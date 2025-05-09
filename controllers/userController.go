package controllers

import (
	"goauthenticator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	userInterface, exists := ctx.Get("user")

	// Verificação de erros
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Usuário não encontrado no contexto"})
		return
	}

	user, ok := userInterface.(models.User)

	// Verificação de erros
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao interpretar os dados do usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
	})
}