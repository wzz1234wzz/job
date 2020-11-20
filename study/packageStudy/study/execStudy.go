package study

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func lookPath(exe string) {
	path, err := exec.LookPath(exe)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(exe, "is at path: ", path)
}

func sliceFunc(cmd... string){
	fmt.Println(cmd)
	if len(cmd) == 0 {
		fmt.Printf("Usage: %s args...\n", os.Args[0])
		os.Exit(-1)
	}
	fmt.Println(cmdFunc(cmd...))
}

func cmdFunc(cmd... string) string {
	fmt.Printf("cmd slice len: %d, value:%v\n", len(cmd),  cmd)
	result, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		fmt.Println("Command failed:", err.Error())
	}

	//  return string(result)  // with '\n'
	return strings.TrimSpace(string(result))
}

func runInLinux(cmd string) string{
	fmt.Println("Running Linux cmd:" , cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func runInWindows(cmd string) string{
	fmt.Println("Running Win cmd:", cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func RunCommand(cmd string) string{
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func RunLinuxCommand(cmd string) string{
	if runtime.GOOS == "windows" {
		return ""
	} else {
		return runInLinux(cmd)
	}
}

func runInLinuxWithErr(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:"+cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func runInWindowsWithErr(cmd string) (string, error){
	fmt.Println("Running Windows cmd:"+cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func RunCommandWithErr(cmd string) (string, error){
	if runtime.GOOS == "windows" {
		return runInWindowsWithErr(cmd)
	} else {
		return runInLinuxWithErr(cmd)
	}
}


func ExecMain() {
	var k int
	k=2
	if k == 0 {
		lookPath("pwd")

		cmd := exec.Command("ls", "-lah")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("cmd.Run() failed with %s\n", err)
		}
		fmt.Printf("combined out:\n%s\n", string(out))
		sliceFunc(os.Args[1:]...)

	}else if k==1{
		fmt.Println("len=",len(os.Args))
		if len(os.Args) == 1{
			fmt.Printf("Usage: %s args...\n", os.Args[0])
			os.Exit(-1)
		}
		str1, err := RunCommandWithErr(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(str1)
		}

		str := RunCommand (os.Args[1])
		fmt.Println(str)
	}else if k==2{
		cmd := exec.Command("tr", "a-z", "A-Z")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("in all caps: %q\n", out.String())//in all caps: "SOME INPUT"
	}
}

