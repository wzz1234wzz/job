package method

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func LoadConfigFromYaml (c *config) error  {
	c.v = viper.New()

	//设置配置文件的名字
	c.v.SetConfigName("config")

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	c.v.AddConfigPath("%GOPATH/src/")
	c.v.AddConfigPath("./")

	//设置配置文件类型
	c.v.SetConfigType("yaml");

	if err := c.v.ReadInConfig(); err != nil{
		return  err;
	}

	log.Printf("age: %s, name: %s \n", c.v.Get("information.age"), c.v.Get("information.name"));
	return nil;
}

//将配置解析为Struct对象
func UmshalStruct(c *config) error  {
	LoadConfigFromYaml(c)
	var cf YamlConfig
	if err := c.v.Unmarshal(&cf); err != nil{
		return err
	}
	fmt.Println(cf)
	return nil
}

func YamlStringSettings(c *config) string {
	c.v = viper.New();
	c.v.Set("name", "wzp");
	c.v.Set("age", 18);
	c.v.Set("aliase",[]string{"one","two","three"})

	cf := c.v.AllSettings()
	bs, err := yaml.Marshal(cf)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}

func JsonStringSettings(c *config) string {
	c.v = viper.New();
	c.v.Set("name", "wzp");
	c.v.Set("age", 18);
	c.v.Set("aliase",[]string{"one","two","three"})

	cf := c.v.AllSettings()
	bs, err := json.Marshal(cf)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}
