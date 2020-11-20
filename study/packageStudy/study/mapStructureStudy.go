package study

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age int
	Job string
	Other map[string]interface{} `mapstructure:",remain"`
}

type  Cat struct {
	Name string
	Age int
	Breed string
}

func MapStuctureMain(){
	datas := []string{`
	{
		"type": "person",
		"name":"dj",
		"age":18,
		"job":"programmer",
		"height":"1.8m",
      	"handsome": true
	}
	`,
	`
	{
		"type": "cat",
		"name":"kitty",
		"Age":1,
		"breed":"Ragdoll"
	}
	`,
	}
    // 将map[string]interface{}解码到 Go 结构体中
	for _, data :=range datas{
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data),&m) // []byte(data) 全是数字
		//fmt.Println("[]byte(data)",[]byte(data))
		if err !=nil{
			log.Fatal(err)
		}
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m,&p)
			fmt.Println("person",p)
		case "cat":
			var cat Cat
			mapstructure.Decode(m,&cat)
			fmt.Println("cat",cat)
		}
	}

	// 将 Go 结构体反向解码为map[string]interface{}
	p := &Person{
		Name: "dj",
		Age:  18,
	}
	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	data, _ := json.Marshal(m)
	fmt.Println(string(data))

}
