package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/utils"
	"net/http"
)

func AuthurizeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := utils.ParseToken(authHeader)
		if err != nil {
			//utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userId := int64(claims["user_id"].(float64))

		c.Set("userId", userId)

		c.Next()
	}

}
