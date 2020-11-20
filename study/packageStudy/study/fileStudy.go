package study

import (
	"fmt"
	"io"
	"os"
	"bufio"
)

func FileMain() {
	// 只读方式打开当前目录下的demo.txt文件
	file, err := os.Open("E:/Baron/study/packageStudy/demo.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}

	// 使用Read方法读取数据
	var tmp = make([]byte, 4)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))

	var content []byte
	var tmp2 = make([]byte, 2)
	for {
		n, err := file.Read(tmp2)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp2[:n]...)
	}
	fmt.Println(string(content))
	file.Close()

	fmt.Println("-----------------------------bufio------------------------------------------------")
	file1, err := os.Open("E:/Baron/study/packageStudy/demo.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}