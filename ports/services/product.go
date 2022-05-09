package services

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
	"log"
	"mime/multipart"
)

const FileSize = 5 * 1024 * 1024

type ProductServices interface {
	CreateProduct(product *models.Product) error
	GetProduct(productID string) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(productID string, product map[string]interface{}) error
	DeleteProduct(productID string) error
	CreateProductCategory(category requests.ProductCategoryRequest) error
	DeleteProductCategory(categoryID string) error
	UploadFileToS3(productImages []*multipart.FileHeader) ([]models.ProductImage, error)
	CreateS3Bucket(session *session.Session, bucketName string) error
}

type ProductService struct {
	repository repository2.ProductRepository
}

func NewProductService(repository repository2.ProductRepository) ProductServices {
	return &ProductService{
		repository: repository,
	}
}

func (p *ProductService) CreateProduct(product *models.Product) error {
	return p.repository.CreateProduct(product)
}

func (p *ProductService) GetProduct(productID string) (*models.Product, error) {
	return p.repository.GetProduct(productID)
}

func (p ProductService) GetAllProducts() ([]models.Product, error) {
	return p.repository.GetAllProducts()
}

func (p ProductService) UpdateProduct(productID string, product map[string]interface{}) error {
	return p.repository.UpdateProduct(productID, product)
}

func (p ProductService) DeleteProduct(productID string) error {
	return p.repository.DeleteProduct(productID)
}

func (p *ProductService) CreateProductCategory(category requests.ProductCategoryRequest) error {
	cate := &models.ProductCategory{
		Name:        category.Name,
		Description: category.Description,
	}
	return p.repository.CreateProductCategory(cate)
}

func (p ProductService) DeleteProductCategory(categoryID string) error {
	return p.repository.DeleteProductCategory(categoryID)
}

func (p ProductService) UploadFileToS3(productImages []*multipart.FileHeader) ([]models.ProductImage, error) {
	var images []models.ProductImage
	for _, f := range productImages {
		file, err := f.Open()
		if err != nil {
			return nil, err
		}
		if f.Size > FileSize {
			return nil, errors.New("file size is too large")
		}

		fileExtension, ok := utils.CheckSupportedFile(f.Filename)
		if ok {
			return nil, errors.New("image file type is not supported")
		}

		session, tempFileName, err := utils.AwsSession(fileExtension, "products")
		if err != nil {
			log.Println("could not upload file", err)
			return nil, err
		}

		url, err := utils.UploadToS3(session, file, tempFileName, f.Size)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		image := models.ProductImage{Url: url}
		images = append(images, image)
	}
	return images, nil
}

func (p *ProductService) CreateS3Bucket(session *session.Session, bucketName string) error {
	// Create S3 service client
	svc := s3.New(session)

	// Create bucket
	result, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	fmt.Println("Bucket created:", aws.StringValue(result.Location))
	return nil
}
