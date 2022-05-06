package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/utils/response"
	"github.com/nade-harlow/E-commerce/ports/services"
	"strconv"
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

func (cart *CartController) GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		item, err := cart.CartService.GetCart()
		if err != nil {
			response.Json(c, 500, "error fetching cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "cart fetched", item, nil)
	}
}

func (cart *CartController) AddCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("productID")
		err := cart.CartService.AddItem(productID)
		if err != nil {
			response.Json(c, 500, "error adding item to cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "item added to cart", nil, nil)
	}
}

func (cart *CartController) DeleteCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("id")
		err := cart.CartService.RemoveItem(itemID)
		if err != nil {
			response.Json(c, 500, "error removing item from cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "item removed from cart", nil, nil)
	}
}

func (cart CartController) UpdateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("id")
		qty := c.Param("quantity")
		quantity, err := strconv.Atoi(qty)
		if err != nil {
			response.Json(c, 500, "error updating item in cart", nil, err.Error())
		}
		err = cart.CartService.UpdateItem(itemID, int16(quantity))
		if err != nil {
			response.Json(c, 500, "error updating item in cart", nil, err.Error())
			return
		}
		response.Json(c, 200, "item updated in cart", nil, nil)
	}
}
