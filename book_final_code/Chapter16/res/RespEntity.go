package res

import (
<<<<<<< HEAD
	"food/myerr"
	"github.com/gin-gonic/gin"
=======
	"github.com/gin-gonic/gin"
	"food/myerr"
>>>>>>> 3bc95edaf0eee6696f3b7d4b974b987237244d4f
	"net/http"
)

type Response struct {
<<<<<<< HEAD
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := myerr.DecodeErr(err)

	//http.StatusOK这个值是200
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
=======
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`

}
func SendResponse(c *gin.Context,err error,data interface{}){
	code,message := myerr.DecodeErr(err)
	c.JSON(http.StatusOK,Response{
		Code:code,
		Message: message,
		Data:data,
>>>>>>> 3bc95edaf0eee6696f3b7d4b974b987237244d4f
	})
}