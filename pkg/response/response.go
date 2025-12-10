package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, ResponseData{Code: code, Msg: msg[code], Data: data})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, ResponseData{Code: code, Msg: msg[code], Data: nil})
}
