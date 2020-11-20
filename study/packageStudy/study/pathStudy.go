package study

import (
	"path"
	"fmt"
)

func  PathMain()  {
	filepath := "/mnt/e/work1/1/51/grid"
	fileFullPath := path.Join(filepath,"grid")
	fmt.Println(fileFullPath)
}