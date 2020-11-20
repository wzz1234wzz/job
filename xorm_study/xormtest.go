package main

import (
	"errors"
	"fmt"
	"log"
     "os"
	"os/user" //获取当前用户信息
	"os/signal" //信号处理
	"os/exec" //执行外部命令
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Account 定义结构体(xorm支持双向映射)
type Account struct {
	ID      int64  `xorm:"pk autoincr"` // ID 指定主键并自增
	Name    string `xorm:"unique"`      //唯一的
	Balance float64
	Version int `xorm:"version"` //乐观锁
}

// PersonTable 人员结构表
type PersonTable struct {
	ID         int64  `xorm:"pk autoincr"`
	PersonName string `xorm:"varchar(24)"`
	PersonAge  int    `xorm:"int default 0"`
	PersonSex  int    `xorm:"notnull"`
	Salary     int    `xorm:"int"`
	//City            CityTable `xorm:"-"`
}

// CityTable 城市表
type CityTable struct {
	CityName      string
	CityLongitude float32
	CityLatitude  float32
}

//定义orm引擎
var x *xorm.Engine

// 0.创建orm引擎
func init() {
	var err error                                                                                    // 不可以x, err ：= xorm.NewEngine
	x, err = xorm.NewEngine("mysql", "lambdacal:yskj2407@tcp(10.0.1.58:3306)/zbb_test?charset=utf8") // engine 可以通过 engine.Close 来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭。
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(Account)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	log.Println("连接数据库成功...")
}

// 1.增
func newAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 2.按ID查
func getAccount(id int64) (*Account, error) {
	a := &Account{}
	has, err := x.Id(id).Get(a) // 查询操作 var a Account; x.Id(1).Get(&a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("账户不存在")
	}
	return a, nil
}

// 3.存
func makeDeposit(id int64, deposit float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	a.Balance += deposit
	_, err = x.Id(id).Update(a)
	return a, nil
}

// 4.取
func makeWithdraw(id int64, withdraw float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	if a.Balance < withdraw {
		return nil, errors.New("账户余额太少，不够转")
	}
	a.Balance -= withdraw
	_, err = x.Id(id).Update(a)
	return a, nil
}

// 5.转账
func makeTransfer(id1, id2 int64, balance float64) error {
	a1, err := getAccount(id1)
	if err != nil {
		return err
	}
	a2, err := getAccount(id2)
	if err != nil {
		return err
	}

	if a1.Balance < balance {
		return errors.New("账户余额太少，不够转")
	}

	a1.Balance -= balance
	a2.Balance += balance

	if _, err = x.Id(id1).Cols("balance").Update(a1); err != nil {
		return err
	} else if _, err = x.Id(id2).Cols("balance").Update(a2); err != nil {
		return err
	}
	return nil
}

// 6. getAccountsAsxID 获取所有账户
func getAccountsAsxID() (as []*Account, err error) {
	err = x.Asc("id").Find(&as)
	return as, err
}

// 7. 练习
func test() {
	fmt.Println("测试7入口...")

	info, err := x.DBMetas() // 展示所有的表
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range info {
		log.Println(v.Name)
	}

	isExit, _ := x.IsTableExist("person_table")
	if isExit {
		log.Println("0:PersonTable已经存在")
		err := x.DropTables(new(PersonTable))
		if err != nil {
			log.Println(err)
		}
		isExit, _ := x.IsTableExist("person_table")
		if !isExit {
			log.Println("0:PersonTable已经删除")
		}
	} else {
		log.Println("0:PersonTable不存在")
	}

	if err := x.Sync(new(PersonTable)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	// 1. 创建表及插入数据
	x.Insert(&PersonTable{ID: 1, PersonName: "wzz", PersonAge: 10, PersonSex: 1})
	x.Insert(&PersonTable{ID: 2, PersonAge: 12, PersonSex: 0})

	// 2. 查询数据
	// a. 根据Id来获得单条数据:
	person1 := &PersonTable{}
	has, _ := x.Id(1).Get(person1) // has=true or false
	fmt.Printf("a:%#v %#v\n", has, person1)
	fmt.Println("----------------------------")

	// b. 根据where获取单条数据
	var person PersonTable
	x.Where("person_age = ? or person_sex = ?", 12, 0).Get(&person)
	fmt.Printf("b:%#v\n", &person)
	fmt.Println("----------------------------")

	// c. 根据Account结构体中存在的非空数据来获取单条数据
	a := &PersonTable{ID: 1}
	x.Get(a)
	fmt.Printf("c:%#v\n", a)
	fmt.Println("----------------------------")

	// 3.更新
	person.Salary = 100
	x.Update(&person)
	fmt.Printf("3:%#v\n", &person)
	fmt.Println("----------------------------")

	SQL := "insert into person_table (PersonAge,PersonSex) values($1,$2)"
	h, _ := x.Exec(SQL, 222, 1)
	log.Println(h)

}

// 8. 删除
func deleteAccount(id int64) error {
	//_,err := x.Delete(&Account{Id:id})
	_, err := x.Id(id).Delete(&Account{})
	return err
}

func os_test(){
	hostname,_ := os.Hostname() // Hostname 函数会返回内核提供的主机名。
	fmt.Printf("hostname=%s\n",hostname)
	fmt.Println("----------------------------")

	environ:= os.Environ() //Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝。
	fmt.Printf("environ=%v\n",environ)
	fmt.Println("----------------------------")

	getenv:=os.Getenv("PATH") // Getenv 函数会检索并返回名为 key 的环境变量的值。如果不存在该环境变量则会返回空字符串。
	fmt.Printf("getenv=%v\n",getenv)
	fmt.Println("----------------------------")

	userId:=os.Getuid() // Getuid 函数可以返回调用者的用户 ID。
	fmt.Printf("userId=%v\n",userId)
	fmt.Println("----------------------------")

	groupId:=os.Getgid() // Getpid Getuid 函数可以返回调用者的用户 ID。
	fmt.Printf("groupId=%v\n",groupId)
	fmt.Println("----------------------------")

	pId:=os.Getpid() // Getpid Getuid 函数可以返回调用者的用户 ID。
	fmt.Printf("pId=%v\n",pId)
	fmt.Println("----------------------------")

	pwd,_:= os.Getwd() // Getwd 函数可以返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd 会返回其中一个
	fmt.Printf("pwd=%v\n",pwd)
	fmt.Println("----------------------------")

	err:=os.Mkdir("./testdir", 0777) //Mkdir 函数可以使用指定的权限和名称创建一个目录。如果出错，会返回 *PathError 底层类型的错误
	if err==nil{
		fmt.Println("testdir创建成功")
	}else{
		fmt.Printf("创建失败,err=%v\n",err)
	}
	fmt.Println("----------------------------")

	err_rm:=os.Remove("./test.txt") // Remove 函数会删除 name 指定的文件或目录。如果出错，会返回 *PathError 底层类型的错误。RemoveAll 函数跟 Remove 用法一样，区别是会递归的删除所有子目录和文件。
	if err_rm==nil{
		fmt.Println("testdir删除成功")
	}else{
		fmt.Printf("删除失败,err=%v\n",err_rm)
	}
	fmt.Println("----------------------------")

	f, err := exec.LookPath("./") // 在环境变量 PATH 指定的目录中搜索可执行文件，如果 file 中有斜杠，则只在当前目录搜索。返回完整路径或者相对于当前目录的一个相对路径。
	if err != nil {
		fmt.Printf("失败,err=%v\n",err)
	}else {
		fmt.Printf("可执行文件f=%v\n", f)
	}
	fmt.Println("----------------------------")

	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户id", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组id：", u.Gid)
	s, _ := u.GroupIds() // 用户所在的所有的组的id
	log.Println("用户所在的所有组：", s)
	fmt.Println("----------------------------")

	c := make(chan os.Signal, 0)
	signal.Notify(c) // 第一个参数表示接收信号的 channel，第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
	signal.Stop(c)   //不允许继续往c中存入内容
	t := <-c // Block until a signal is received.//c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
	fmt.Println("Got signal:", t)
	fmt.Println("--------运行该程序，然后在 CMD 窗口中通过 Ctrl+C 来结束该程序，便会得到输出结果-------------")

	os.Exit(0) // Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
}
