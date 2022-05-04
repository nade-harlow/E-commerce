package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/adapter/api/controllers/product"
	"github.com/nade-harlow/E-commerce/adapter/api/controllers/user"
	"github.com/nade-harlow/E-commerce/adapter/api/middleware"
)

func ProductRoutes(r *gin.Engine, pc *product.ProductController) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{"message": "Hello World!"})
	})
	r.GET("/products", pc.GetAllProduct())
	r.POST("/products/add", middleware.AuthorizeToken(), pc.AddProduct())
	r.POST("/products/category/add", middleware.AuthorizeToken(), pc.AddProductCategory())
	r.PUT("/products/update/:id", middleware.AuthorizeToken(), pc.UpdateProduct())
	r.DELETE("/products/category/remove/:id", middleware.AuthorizeToken(), pc.RemoveProductCategory())
	r.DELETE("/products/delete/:id", middleware.AuthorizeToken(), pc.DeleteProduct())
}

func UserRoutes(r *gin.Engine, uc *user.UserController) {
	r.POST("/signup", uc.SignUpUser())
	r.POST("/login", uc.SignInUser())
	r.POST("/verify/otp", uc.VerifyUser())
}
