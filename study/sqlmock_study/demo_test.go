package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func (a *api) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestShouldGetPosts(t *testing.T) {
	db, mock, err := sqlmock.New() //初始化sqlmock
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &api{db} // 使用mock的虚拟数据库进行模拟数据库测试
	req, err := http.NewRequest("GET", "http://localhost/posts", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// 在我们实际执行api函数之前，我们需要预期所需的DB操作,设置期望逻辑
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")
	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows) // 定义执行的数据库操作、返回row的头和行

	app.posts(w, req) //现在我们执行我们的请求
	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}
	data := struct {
		Posts []*post
	}{Posts: []*post{
		{ID: 1, Title: "post 1", Body: "hello"},
		{ID: 2, Title: "post 2", Body: "world"},
	}}
	app.assertJSON(w.Body.Bytes(), data, t) // /测试w.Body的byte是否可以解码成功
	if err := mock.ExpectationsWereMet(); err != nil { // 我们确保所有的期望都得到满足
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldRespondWithErrorOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	app := &api{db}
	req, err := http.NewRequest("GET", "http://localhost/posts", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	//mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectQuery("SELECT id, title, body FROM posts").WillReturnError(fmt.Errorf("some error"))

	// now we execute our request
	app.posts(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	data := struct {
		Error string
	}{"failed to fetch posts: some error"}
	app.assertJSON(w.Body.Bytes(), data, t)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
