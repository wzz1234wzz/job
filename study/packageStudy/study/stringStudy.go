package study

import (
	"fmt"
	"strings"
)

func StringMain(){
	fmt.Println("-------------------Trim------------------------")
	fmt.Println(strings.TrimRight("abba", "ba"))         // 结果为空，即从右边开始，碰到与“ba"中的一个即删除
	fmt.Println(strings.TrimRight("abcdaa8aaa", "abcd")) // 结果为abcdaa8
	fmt.Println(strings.TrimSuffix("abcddcba", "dcba"))   // 结果为abcddcba
	s:=";ff;dd/d;;/"
	fmt.Println(strings.Trim(s, "\""))

	line:=" nBoundaryConditons = 4\""
	arr := strings.Split(line, "=") // 以=切割

	key := strings.TrimSpace(arr[0]) // 去除两端空格
	value := strings.Trim(strings.TrimSpace(arr[1]), "\"")
	fmt.Println("key:",key)
	fmt.Println("value:",value)
	a:="123"
	fmt.Println("a:=%q\n",a) // %q为加引号
	fmt.Println("a:=%s\n",a)

    fmt.Println("-------------------根据空格键进行分割------------------------")
	s1:=" ab cd          ef gh ij kl "
	arr1:=strings.Fields(s1) // arr:["ab" "cd" "ef" "gh" "ij" "kl"]
	fmt.Printf("arr:%q\n",arr1)
	arr2:=strings.Split(s1," ") // arr:["" "ab" "cd" "" "" "" "" "" "" "" "" "" "ef" "gh" "ij" "kl" ""]
	fmt.Printf("arr:%q\n",arr2)

	fmt.Println("-------------------根据空格键进行分割------------------------")
	s2:="iiaiibiiciiiidiiii"
	sep:="ii"
	arr3:=strings.Split(s2,sep)
	fmt.Println("arr:",arr3)

	fmt.Println("-------------------其他------------------------")
	s4 := " aBc"
	s5 := "100a"
	s3 := s4 + s5
	fmt.Println(s3) // aBc100a
	fmt.Println(strings.HasPrefix(s3, "a")) //判断前缀
	fmt.Println(strings.HasSuffix(s3, "0")) //判断后缀
	fmt.Println(strings.Contains(s3, "9"))  //字符串包含关系
	fmt.Println(strings.Index(s3, "0"))     //判断子字符串或字符在父字符串中出现的位置（索引）5
	fmt.Println(strings.LastIndex(s3, "0")) //最后出现位置的索引 6
	fmt.Println(strings.Replace(s3,"0","1",-1))//如果 n = -1 则替换所有字符串  aBc111a
	fmt.Println(strings.Count(s3,"0"))//出现的非重叠次数 2
	fmt.Println(strings.Repeat(s3,2))//重复字符串  aBc100a aBc100a
	fmt.Println(strings.ToLower(s3))//修改字符串大小写 abc100a
	fmt.Println(strings.ToUpper(s3))//修改字符串大小写 ABC100A
	fmt.Println(strings.TrimSpace(s3))//修剪字符串 去掉开头和结尾空格
	fmt.Println(strings.Trim(strings.TrimSpace(s3),"a"))//修剪字符串 去掉开头和结尾字符串 Bc100
}