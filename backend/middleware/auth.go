package middleware

import (
	"goauthenticator/backend/database"
	"goauthenticator/backend/environment"
	"goauthenticator/backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Verifica se o header é nulo
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
			ctx.Abort()
			return
		}

		// Verifica o conteúdo do header 
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			ctx.Abort()
			return
		}
		
		// Identifica se o token foi assinado com a chave configurada
		tokenString := tokenParts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenInvalidSubject
			}
			return environment.TokenKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		// Verificação dos claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["username"] == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token malformado"})
			ctx.Abort()
			return
		}

		// Busca o usuário baseado nos claims
		var user models.User
		if err := database.DB.First(&user, claims["username"]).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	}
}
