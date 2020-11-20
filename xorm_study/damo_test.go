package main

import (
	"fmt"
	"os"
	"testing"
)

//type calcCase struct{ A, B, Expected int }
//
//func createMulTestCase(t *testing.T, c *calcCase) {
//	//t.Helper()
//	if ans := Multiply(c.A, c.B); ans != c.Expected {
//		t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
//	}
//
//}
//
//func TestMultiply(t *testing.T) {
//	createMulTestCase(t, &calcCase{2, 3, 6})
//	createMulTestCase(t, &calcCase{2, -3, -6})
//	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
//}

//func TestAdd(t *testing.T){
//	sum:=Add(1,2)
//	if sum==-1{
//		t.Fatal("TestAdd Pass!")
//	}
//	t.Error("TestAdd fail! Add 1 and 2 result isn't 3!")
//}
//
//func TestSubtract(t *testing.T){
//	subtract:=Subtract(1,2)
//	if subtract==-1{
//		t.Fatal( "TestSubtract Pass!")
//	}
//	t.Error("TestSubtract fail! 1 subtract 2 result isn't -1!")
//}
//
////func TestMultiAdd(t *testing.T){
////	var tests=[]struct{
////		Arg1 int
////		Arg2 int
////		Sum int
////	}{
////		{1, 1, 2},
////		{-1, -1, -2},
////	    {1,-1,0},
////	    {0,1,0},
////	}
////	for _,test:=range tests{
////		sum := Add(test.Arg1,test.Arg2)
////		if sum != test.Sum{
////			t.Errorf("TestMultiAdd fail! Add %v and %v result isn't %v!",test.Arg1,test.Arg2,test.Sum)
////		}
////	}
////}
//
//func TestMultiAdd(t *testing.T){
//	var tests=map[string]struct{
//		Arg1 int
//		Arg2 int
//		Sum int
//	}{
//		"TestMultiAdd case1":{1, 1, 2},
//		"TestMultiAdd case2":{-1, -1, -2},
//		"TestMultiAdd case3":{1,-1,0},
//		"TestMultiAdd case4":{0,1,0},
//	}
//	for caseName,test:=range tests{
//		sum := Add(test.Arg1,test.Arg2)
//		if sum != test.Sum{
//			t.Errorf("%v fail! Add %v and %v result isn't %v!",caseName,test.Arg1,test.Arg2,test.Sum)
//		}else{
//			t.Log(caseName,"success!")
//		}
//	}
//}
//
//func TestMul(t *testing.T) {
//	cases := []struct {
//		Name           string
//		A, B, Expected int
//	}{
//		{"pos", 2, 3, 6},
//		{"neg", 2, -3, -6},
//		{"zero", 2, 0, 0},
//	}
//
//	for _, c := range cases {
//		t.Run(c.Name, func(t *testing.T) {
//			if ans := Multiply(c.A, c.B); ans != c.Expected {
//				t.Fatalf("%d * %d expected %d, but %d got",
//					c.A, c.B, c.Expected, ans)
//			}
//		})
//	}
//}
//


func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	fmt.Println("111111111111111")
	code := m.Run()
	fmt.Println("222222222222222")
	teardown()
	os.Exit(code)
}