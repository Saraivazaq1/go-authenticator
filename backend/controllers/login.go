package controllers

import (
	"goauthenticator/backend/database"
	"goauthenticator/backend/environment"
	"goauthenticator/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {

	// Entrada de dados
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Verificação de erros
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Não possível entrar"})
		return
	}

	// Instancia da struct
	var user models.User

	// Verificação de erros
	if err := database.DB.First(&user, "username = ?", input.Username).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Comparação da senha e do hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Criando o token de autenticação
	expiration := time.Now().Add(time.Duration(environment.GetTokenExpirationMinutes()) * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.ID,
			"exp":      expiration.Unix(),
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
