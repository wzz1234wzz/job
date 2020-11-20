package main
import (
	"fmt"
	"yamlStudy/method"
	"yamlStudy/mode"
)

func main(){
	var c mode.Config
	method.UmshalStruct(&c)
	fmt.Println(method.YamlStringSettings(&c))
	fmt.Println(method.JsonStringSettings(&c))
}