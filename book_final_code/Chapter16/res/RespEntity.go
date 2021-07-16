package res

import (
	"food/myerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
func SendResponse(c *gin.Context,err error,data interface{}){
	code,message := myerr.DecodeErr(err)
	c.JSON(http.StatusOK,Response{
		Code:code,
		Message: message,
		Data:data,
	})
}