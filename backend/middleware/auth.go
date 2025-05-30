package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId uint   `json:"userId"`
	Rol    string `json:"rol"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	key := []byte(os.Getenv("JWT_SECRET"))
	return func(c *gin.Context) {
		if len(key) == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET no configurado"})
			return
		}
		//Leer el header Authorization
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			// Si no tiene el prefijo "Bearer ", retornar error
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			return
		}

		// Soportar formato: "Bearer <token>"
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		// Guardar claims en el contexto
		c.Set("userId", claims.UserId)
		c.Set("rol", claims.Rol)
		c.Next() // Continuar con la siguiente función del middleware
	}
}
