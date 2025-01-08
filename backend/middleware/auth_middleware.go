package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kh3rld/e-waste-soko/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Authorization header missing or invalid"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token payload"})
			c.Abort()
			return
		}

		userType, ok := claims["user_type"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token payload"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(userID))
		c.Set("user_type", userType)

		c.Next()
	}
}
