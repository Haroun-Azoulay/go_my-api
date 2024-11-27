package security

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	tokenString := c.GetHeader("Authorization") 
	

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]



	err := VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func IsAdmin(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to parse claims"})
		c.Abort()
		return
	}

	isAdmin, ok := claims["IsAdmin"].(bool)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "IsAdmin claim not found or invalid type"})
		c.Abort()
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have admin privileges"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Admin access granted"})
	c.Next()
}
