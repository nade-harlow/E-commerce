package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils"
	"github.com/nade-harlow/E-commerce/core/utils/response"
	"github.com/nade-harlow/E-commerce/ports/services"
	"log"
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
			response.Json(c, 400, "Error validating data", nil, err)
			return
		}
		if err := user.UserService.SignUpUser(&userRequest); err != nil {
			response.Json(c, 500, "Error creating user", nil, err.Error())
			return
		}
		response.Json(c, 200, "User created", nil, nil)
	}
}

func (user *UserController) SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest requests.UserLoginRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			response.Json(c, 500, "Error binding json", nil, err.Error())
			return
		}
		if err := utils.ValidateStruct(userRequest); err != nil {
			response.Json(c, 400, "Error validating data", nil, err)
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

func (user *UserController) VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.PostForm("code")
		err := user.UserService.VerifyUser(code)
		if err != nil {
			response.Json(c, 500, "Error verifying user", nil, err.Error())
			return
		}
		response.Json(c, 200, "User verified", nil, nil)
	}
}

func (user UserController) AddUserAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest models.UserAddress
		userID := c.Param("id")
		userRequest.UserID = userID
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			response.Json(c, 500, "Error binding json", nil, err.Error())
			return
		}
		if err := utils.ValidateStruct(userRequest); err != nil {
			response.Json(c, 400, "Error validating data", nil, err)
			return
		}
		if err := user.UserService.AddUserAddress(&userRequest); err != nil {
			response.Json(c, 500, "Error adding user address", nil, err.Error())
			return
		}
		response.Json(c, 200, "User address added", nil, nil)
	}
}

func (user UserController) UpdateUserAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest models.UserAddress
		userID := c.Param("id")
		userRequest.UserID = userID
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			response.Json(c, 500, "Error binding json", nil, err.Error())
			return
		}
		if err := utils.ValidateStruct(userRequest); err != nil {
			response.Json(c, 400, "Error validating data", nil, err)
			return
		}
		if err := user.UserService.UpdateUserAddress(&userRequest); err != nil {
			response.Json(c, 500, "Error updating user address", nil, err.Error())
			return
		}
		response.Json(c, 200, "User address updated", nil, nil)
	}
}

func (user UserController) ForgotPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		ID := c.Param("id")
		if err := user.UserService.ForgotPassword(ID, email); err != nil {
			response.Json(c, 500, "Error forgot password", nil, err.Error())
			return
		}
		response.Json(c, 200, "Password reset email sent", nil, nil)
	}
}

func (user UserController) RecoverPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")
		c.HTML(200, "change_password.html", ID)
	}
}

func (user UserController) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.PostForm("password")
		confirm := c.PostForm("confirm")
		userID := c.Param("id")
		log.Println(userID)
		if password != confirm {
			response.Json(c, 400, "Passwords do not match", nil, "Passwords do not match")
		}
		if err := user.UserService.ResetUserPassword(userID, password); err != nil {
			response.Json(c, 500, "Error resetting password", nil, err.Error())
			return
		}
		response.Json(c, 200, "Password reset", nil, nil)
	}
}
