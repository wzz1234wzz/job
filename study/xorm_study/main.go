package main

import (
	"fmt"
)

const prompt = "请输入操作的数字: 1.创建一个账户 2.获取账户的细节 3.存 4.取 5.转账 6. 查看所有账户 9.退出"

func main() {

	fmt.Println("欢迎来到xorm!")
	// Exit:
	// 	for{
	fmt.Println(prompt)

	var num int
	//fmt.Scanf("%d\n",&num)
	num = 10

	switch num {
	case 1: //创建用户
		fmt.Println("请输入<name><balence>")
		var name string
		var balance float64
		fmt.Scanf("%s %f\n", &name, &balance)
		if err := newAccount(name, balance); err != nil {
			fmt.Println(err)
		}
		fmt.Println("成功创建！")
	case 2: // 查看某一用户id
		fmt.Println("请输入用户的<id>")
		var id int64
		fmt.Scanf("%d\n", &id)
		a, err := getAccount(id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%#v\n", a)
		}
	case 3: // 向用户账户打入钱
		fmt.Println("请输入<id><deposit>")
		var id int64
		var deposit float64
		fmt.Scanf("%d %f\n", &id, &deposit)
		a, err := makeDeposit(id, deposit)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%#v\n", a)
		}
		fmt.Println("存钱成功！")
	case 4: // 向用户账户取出钱
		fmt.Println("请输入<id><withdraw>")
		var id int64
		var withdraw float64
		fmt.Scanf("%d %f\n", &id, &withdraw)
		a, err := makeWithdraw(id, withdraw)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%#v\n", a)
		}
		fmt.Println("取出成功！")
	case 5: // 向用户账户转账
		fmt.Println("请输入<id1><withdraw><id2>")
		var id1, id2 int64
		var balance float64
		fmt.Scanf("%d %f %d\n", &id1, &balance, &id2)
		err := makeTransfer(id1, id2, balance)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("转账成功！")
		}
	case 6: //查询所有账户
		as, err := getAccountsAsxID()
		if err != nil {
			fmt.Println(err)
		} else {
			for i, a := range as {
				fmt.Printf("%d: %#v\n", i, a)
			}
		}
	case 7:
		test()
		fmt.Println("测试7")
	case 8:
		fmt.Println("请输入<id>")
		var id int64
		fmt.Scanf("%d\n", &id)
		if err := deleteAccount(id); err != nil {
			fmt.Println(err)
		}
	case 9:
		//break Exit
		fmt.Println("测试9")
	case 10:
		os_test()
		fmt.Println("测试10结束")
	}
	//}

}
