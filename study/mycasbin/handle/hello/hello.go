package hello

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags Hello
// @Security ApiKeyAuth
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "fail"}"
// @Router /hello [get]
func Hello(c *gin.Context) {
	log.Println("Hello 接收到GET请求..")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": "Hello 接收到GET请求..",
	})
}
