package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/utils"
	"net/http"
)

func AuthorizeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := utils.ParseToken(authHeader)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(string)

		c.Set("userId", userId)

		c.Next()
	}

}

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := utils.ParseToken(authHeader)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		if role != "admin" {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Next()
	}
}
