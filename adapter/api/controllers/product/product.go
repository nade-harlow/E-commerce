package product

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/ports"
	"log"
)

type ProductController struct {
	ProductService ports.Service
	route          *gin.Engine
}

func NewProductController(productService ports.Service) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (products *ProductController) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		product := models.Product{}
		c.ShouldBindJSON(&product)
		err := products.ProductService.Create(&product)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
