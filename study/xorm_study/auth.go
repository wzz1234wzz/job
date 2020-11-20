package main

type Client interface {GetType(flag int)string}

type Service struct {}

func (c *Service)GetType(flag int)string{
	if flag>0{return "add"}
	if flag<0{return "sub"}
	return "zero"
}

func Handle(a,b,flag int,client Client)int{
	switch client.GetType(flag) {
	case "add":return a + b
	case "sub":return a-b
	default:return 0
	}
	return 0
}


// Stub案例1
//func GetType(flag int)string{
//	if flag>0{return "add"}
//	if flag<0{return "sub"}
//	return "zero"
//}
//
//var CallFunc func(int)(string)=GetType
//
//func Handle(a,b,flag int)int{
//	switch CallFunc(flag) {
//	case "add":return a + b
//	case "sub":return a-b
//	default:return 0
//	}
//	return 0
//}


// Mock案例1
////假设我们有一个依赖http请求的鉴权接口
//type AuthService interface{
//	Login(username string,password string) (token string,e error)
//	Logout(token string) error
//}