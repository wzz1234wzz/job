package study

import (
	"bytes"
	"fmt"
)

func BytesMain(){
	var b = []byte("seafood")  //强制类型转换

	a := bytes.ToUpper(b)

	fmt.Println(a, b)     //输出结果   [83 69 65 70 79 79 68] [115 101 97 102 111 111 100]
	fmt.Printf("b[0]=%v,type=%T\n",b[0],b[0]) // b[0]=115,type=uint8

	c := b[0:4]
	c[0] = 'A'
	fmt.Println(c, b)     //输出结果   [65 101 97 102] [65 101 97 102 111 111 100]

	s1 := "Φφϕ kKK"
	s2 := "ϕΦφ KkK"

	// 看看 s1 里面是什么
	for _, c := range s1 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	// 看看 s2 里面是什么
	for _, c := range s2 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	// 看看 s1 和 s2 是否相似
	fmt.Println(bytes.EqualFold([]byte(s1), []byte(s2)))

	bs := [][]byte{                        //[][]byte   类似于  []string     字节切片 二维数组
		[]byte("Hello World !"),
		[]byte("Hello 世界！"),
		[]byte("hello golang ."),
	}
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!！. "), r)    //判断r字符是否包含在    "!！. "  内
	}
	for _, b := range bs {          //range bs  取得下标和[]byte
		fmt.Printf("%q\n", bytes.TrimFunc(b, f))         //去掉两边满足函数的字符
	}
	// "Hello World"
	// "Hello 世界"
	// "Hello Golang"
	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimPrefix(b, []byte("Hello "))) //去掉前缀
	}
	fmt.Println("----------------------------------------------------------------")
	b1 := []byte("  Hello   World !  ")
	fmt.Printf("%q\n", bytes.Split(b1, []byte{' '}))
	// ["" "" "Hello" "" "" "World" "!" "" ""]
	fmt.Printf("%q\n", bytes.Fields(b1))
	// ["Hello" "World" "!"]
	f1 := func(r rune) bool {
		return bytes.ContainsRune([]byte(" !"), r)
	}
	fmt.Printf("%q\n", bytes.FieldsFunc(b, f1))
	// ["Hello" "World"]

	fmt.Println("----------------------------------------------------------------")
	b3 := []byte("Hello World!")
	b2 := []byte("Hello 世界！")
	buf := make([]byte, 6)
	rd := bytes.NewReader(b3)
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "World!"

	rd.Reset(b2)
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	fmt.Printf("Size:%d, Len:%d\n", rd.Size(), rd.Len())

	fmt.Println("----------------------------------------------------------------")
	rd2 := bytes.NewBufferString("Hello World!")
	buf2 := make([]byte, 6)
	// 获取数据切片
	b4 := rd2.Bytes()
	// 读出一部分数据，看看切片有没有变化
	rd2.Read(buf2)
	fmt.Printf("%s\n", rd2.String()) // World!
	fmt.Printf("%s\n", b4)         // Hello World!
	fmt.Printf("%s\n\n", buf2) //[72 101 108 108 111 32]

	// 写入一部分数据，看看切片有没有变化
	rd2.Write([]byte("abcdefg"))
	fmt.Printf("%s\n", rd2.String()) // World!abcdefg
	fmt.Printf("%s\n\n", b4)         // Hello World!

	// 再读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd2.String()) // abcdefg
	fmt.Printf("%s\n", b4)           // Hello World!
}
