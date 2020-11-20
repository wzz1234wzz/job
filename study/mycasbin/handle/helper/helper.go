package helper

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"mycasbin/errcode"
	"fmt"
)

// Resp defines a generic type to be used as an http response
type Resp struct {
	Code    errcode.Code `json:"code" binding:"required"`
	Message string       `json:"message" binding:"required"`
	Data    interface{}  `json:"data" binding:"required"`
}

// Ok
func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Resp{
		Code: errcode.ErrNone,
		Data: data,
	})
}

// ReplyErr
func ReplyErr(c *gin.Context, code errcode.Code, e ...interface{}) {
	log.Print("err_code: ",code," err= ", e)
	c.JSON(http.StatusInternalServerError, Resp{
		Code:    code,
		Message: fmt.Sprint(e...),
		Data:    ""})
}

// ReplyErr
func ReplyErrf(c *gin.Context, code errcode.Code, format string, e ...interface{}) {
	log.Print("err_code: ",code," err= ", e)
	c.JSON(http.StatusInternalServerError, Resp{
		Code:    code,
		Message: fmt.Sprintf(format, e...),
		Data:    ""})
}

// ReplyErr
func ReplyAuthErr(c *gin.Context, code errcode.Code, e ...interface{}) {
	log.Print("err_code: ",code," err= ", e)
	c.JSON(http.StatusUnauthorized, Resp{
		Code:    code,
		Message: fmt.Sprint(e...),
		Data:    ""})
}

// ReplyErr
func ReplyAuthErrf(c *gin.Context, code errcode.Code, format string, e ...interface{}) {
	log.Print("err_code: ",code," err= ", e)
	c.JSON(http.StatusUnauthorized, Resp{
		Code:    code,
		Message: fmt.Sprintf(format, e...),
		Data:    ""})
}

