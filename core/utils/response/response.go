package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//GeneralResponse returns a standard response format
type GeneralResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
}

func (response *GeneralResponse) GetFailedResponse(err interface{}, message string) GeneralResponse {
	response.Error = err
	response.Message = message
	response.Success = false
	return *response
}

func (response *GeneralResponse) GetFailedResponseWithData(data interface{}, err interface{}, message string) GeneralResponse {
	response.Data = data
	response.Error = err
	response.Message = message
	response.Success = false
	return *response
}

func (response *GeneralResponse) GetSuccessfulResponse(data interface{}, message string) GeneralResponse {
	response.Data = data
	response.Message = message
	response.Success = true
	return *response
}

func Json(c *gin.Context, status int, message string, data interface{}, errs interface{}) {
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}
	c.JSON(status, responsedata)
}
