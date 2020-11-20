package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"mycasbin/consts"
	"mycasbin/errcode"
	"mycasbin/handle/helper"
	"strings"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		log.Print("beartoken = ", bearToken)
		if len(bearToken) == 0 {
			helper.ReplyAuthErr(c, errcode.ErrAuthHeaderNoToken, "Header has not found auth token")
			c.Abort()
			return
		}
		token := strings.ReplaceAll(bearToken, "Bearer ", "")
		if len(token) == 0 {
			helper.ReplyAuthErr(c, errcode.ErrAuthHeaderNoToken, "Header has not found auth token")
			c.Abort()
			return
		}

		cl, err := ParseToken(token)
		if err != nil {
			helper.ReplyAuthErr(c, errcode.ErrAuthInvalidToken, "Invalid auth token")
			c.Abort()
			return
		}
		log.Print("current UserID :", cl.UserID)
		c.Set(consts.HeaderUserID, cl.UserID)
		c.Next()
	}
}
