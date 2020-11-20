package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"mycasbin/consts"
	"mycasbin/dao"
	"mycasbin/db"
)

// Authorize 拦截器
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := dao.Casbin.Enforcer //var e *casbin.Enforcer
		e.LoadPolicy() //从DB加载策略
		obj := c.Request.URL.RequestURI() //获取请求的URI
		act := c.Request.Method //获取请求方法
		userID := c.GetInt64(consts.HeaderUserID)
		session := db.Engine.Context(c)
		defer session.Close()
		user := dao.User.QueryByUserID(c, session, userID)
		sub := user.Name //获取用户的角色 应该从db中读取
		//判断策略中是否存在
		if ok := e.Enforce(sub, obj, act); ok {
			log.Println("恭喜您,权限验证通过")
			c.Next() // 进行下一步操作
		} else {
			log.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}
