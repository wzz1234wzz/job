package main

import "errors"

func Add(a int,b int)int{
	return a+b
}

func Subtract(a int,b int)int{
	return a-b
}

func Multiply(a int,b int)int{
	return a*b
}

func Divide(a int,b int)int{
	if b==0{
		errors.New("除数不能为0！")
	}
	return a/b
}