package study

import (
	"regexp"
	"fmt"
)

func RegexMain() {
	var arrayNameReg = regexp.MustCompile(`(.*?)\[.*\]`)
	var key = "localCoordAxis0[3]"
	arrayKey := arrayNameReg.FindStringSubmatch(key)
	fmt.Println(arrayKey)

	var key1 = "localCoordAxis0[]"
	arrayKey1 := arrayNameReg.FindStringSubmatch(key1)
	fmt.Println(arrayKey1)
}