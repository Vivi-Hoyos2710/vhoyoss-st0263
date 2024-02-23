package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseJSON struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Token interface{} `json:"token,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, ResponseJSON{
		Code: http.StatusText(status),
		Data: data})
}
func SuccessLogin(c *gin.Context, status int, data interface{}) {
	Response(c, status, ResponseJSON{Token: data})
}

func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := ErrorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}
	Response(c, status, err)
}
