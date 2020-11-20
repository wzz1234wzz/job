package xorm_sqlmock

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-xorm/xorm"
)

func TestNewAccount(t *testing.T)  {
	var x xorm.Engine
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	x.db=db

	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
}

