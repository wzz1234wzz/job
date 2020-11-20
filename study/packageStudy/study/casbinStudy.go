package study


import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/casbin/xorm-adapter/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func baseCase(){
	fmt.Println("-----------1---------------------")
	e, err := casbin.NewEnforcer("./static/model1.conf", "./static/policy1.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
	check(e, "root", "data3", "rwx")

	fmt.Println("-----------2---------------------")
	e, err = casbin.NewEnforcer("./static/model2.conf", "./static/policy2.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	check(e, "dajun", "data", "read")
	check(e, "dajun", "data", "write")
	check(e, "lizi", "data", "read")
	check(e, "lizi", "data", "write")
	check(e, "dajun", "prod.data", "read")
	check(e, "dajun", "prod.data", "write")
	check(e, "lizi", "dev.data", "read")
	check(e, "lizi", "dev.data", "write")
	check(e, "lizi", "prod.data", "write")

	fmt.Println("-----------3---------------------")
	e, err = casbin.NewEnforcer("./static/model3.conf", "./static/policy3.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	check(e, "dajun", "data", "read")
	check(e, "dajun", "data", "write")
	check(e, "lizi", "data", "read")
	check(e, "lizi", "data", "write")
}

func CasbinMain() {
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := xormadapter.NewAdapter("mysql", "lambdacal:yskj2407@tcp(10.0.1.58:3306)/") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e, _ := casbin.NewEnforcer("./static/model1.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	ok, _ := e.Enforce("alice", "data1", "read")
	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	roles, err:= e.GetRolesForUser("dajun")
	fmt.Println("s=",roles)
	fmt.Println("err=",err)

	// Save the policy back to DB.
	e.SavePolicy()
}