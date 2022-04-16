package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils"
	"github.com/nade-harlow/E-commerce/core/utils/response"
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

func (user *UserController) SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest models.User
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			response.Json(c, 500, "Error binding json", nil, err.Error())
			return
		}
		if err := utils.ValidateStruct(userRequest); err != nil {
			response.Json(c, 400, "Error validating data", nil, err.Error())
			return
		}
		if err := user.UserService.SignUpUser(&userRequest); err != nil {
			response.Json(c, 500, "Error creating user", nil, err.Error())
			return
		}
		token, err := utils.GenerateToken(&userRequest)
		if err != nil {
			response.Json(c, 500, "Error generating token", nil, err.Error())
			return
		}
		data := map[string]interface{}{
			"token":           token,
			"user_id":         userRequest.ID,
			"role":            userRequest.Role,
			"user_first_name": userRequest.FirstName,
		}
		response.Json(c, 200, "User created", data, nil)
	}
}

func (user UserController) SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest requests.UserLoginRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			response.Json(c, 500, "Error binding json", nil, err.Error())
			return
		}
		if err := utils.ValidateStruct(userRequest); err != nil {
			response.Json(c, 400, "Error validating data", nil, err.Error())
			return
		}
		userdata, err := user.UserService.SignInUser(&userRequest)
		if err != nil {
			response.Json(c, 500, "Error signing in user", nil, err.Error())
			return
		}
		token, err := utils.GenerateToken(userdata)
		if err != nil {
			response.Json(c, 500, "Error generating token", nil, err.Error())
			return
		}
		data := map[string]interface{}{
			"token":           token,
			"user_id":         userdata.ID,
			"role":            userdata.Role,
			"user_first_name": userdata.FirstName,
		}
		response.Json(c, 200, "User signed in", data, nil)
	}
}
