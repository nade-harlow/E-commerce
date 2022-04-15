package product

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/utils/response"
	"github.com/nade-harlow/E-commerce/ports/services"
	"log"
)

type ProductController struct {
	ProductService services.ProductServices
	route          *gin.Engine
}

func NewProductController(productService services.ProductServices) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (products *ProductController) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		product := models.Product{}
		c.ShouldBindJSON(&product)
		err := products.ProductService.CreateProduct(&product)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error creating Product", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product created successfully", nil, nil)
	}
}
