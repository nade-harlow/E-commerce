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

func (cart *CartController) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		item, err := cart.CartService.GetCart()
		if err != nil {
			response.Json(c, 500, "error fetching cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "cart fetched", item, nil)
	}
}

func (cart *CartController) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("itemID")
		err := cart.CartService.AddItem(itemID)
		if err != nil {
			response.Json(c, 500, "error adding item to cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "item added to cart", nil, nil)
	}
}

func (cart *CartController) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("itemID")
		err := cart.CartService.RemoveItem(itemID)
		if err != nil {
			response.Json(c, 500, "error removing item from cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "item removed from cart", nil, nil)
	}
}
