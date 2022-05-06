package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/ports/services"
)

type CartController struct {
	CartService services.CartServices
	route       *gin.Engine
}

func NewCartController(productService services.CartServices) *CartController {
	return &CartController{
		CartService: productService,
	}
}
