package study


import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"path/filepath"
)

func OsMain() {
	t := 0
	if t == 0 {
		err1 := Zip("/mnt/e/Baron/study/packageStudy/study", "/mnt/e/AAA/study.zip")
		if err1 == nil {
			fmt.Println("压缩成功！")
		}else{
			fmt.Println(err1)
		}

		//err1 = Unzip("wzz0.zip", "/mnt/e/AAA")
		//if err1 == nil {
		//	fmt.Println("解压成功！")
		//}else{
		//	fmt.Println(err1)
		//}
	} else {
		fmt.Println("----------------------------")
		hostname, _ := os.Hostname() // Hostname 函数会返回内核提供的主机名。
		fmt.Printf("hostname=%s\n", hostname)
		fmt.Println("-----------os-----------------")

		//environ := os.Environ() //Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝。
		//fmt.Printf("environ=%v\n", environ)
		fmt.Println("----------------------------")

		getenv := os.Getenv("PATH") // Getenv 函数会检索并返回名为 key 的环境变量的值。如果不存在该环境变量则会返回空字符串。
		fmt.Printf("getenv=%v\n", getenv)
		fmt.Println("----------------------------")

		userId := os.Getuid() // Getuid 函数可以返回调用者的用户 ID。
		fmt.Printf("userId=%v\n", userId)
		fmt.Println("----------------------------")

		groupId := os.Getgid() // Getpid Getuid 函数可以返回调用者的用户 ID。
		fmt.Printf("groupId=%v\n", groupId)
		fmt.Println("----------------------------")

		pId := os.Getpid() // Getpid Getuid 函数可以返回调用者的用户 ID。
		fmt.Printf("pId=%v\n", pId)
		fmt.Println("----------------------------")

		pwd, _ := os.Getwd() // Getwd 函数可以返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd 会返回其中一个
		fmt.Printf("pwd=%v\n", pwd)
		fmt.Println("----------------------------")

		err := os.Mkdir("./testdir", 0777) //Mkdir 函数可以使用指定的权限和名称创建一个目录。如果出错，会返回 *PathError 底层类型的错误
		if err == nil {
			fmt.Println("testdir创建成功")
		} else {
			fmt.Printf("创建失败,err=%v\n", err)
		}
		fmt.Println("----------------------------")

		err_rm := os.Remove("./test.txt") // Remove 函数会删除 name 指定的文件或目录。如果出错，会返回 *PathError 底层类型的错误。RemoveAll 函数跟 Remove 用法一样，区别是会递归的删除所有子目录和文件。
		if err_rm == nil {
			fmt.Println("testdir删除成功")
		} else {
			fmt.Printf("删除失败,err=%v\n", err_rm)
		}
		fmt.Println("----------------------------")

		f, err := exec.LookPath("./") // 在环境变量 PATH 指定的目录中搜索可执行文件，如果 file 中有斜杠，则只在当前目录搜索。返回完整路径或者相对于当前目录的一个相对路径。
		if err != nil {
			fmt.Printf("失败,err=%v\n", err)
		} else {
			fmt.Printf("可执行文件f=%v\n", f)
		}

		fmt.Println("-----------------------------")
		u, _ := user.Current()
		log.Println("用户名：", u.Username)
		log.Println("用户id", u.Uid)
		log.Println("用户主目录：", u.HomeDir)
		log.Println("主组id：", u.Gid)
		s, _ := u.GroupIds() // 用户所在的所有的组的id
		log.Println("用户所在的所有组：", s)
		fmt.Println("----------------------------")

		path := "./a"
		err = os.RemoveAll(path)
		if err == nil {
			fmt.Println("删除成功！")
		}
		c := make(chan os.Signal, 0)
		signal.Notify(c) // 第一个参数表示接收信号的 channel，第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
		signal.Stop(c)   //不允许继续往c中存入内容
		t := <-c         // Block until a signal is received.//c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
		fmt.Println("Got signal:", t)
		fmt.Println("--------运行该程序，然后在 CMD 窗口中通过 Ctrl+C 来结束该程序，便会得到输出结果-------------")
		os.Exit(0) // Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
	}
}

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if ! info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	return err
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}