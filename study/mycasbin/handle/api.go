package handle

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "mycasbin/docs"
	"mycasbin/handle/casbinRule"
	"mycasbin/handle/hello"
	"mycasbin/handle/job"
	"mycasbin/handle/middleware"
	"mycasbin/handle/user"
)

// @title Casbin Example API
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com
// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email wzz@yuansuan.cn
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /api
func InitHandlers() {
	//获取router路由对象
	r := gin.New()
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//User API
	userApi := r.Group("/api/user")
	{
		userApi.POST("/auth", user.HandleAuth)
		userApi.POST("/adduser", middleware.JWT(), middleware.Authorize(), user.HandleAddUser)
		userApi.GET("/current", middleware.JWT(), middleware.Authorize(), user.HandleCurrentUser)
	}

	//Policy API
	policyApi := r.Group("/api/policy")
	policyApi.Use(middleware.JWT(),middleware.Authorize())
	{
		policyApi.POST("/add", casbinRule.HandleAddPolicy)
		policyApi.POST("/delete", casbinRule.HandleDeletePolicy)
		policyApi.GET("/get", casbinRule.HandleGetPolicy)
	}

	//Role API
	roleApi := r.Group("/api/role")
	roleApi.Use(middleware.JWT(),middleware.Authorize())
	{
		roleApi.POST("/add", casbinRule.HandleAddRole)
		roleApi.POST("/delete", casbinRule.HandleDeleteRole)
		roleApi.GET("/get", casbinRule.HandleGetRole)
	}

	//创建请求
	r.GET("/api/hello", middleware.JWT(), middleware.Authorize(), hello.Hello)

	jobApi := r.Group("/api/job")
	jobApi.Use(middleware.JWT(), middleware.Authorize())
	{
		jobApi.GET("/add", job.HandleAddJob)
		jobApi.GET("/delete", job.HandleDeleteJob)
	}

	r.Run(":9090") //参数为空 默认监听9090端口
}
