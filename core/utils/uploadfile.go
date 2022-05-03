package utils

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func AwsSession(fileExtension, folder string) (*session.Session, string, error) {
	tempFileName := folder + "/" + fmt.Sprint(rand.Int()) + fileExtension

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_SECRET_ID"),
			os.Getenv("AWS_SECRET_KEY"),
			os.Getenv("AWS_TOKEN"),
		),
	})

	return session, tempFileName, err
}

func UploadToS3(sess *session.Session, file multipart.File, fileName string, size int64) (string, error) {
	bucket := os.Getenv("S3_BUCKET_NAME")

	defer file.Close()

	buffer := make([]byte, size)
	file.Read(buffer)

	uploader := s3manager.NewUploader(sess)
	url := "https://s3-eu-west-3.amazonaws.com/e-pine/" + fileName
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		log.Printf("Unable to upload  %q to %q %v", fileName, bucket, err)
		return "", err
	}

	log.Println("link:", url)
	log.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	return url, nil
}

func CheckSupportedFile(filename string) (string, bool) {
	supportedFileTypes := map[string]bool{
		".png":  true,
		".jpeg": true,
		".jpg":  true,
	}
	fileExtension := filepath.Ext(filename)

	return fileExtension, !supportedFileTypes[fileExtension]
}
