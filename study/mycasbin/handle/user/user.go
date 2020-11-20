package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"mycasbin/consts"
	"mycasbin/dao"
	"mycasbin/db"
	"mycasbin/errcode"
	"mycasbin/handle/helper"
	"mycasbin/handle/middleware"
)

type reqAuth struct {
	Username string `json:"username" binding:"required" example:"root"`
	Password string `json:"passwd" binding:"required" example:"phenglei!@#"`
}

type respAuth struct {
	Token    string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJleHAiOjE1OTI3MDY3MTIsImlzcyI6InBoZW5nbGVpIn0.lKUH_H9lAMPgRrjpK4w67qYox1grQ7qWSh_0ZrDwu68"`
	JwtToken string `json:"jwttoken" binding:"required" example:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJleHAiOjE1OTI3MDY3MTIsImlzcyI6InBoZW5nbGVpIn0.lKUH_H9lAMPgRrjpK4w67qYox1grQ7qWSh_0ZrDwu68"`
}

// @Summary 登录鉴权
// @Description  登录鉴权，输入用户名和密码，返回Token
// @Tags User
// @Accept json
// @Produce json
// @Param param body user.reqAuth true "用户名和密码"
// @Success 200 {object} user.respAuth  "验证通过后返回Token"
// @Router /user/auth [post]
func HandleAuth(c *gin.Context) {
	var req reqAuth
	var resp respAuth
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Fatal("解析数据错误，err=",err)
		return
	}
	session := db.Engine.Context(c)
	defer session.Close()

	user := dao.User.QueryByUserName(c, session, req.Username)
	if user == nil {
		helper.ReplyErr(c, errcode.ErrAuthUserIsNotExist, err)
		return
	}
	err = helper.PwdVerify(req.Password, user.PwdHash)
	if err != nil {
		helper.ReplyErr(c, errcode.ErrAuthPwdNotMatch, err)
		return
	}

	resp.Token, err = middleware.GenerateToken(user.Ysid)
	resp.JwtToken = "Bearer " + resp.Token
	if err != nil {
		helper.ReplyErr(c, errcode.ErrInvalidParam, err)
		return
	}
	helper.Ok(c, resp)
	return
}

type reqAddUser struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"passwd" binding:"required" example:"admin"`
}

type respAddUser struct {
	Token string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJleHAiOjE1OTI3MDY3MTIsImlzcyI6InBoZW5nbGVpIn0.lKUH_H9lAMPgRrjpK4w67qYox1grQ7qWSh_0ZrDwu68"`
	ID    int64  `json:"id" binding:"required" example:"1"`
}

// @Summary 添加用户
// @Description  添加用户，输入用户的用户名和密码即可
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param param body user.reqAddUser true "用户名和密码"
// @Success 200 {object} user.respAddUser  "正常返回"
// @Router /user/adduser [post]
func HandleAddUser(c *gin.Context) {
	var req reqAddUser
	var resp respAddUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ReplyErr(c, errcode.ErrInvalidParam, err)
		return
	}

	session := db.Engine.Context(c)
	defer session.Close()

	pwdHash, _ := helper.PwdGenerate(req.Password)
	userID, err := dao.User.AddUser(c, session, req.Username, pwdHash)
	if err != nil {
		helper.ReplyErr(c, errcode.ErrInvalidParam, "用户名重复，添加失败！ "+err.Error())
		return
	}

	resp.Token, err = middleware.GenerateToken(userID)
	if err != nil {
		helper.ReplyErr(c, errcode.ErrInvalidParam, err)
		return
	}
	resp.ID = userID
	helper.Ok(c, resp)
	return
}

type reqCurrentUser struct {
}

type respCurrentUser struct {
	UserID   int64  `json:"userid" binding:"required" example:"2"`
	UserName string `json:"username" binding:"required" example:"admin"`
}

// @Summary 查询当前用户
// @Description  通过Token查询当前用户信息，通过token,返回用户ID和用户名
// @Tags User
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} user.respCurrentUser  "返回用户ID和用户名"
// @Router /user/current [get]
func HandleCurrentUser(c *gin.Context) {
	userID := c.GetInt64(consts.HeaderUserID)
	session := db.Engine.Context(c)
	defer session.Close()
	user := dao.User.QueryByUserID(c, session, userID)
	if user == nil {
		helper.ReplyAuthErr(c, errcode.ErrAuthUserIsNotExist, errors.New("User is not exist.\n"))
	}
	var resp respCurrentUser
	resp.UserID = user.Ysid
	resp.UserName = user.Name
	helper.Ok(c, resp)
}
