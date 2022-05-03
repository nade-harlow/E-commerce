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

func (products ProductController) GetAllProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		product, err := products.ProductService.GetAllProducts()
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error getting all Product", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product retrieved successfully", product, nil)
	}
}

func (products *ProductController) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := products.ProductService.DeleteProduct(id)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error deleting Product", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product deleted successfully", nil, nil)
	}
}

func (products *ProductController) AddProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := models.ProductCategory{}
		c.ShouldBindJSON(&category)
		err := products.ProductService.CreateProductCategory(&category)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error creating Product Category", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product Category created successfully", nil, nil)
	}
}

func (products *ProductController) RemoveProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("id")
		err := products.ProductService.DeleteProductCategory(categoryID)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error deleting Product Category", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product Category deleted successfully", nil, nil)
	}
}
