package xorm_sqlmock

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Account struct {
	ID      int64  `xorm:"pk autoincr"` // ID 指定主键并自增
	Name    string `xorm:"unique"`      //唯一的
	Balance float64
	Version int `xorm:"version"` //乐观锁
}
// 1.增
func newAccount(db *xorm.Engine,name string, balance float64) error {
	_, err := db.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 2.按ID查
func getAccount(db *xorm.Engine,id int64) (*Account, error) {
	a := &Account{}
	has, err := db.Id(id).Get(a) // 查询操作 var a Account; x.Id(1).Get(&a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("账户不存在")
	}
	return a, nil
}
