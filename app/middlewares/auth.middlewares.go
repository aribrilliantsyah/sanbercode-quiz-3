package middlewares

import (
	"book-management/util/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "authorization header is required",
			})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "invalid authorization header format",
			})
			c.Abort()
			return
		}

		userID, err := jwt.VerifyToken(bearerToken[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
