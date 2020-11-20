package db

import (
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"mycasbin/dao"
	"mycasbin/dao/models"
)

var Engine *xorm.Engine

func Init() {
	CasbinSetup()
	InitDB()
}

// 初始化casbin
func CasbinSetup(){
	a:= xormadapter.NewAdapter(dao.DBNAME, dao.DNS,true)
	e := casbin.NewEnforcer("./conf/rbac_models.conf", a)
	dao.Casbin.Enforcer = e
}

func InitDB() {
	var err error                                                                                    // 不可以x, err ：= xorm.NewEngine
	Engine, err = xorm.NewEngine(dao.DBNAME, dao.DNS) // engine 可以通过 engine.Close 来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭。

	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := Engine.Sync(new(models.User)); err != nil {
		log.Fatal("用户数据表同步失败:", err)
	}
	if err := Engine.Sync(new(models.CasbinRule)); err != nil {
		log.Fatal("Casbin数据表同步失败:", err)
	}
	log.Println("连接数据库成功...")
}