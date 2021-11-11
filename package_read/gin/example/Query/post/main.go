package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.POST("/post",func(c *gin.Context){
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name:= c.PostForm("name")
		message:= c.PostForm("message")
		fmt.Println("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	r.Run(":8080")
}
/*curl -F "name=manu&message=this_is_great" http://127.0.0.1:8080/post -v
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /post HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.64.1
> Accept: *
> Content-Length: 165
> Content-Type: multipart/form-data; boundary=------------------------bb3de0b2208f8ac2
>
* We are completely uploaded and fine
< HTTP/1.1 200 OK
< Date: Tue, 12 Oct 2021 15:42:21 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
* Closing connection 0
*/