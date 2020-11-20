package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "io/ioutil"
)

func getApi (c *gin.Context){
    fmt.Println(c.Query("id"))
    c.String(http.StatusOK,"ok!恭喜你访问成功！")
}

func postApi (c *gin.Context){
    fmt.Println(c.PostForm("id"))
    c.String(http.StatusOK,"ok")
}

func postjson (c *gin.Context) {
    var data = &struct {
        Name string `json:"title"`
    }{}
    
    c.BindJSON(data)
    
    fmt.Println(data)
    c.String(http.StatusOK,"ok")
    
}

func sayHello(w http.ResponseWriter, r *http.Request){
    b,_:=ioutil.ReadFile("./hello.txt")
    _,_=fmt.Fprintln(w,string(b))
}

func sayHello1 (c *gin.Context) {
    c.JSON(200,gin.H{"message":"hallo golang"}) // 200表示状态码
}

func studyCookie (c *gin.Context) {
    cookie, err := c.Cookie("gin_cookie") // 获取Cookie
    if err != nil {
        cookie = "NotSet"
        // 设置Cookie
        c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
    }else {
        fmt.Printf("Cookie value: %s \n", cookie)
    }
    c.JSON(200,gin.H{"message":cookie}) // 200表示状态码
}

func main() {
    r := gin.Default() // 返回默认引擎
    r.GET("/getApi",getApi)//http://127.0.0.1:8080/getApi
    r.POST("/postApi",postApi)//注册接口
    r.POST("/postjson",postjson)//注册接口
    r.GET("/hello",sayHello1)
    r.GET("/cookie", studyCookie)
    r.Run()
}