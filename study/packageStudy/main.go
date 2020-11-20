package main

import (
	"flag"
	"packageStudy/study"
)

func main() {
	var num int
	flag.IntVar(&num,"id",0,"输入一个数字")

	//用flag.Parse()解析后
	flag.Parse()

	switch num{
	case 0:
		study.ReflectMain()
	case 1:
		study.OsMain()
	case 2:
		study.BytesMain()
	case 3:
		study.IoutilMain()
	case 4:
		study.ExecMain()
	case 5:
		study.FileMain()
	case 6:
		study.RegexMain()
	case 7:
		study.StringMain()
	case 8:
		study.PathMain()
	case 9:
		study.ZipMain()
	case 10:
		study.CgoMain()
	case 11:
		study.ByteMain()
	case 12:
		study.MapStuctureMain()
	case 13:
		study.SyncMain()
	case 14:
		study.CasbinMain()
	}
}
