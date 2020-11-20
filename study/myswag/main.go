package main

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "myswag/docs"
	"fmt"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /wzz

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	hi := r.Group("/wzz/hi")
	{
		hi.GET("/hello", HandleHello)
	}

	download := r.Group("/wzz/runtime")
	{
		download.POST("/download", DownAccountImage)
	}

	nihao := r.Group("/wzz/load")
	{
		nihao.POST("/login", HandleLogin)
	}
	r.Run(":9090")
}

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hi/hello [get]
func HandleHello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}

// DownloadTxtFile godoc
// @Summary Download txt file
// @Description Download file
// @Tags download
// @Accept json
// @Produce application/octet-stream
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /runtime/download [post]
func DownAccountImage(ctx *gin.Context) {
	//ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	//ctx.File("./a.txt")
	srcFile := "./testDir/testDirA/testFile.txt"
	destZip := "/mnt/e/wzz.zip"
	err := Zip(srcFile, destZip)
	if err !=nil{
		fmt.Println("err:",err)
	}
	ctx.Writer.Header().Add("Content-Type", "application/zip")
	ctx.File(destZip)
}

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

func DirUnzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type reqAuth struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"passwd" binding:"required" example:"admin"`
}

type respAuth struct {
	Isok string `json:"isok" binding:"required" example:"yes"`
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"passwd" binding:"required" example:"admin"`
}

// @Summary 模拟登录
// @Description  登录鉴权，输入用户名和密码，返回Token
// @Tags User
// @Accept json
// @Produce json
// @Param param body reqAuth true "用户名和密码"
// @Success 200 {object} respAuth  "验证是否通过"
// @Router /load/login [post]
func HandleLogin(c *gin.Context) {
	var req reqAuth
	var resp respAuth
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
		return
	}
	if req.Password == req.Username{
		resp.Isok = "yes"
		resp.Username = req.Username
		resp.Password = req.Username
	}else{
		resp.Isok = "no"
	}
	c.JSON(http.StatusOK, resp)
}