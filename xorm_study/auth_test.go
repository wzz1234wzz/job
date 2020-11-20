package main

import "testing"

// Mock与Stub相结合：mock一个struct实现interface里面的函数
// 该函数采用调用函数的方式来处理
type MockService struct{
	StubFunc func(flag int)string
	StubFuncUsed bool
}

func (s *MockService)GetType(flag int)string{
	s.StubFuncUsed=true
	return s.StubFunc(flag)
}

func TestHandleMock(t *testing.T) {
       client:=&MockService{}
       client.StubFunc=func(flag int)string{return "add"}
		res:=Handle(1,3,0,client)
		t.Log(res,1,3)
		res =Handle(2,5,1,client)
		t.Log(res,2,5)
		client.StubFunc=func(flag int)string{return "zero"}
		res=Handle(5,1,-2,client)
		t.Log(res,5,1)
}

//func TestHandle(t *testing.T) {
//	CallFunc=StubFunc
//	res:=Handle(1,3,0)
//	t.Log(res,1,3)
//	res =Handle(2,5,1)
//	t.Log(res,2,5)
//	res=Handle(5,1,-2)
//	t.Log(res,5,1)
//}

//func StubFunc(a int)string{return "add"}

//func TestHandle(t *testing.T) {
//	client:=&Service{}
//	res:=Handle(1,3,0,client)
//	t.Log(res,1,3)
//	res =Handle(2,5,1,client)
//	t.Log(res,2,5)
//	res=Handle(5,1,-2,client)
//	t.Log(res,5,1)
//}

//import "errors"
//
//type authService struct {}
//
//func (auth *authService) Login (username string,password string) (string,error){
//	return "token", nil
//}
//
//func (auth *authService) Logout(token string) error{
//	return nil
//}
//
////模拟登录失败
//type authLoginErr struct {
//	auth AuthService  //可以使用组合的特性，Logout方法我们不关心，只用“覆盖”Login方法即可
//}
//func (auth *authLoginErr) Login (username string,password string) (string,error) {
//	return "", errors.New("用户名密码错误")
//}
//
////模拟api服务器宕机
//type authUnavailableErr struct {
//}
//func (auth *authUnavailableErr) Login (username string,password string) (string,error) {
//	return "", errors.New("api服务不可用")
//}
//func (auth *authUnavailableErr) Logout(token string) error{
//	return errors.New("api服务不可用")
//}