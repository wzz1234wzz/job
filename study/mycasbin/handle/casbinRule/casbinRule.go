package casbinRule

import (
	"github.com/gin-gonic/gin"
	"log"
	"mycasbin/dao"
	"net/http"
)

type reqAdd struct {
	V0    string `json "v0" example:"admin"`
	V1    string `json "v1" example:"/api/hello"`
	V2    string `json "v2" example:"GET"`
}

type respAdd struct {
	Message string  `json:"message" binding:"required"`
}

// @Summary 新增策略
// @Description  新增策略
// @Security ApiKeyAuth
// @Tags Policy
// @Accept json
// @Produce json
// @Param param body casbinRule.reqAdd true "新增策略"
// @Success 200 {object}  casbinRule.respAdd "正常返回表示创建成功，否则返回错误码和错误信息"
// @Router /policy/add [post]
func HandleAddPolicy(c *gin.Context){
	log.Println("增加Policy")
	var req reqAdd
	var resp respAdd
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("err=",err)
		return
	}
	if ok:= dao.Casbin.Enforcer.AddPolicy(req.V0, req.V1, req.V2); !ok {
		resp.Message = "Policy已经存在"
		c.JSON(http.StatusBadRequest, resp)
	} else {
		resp.Message = "Policy增加成功"
		c.JSON(http.StatusOK, resp)
	}
	log.Println(resp.Message)
}

type reqDelete struct {
	V0    string `json "v0" example:"admin"`
	V1    string `json "v1" example:"/api/hello"`
	V2    string `json "v2" example:"GET"`
}

type respDelete struct {
	Message string  `json:"message" binding:"required"`
}

// @Summary 删除policy
// @Description  删除policy
// @Security ApiKeyAuth
// @Tags Policy
// @Accept json
// @Produce json
// @Param param body casbinRule.reqDelete true "删除策略"
// @Success 200 {object}  casbinRule.respDelete "正常返回表示删除成功，否则返回错误码和错误信息"
// @Router /policy/delete [post]
func HandleDeletePolicy(c *gin.Context){
	log.Println("删除Policy")
	var req reqDelete
	var resp respDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if ok:= dao.Casbin.Enforcer.RemovePolicy(req.V0, req.V1, req.V2); !ok {
		resp.Message = "Policy不存在"
		c.JSON(http.StatusNotFound, resp)
	} else {
		resp.Message = "Policy删除成功"
		c.JSON(http.StatusOK, resp)
	}
	log.Println(resp.Message)
}

//获取policy
type reqGet struct {

}

type respGet struct {
	Policys [][]string `json:"policy" binding:"required"`
}

// @Summary 获取policy
// @Description  获取policy
// @Security ApiKeyAuth
// @Tags Policy
// @Accept json
// @Produce json
// @Success 200 {object}  casbinRule.respGet "正常返回表示获取成功，否则返回错误码和错误信息"
// @Router /policy/get [get]
func HandleGetPolicy(c *gin.Context) {
	log.Println("查看policy")
	var resp respGet
	list := dao.Casbin.Enforcer.GetPolicy()
	for _, vlist := range list {
		var temp []string
		for _, v := range vlist {
			temp = append(temp,v)
		}
		resp.Policys = append(resp.Policys,temp)
	}
	c.JSON(http.StatusOK, resp)
	log.Println("查看policy成功！")
}

type reqRoleAdd struct {
	User    string `json "user" binding:"required" example:"admin"`
	Role    string `json "role" binding:"required" example:"root"`
}

type respRoleAdd struct {
	Message string  `json:"message" binding:"required"`
}

// @Summary 新增角色
// @Description  新增角色
// @Security ApiKeyAuth
// @Tags Role
// @Accept json
// @Produce json
// @Param param body casbinRule.reqRoleAdd true "新增角色"
// @Success 200 {object}  casbinRule.respRoleAdd "正常返回表示创建成功，否则返回错误码和错误信息"
// @Router /role/add [post]
func HandleAddRole(c *gin.Context){
	log.Println("增加Role")
	var req reqRoleAdd
	var resp respRoleAdd
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("err=",err)
		return
	}
	if ok:= dao.Casbin.Enforcer.AddRoleForUser(req.User, req.Role); !ok {
		resp.Message = "Role已经存在"
		c.JSON(http.StatusBadRequest, resp)
	} else {
		resp.Message = "Role增加成功"
		c.JSON(http.StatusOK, resp)
	}
	log.Println(resp.Message)
}

type reqRoleDelete struct {
	User    string `json "user" binding:"required" example:"root"`
	Role    string `json "role" binding:"required" example:"admin"`
}

type respRoleDelete struct {
	Message string  `json:"message" binding:"required" example:"sucess!"`
}

// @Summary 删除Role
// @Description  删除Role
// @Security ApiKeyAuth
// @Tags Role
// @Accept json
// @Produce json
// @Param param body casbinRule.reqRoleDelete true "删除策略"
// @Success 200 {object}  casbinRule.respRoleDelete "正常返回表示删除成功，否则返回错误码和错误信息"
// @Router /role/delete [post]
func HandleDeleteRole(c *gin.Context){
	log.Println("删除Role")
	var req reqRoleDelete
	var resp respRoleDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if ok:= dao.Casbin.Enforcer.DeleteRoleForUser(req.User, req.Role); !ok {
		resp.Message = "Role不存在"
		c.JSON(http.StatusNotFound, resp)
	} else {
		resp.Message = "Role删除成功"
		c.JSON(http.StatusOK, resp)
	}
	log.Println(resp.Message)
}

//获取Role
type reqRoleGet struct {

}

type respRoleGet struct {
	Roles []string `json:"policy" binding:"required"`
}

// @Summary 获取Role
// @Description  获取Role
// @Security ApiKeyAuth
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {object}  casbinRule.respRoleGet "正常返回表示获取成功，否则返回错误码和错误信息"
// @Router /role/get [get]
func HandleGetRole(c *gin.Context) {
	log.Println("查看Role")
	var resp respRoleGet
	list := dao.Casbin.Enforcer.GetAllRoles()
	for _, vlist := range list {
		resp.Roles = append(resp.Roles,vlist)
	}
	c.JSON(http.StatusOK, resp)
	log.Println("查看Role成功！")
}
