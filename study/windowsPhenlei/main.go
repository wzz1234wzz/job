package main

import (
	"fmt"
	"os/exec"
)
const (
	WindowsSolverShellScript     = "PHengLEIv3d0.exe >> cfd.log"
	WindowsStopSolverShellScript = "taskkill /f /im PHengLEIv3d0.exe"
	GridToWebGridScriptTpl = "cmdgrid %q %q %q"
	//gridscript="cmdgrid \"default\" \"E:/work1/1/113/grid/m6_str_0.bcmesh\" \"E:/fts/1/113/default\""
	gridscript="\"default\" \"E:/work1/1/113/grid/m6_str_0.bcmesh\" \"E:/fts/1/113/default\""
)

func runWindowsCmd(workDir string, cmd string) (string, error){
	fmt.Println("Running Windows cmd: "+cmd)
	command := exec.Command("cmd", "/c", cmd)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}

// RunWindowsCmdPost
func RunWindowsCmdPost(workDir, srcPath, dstPath string) (string, error){
	command := exec.Command("cmdpost", srcPath,dstPath)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}

func runWindowsCmds(workDir string) error{
	fmt.Println("Running Windows cmd: ")
	command := exec.Command("cmdgrid", "default","E:/work1/1/113/grid/m6_str_0.bcmesh","E:/fts/1/113/default")
	command.Dir = workDir
	err := command.Run()
	return err
}

func runWindowsCmdgrid(workDir, bcName, srcPath, dstPath string) (string, error){
	fmt.Println("Running Windows cmd: ")
	command := exec.Command("cmdgrid", bcName,srcPath,dstPath)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}

func ExecString(workDir string, cmd string) (string, error) {
	command := exec.Command("sh", "-c", cmd)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}

func WindowsDoStopSolver(workDir string) error {
	_, err := runWindowsCmd(workDir, WindowsStopSolverShellScript)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	a:=0
	if a==0{ // windows
		//workDir:="D:\\wzz\\study\\windowsPhenlei\\m6\\Ma10H10α10"
		//workDir:="D:/wzz/study/windowsPhenlei/m6/Ma10H10α10"
		//workDir:="D:/wzz/study/windowsPhenlei/m6"
		//workDir:="E:/work1/1/113"
		//_,err := runWindowsCmd(workDir,WindowsStopSolverShellScript)
		//out,err := runWindowsCmd(workDir,WindowsSolverShellScript )
		//script := fmt.Sprintf(GridToWebGridScriptTpl, "default", "E:/work1/1/113/grid/m6_str_0.bcmesh", "E:/fts/1/113/default")
		//out,err := runWindowsCmd(workDir,script)
		workDir := "./"

		//out,err := runWindowsCmd(workDir,gridscript)
		//fmt.Println("error=",err)
		//if err != nil && err.Error() != "exit status 1"{
		//	fmt.Println("err.Error() == \"exit status 1\"=",err.Error() == "exit status 1")
		//}
		//fmt.Println("out=",out)

		//err := runWindowsCmds(workDir)
		//fmt.Println("error=",err)
		//if err != nil && err.Error() != "exit status 1"{
		//	fmt.Println("err.Error() == \"exit status 1\"=",err.Error() == "exit status 1")
		//}
		//fmt.Println("out=",err)
		//bcName:="default"
		//srcPath:="E:/work1/1/113/grid/m6_str_0.bcmesh"
		//dstPath:="E:/fts/1/113/default"
		//out,err := runWindowsCmdgrid(workDir, bcName, srcPath, dstPath)
		//if err != nil && err.Error() != "exit status 1"{
		//	fmt.Println("err.Error() =",err)
		//}
		//fmt.Println("out=",out)
		srcPath:= "E:/work1/1/129/Ma0.8395H0α3.06/results/tecflow.plt.bak"
		dstPath:= "E:/post/1/129/223"
		out,err := RunWindowsCmdPost(workDir, srcPath, dstPath)
		if err != nil && err.Error() != "exit status 1"{
			fmt.Println("err.Error() =",err)
		}
		fmt.Println("out=",out)
	}else{ // linux
		workDir:="/mnt/d/wzz/study/windowsPhenlei/m6/Ma10H10α10"
		//out,err := ExecString(workDir,LocalSolverShellScript)
		script := fmt.Sprintf(GridToWebGridScriptTpl, "default", "/mnt/e/work1/1/113/grid/m6_str_0.bcmesh", "/mnt/e/fts/1/113/default")
		out,err := ExecString(workDir,script)
		if err != nil {
			fmt.Println("error=",err)
		}
		fmt.Println("out=",out)
	}
}

