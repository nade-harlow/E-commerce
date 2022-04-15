package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/ports/services"
)

type UserController struct {
	UserService services.UserServices
	route       *gin.Engine
}

func NewUserController(productService services.UserServices) *UserController {
	return &UserController{
		UserService: productService,
	}
}
