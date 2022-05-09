package product

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils/response"
	"github.com/nade-harlow/E-commerce/ports/services"
	"log"
	"strconv"
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
		price, _ := strconv.Atoi(c.PostForm("price"))
		qty, _ := strconv.Atoi(c.PostForm("quantity"))

		form, err := c.MultipartForm()
		if err != nil {
			log.Printf("error parsing multipart form: %v", err)
			response.Json(c, 500, "error parsing multipart form", nil, err.Error())
			return
		}
		productImages := form.File["image"]
		images, err := products.ProductService.UploadFileToS3(productImages)
		if err != nil {
			response.Json(c, 500, "error Uploading file to S3", nil, err.Error())
			return
		}
		product := models.Product{
			Name:              c.PostForm("name"),
			Description:       c.PostForm("description"),
			Sku:               c.PostForm("sku"),
			ProductImage:      images,
			ProductCategoryID: c.PostForm("category_id"),
			ProductCategory:   models.ProductCategory{},
			Price:             float32(price),
			Quantity:          int16(qty),
		}
		err = products.ProductService.CreateProduct(&product)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error creating Product", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product created successfully", nil, nil)
	}
}

func (products ProductController) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")
		product, err := products.ProductService.GetProduct(productID)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error getting Product", nil, err.Error())
		}
		response.Json(c, 200, "Product retrieved successfully", product, nil)
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

func (products *ProductController) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")
		price, _ := strconv.Atoi(c.PostForm("price"))
		qty, _ := strconv.Atoi(c.PostForm("quantity"))

		form, err := c.MultipartForm()
		if err != nil {
			log.Printf("error parsing multipart form: %v", err)
			response.Json(c, 500, "error parsing multipart form", nil, err.Error())
			return
		}
		productImages := form.File["image"]
		images, err := products.ProductService.UploadFileToS3(productImages)
		if err != nil {
			response.Json(c, 500, "error Uploading file to S3", nil, err.Error())
			return
		}
		product := map[string]interface{}{
			"ID":                productID,
			"Name":              c.PostForm("name"),
			"Description":       c.PostForm("description"),
			"Sku":               c.PostForm("sku"),
			"ProductImage":      images,
			"ProductCategoryID": c.PostForm("category_id"),
			"Price":             float32(price),
			"Quantity":          int16(qty),
		}
		err = products.ProductService.UpdateProduct(productID, product)
		if err != nil {
			log.Println(err.Error())
			response.Json(c, 500, "Error updating Product", nil, err.Error())
			return
		}
		response.Json(c, 200, "Product updated successfully", nil, nil)
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
		category := requests.ProductCategoryRequest{}
		c.ShouldBindJSON(&category)
		err := products.ProductService.CreateProductCategory(category)
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
