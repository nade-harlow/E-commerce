package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/adapter/api/controllers/product"
)

func DefineRoutes(r *gin.Engine, pc *product.ProductController) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{"message": "Hello World!"})
	})
	r.POST("/products", pc.GetProduct())
}
