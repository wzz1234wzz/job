package study

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)
func IoutilMain(){
	content, err := ioutil.ReadFile("./demo.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file content: %s,\n 类型：%T\n", content,content)

	s:=string(content)
	fmt.Printf("内容:\n%v\n 类型：%T,长度l=%d\n", s,s,len(s))
	fmt.Printf("ft:",strings.Contains(s, "Co"))

	files, err := ioutil.ReadDir("./a")
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	fmt.Println("获取文件，并输出它们的名字")
	for _, file := range files {
		println(file.Name())
	}
	fmt.Println("复制文件")
	err1 := CopyAllFile("/mnt/e/Baron/study/packageStudy/b", "/mnt/e/Baron/study/packageStudy/a")
	if err != nil {
		fmt.Println(err1)
	}

	CopyFile("/mnt/e/Baron/study/packageStudy/study/b.txt", "/mnt/e/Baron/study/packageStudy/study/data.txt")
}

// CopyAllFile 拷贝文件函数
func CopyAllFile(dst, src string) (err error) {
	ok, err := IsDir(src)
	if err !=nil{
		return err
	}
	if ok{ // if file is dir
		projectFiles, err2 := ioutil.ReadDir(src)
		if err2 != nil{
			return err2
		}
		for _, projectFile := range projectFiles {
			fileName := projectFile.Name()
			newSrcFile := path.Join(src, fileName)
			newDstFile := path.Join(dst, fileName)
			isok, _ := PathExists(dst)
			if !isok {
				err = os.MkdirAll(dst, 0664)
				if err !=nil{
					return err
				}
			}
			err3 := CopyAllFile(newDstFile, newSrcFile)
			if err3 != nil{
				return err3
			}
		}
	}else{
		_, err = CopyFile(dst, src)
		if err != nil{
			return err
		}
	}
	return nil
}

// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName) // 以读方式打开源文件
	if err != nil {
		return -1, err
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_TRUNC| os.O_CREATE, 0600)
	if err != nil {
		return -1, err
	}

	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

func IsDir(file string) (bool, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, err
	}
	return false, errors.New("unknown file type")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}