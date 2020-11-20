package job

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type resp struct {
	Message string  `json:"message" binding:"required"`
}

// @Summary 增
// @Description 测试增
// @Security ApiKeyAuth
// @Tags Job
// @Success 200 {object}  job.resp "正常返回表示创建成功，否则返回错误码和错误信息"
// @Failure 400 {string} job.resp  "{"msg": "fail"}"
// @Router /job/add [get]
func HandleAddJob(c *gin.Context) {
	var resp resp
	resp.Message = "接收到add job请求.."
	log.Println("请求add job")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})
}

// @Summary 删
// @Description 测试删
// @Security ApiKeyAuth
// @Tags Job
// @Success 200 {object}  job.resp "正常返回表示删成功，否则返回错误码和错误信息"
// @Failure 400 {string} job.resp  "{"msg": "fail"}"
// @Router /job/delete [get]
func HandleDeleteJob(c *gin.Context) {
	var resp resp
	resp.Message = "接收到delete job请求.."
	log.Println("请求delete job")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})
}