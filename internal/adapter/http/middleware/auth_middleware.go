package middleware

import (
	"jamlink-backend/internal/shared/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(securitySvc security.SecurityService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing or invalid"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := securitySvc.ValidateJWT(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		isVerified, ok := claims["isVerified"].(bool)

		if !ok || !isVerified {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "your account is not verified"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["id"])

		c.Next()
	}
}
