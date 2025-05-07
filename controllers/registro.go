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

func Registrar(ctx *gin.Context) {

	// Struct de input de valores
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// Tratamento de erros
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Informações inválidas"})
		return
	}

	var count int64
	database.DB.Model(&models.User{}).Where("email = ? OR username = ?", input.Email, input.Username).Count(&count)

	if count > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Usuário já cadastrado"})
		return
	}
	// Criando o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	// Tratamento de erros
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	// Criando um usuário no banco de dados
	user := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}

	// Tratamento de erros
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
		return
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
		"message": "Usuário registrado com sucesso",
		"token":   tokenString,
	})
}
