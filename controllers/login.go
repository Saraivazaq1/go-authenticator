package controllers

import (
	"goauthenticator/database"
	"goauthenticator/environment"
	"goauthenticator/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {

	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Não possível entrar"})
	}

	var user models.User

	if err := database.DB.First(&user, "username = ?", input.Username).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credenciais inválidas"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credenciais inválidas"})
	}

	// Criando o token de autenticação
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.ID,
			"exp":      time.Now().Add(time.Duration(environment.GetTokenExpirationMinutes())).Unix(),
		})

	tokenString, err := token.SignedString(environment.TokenKey)

	// Tratamento de erros
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login bem-sucedido",
		"token": tokenString,
	})

}
