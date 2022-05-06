package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/utils/response"
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

func (cart CartController) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("id")
		item, err := cart.CartService.GetCart(itemID)
		if err != nil {
			response.Json(c, 500, "error fetching cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "cart fetched", item, nil)
	}
}
