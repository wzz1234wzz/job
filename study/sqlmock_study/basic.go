package main

import (
	"database/sql"
	"fmt"
)

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	var rows *sql.Rows
	rows, err = tx.Query("SELECT id, title, body FROM user")
	if err != nil {
		return
	}

	var id int
	var title,body string
	for rows.Next() {
		if err = rows.Scan(&id, &title, &body); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("id=%v,title=%v,body=%v\n",id,title,body)
	}
	if rows.Err() != nil {
		fmt.Println(err)
		return
	}
	rows.Close()

	var result sql.Result
	result, err = tx.Exec("UPDATE products SET views = views + 1")
	if err != nil {
		fmt.Println("更新错误，err=",err)
		return
	}else{
		id,_:=result.LastInsertId()
		fmt.Printf("更新id=%v\n",id)
	}

	result, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID)

	if err != nil {
		fmt.Println("插入错误，err=",err)
		return
	}else{
		id,_:=result.LastInsertId()
		fmt.Printf("插入id=%v\n",id)
	}
	return
}

//func main() {
//	// @NOTE: the real connection is not required for tests
//	db, err := sql.Open("mysql", "root@/blog")
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	if err = recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
//		panic(err)
//	}
//}
