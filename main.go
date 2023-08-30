package main

import (
	"encoding/json"
	"fmt"
	"github.com/goland-cn/apollo/initialize"
)

type Config struct{} //改成Apollo的相对应配置的结构体

//示例
//type Config struct { //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Apollo 配置
//	Host   string       `json:"Host"`
//	Port   string       `json:"Port"`
//	Mysql  MysqlConfig  `json:"Mysql"`
//	Consul ConsulConfig `json:"Consul"`
//}
//
//type MysqlConfig struct {
//	Username string `json:"Username"`
//	Password string `json:"Password"`
//	Host     string `json:"Host"`
//	Port     string `json:"Port"`
//	DbName   string `json:"DbName"`
//}
//
//type ConsulConfig struct {
//	Host string   `json:"Host"`
//	Port string   `json:"Port"`
//	Name string   `json:"Name"`
//	Tags []string `json:"Tags"`
//}

var (
	cfg = &Config{} //全局配置信息设置
)

func main() {
	//1.定义配置信息
	AppID := "wkxx-1"                            // AppId
	cluster := "test"                            //对应环境
	NamespaceName := "application"               // 替换为您的应用程序所使用的命名空间名称
	Secret := "a42772d5d4014f3b822978afca0144a2" // 密钥管理可以查看
	//2.调用配置信息初始化
	apolloInit, err := initialize.ApolloInit(AppID, cluster, NamespaceName, Secret, true)
	if err != nil {
		return
	}
	//3.指定获取命名空间
	cache := apolloInit.GetConfigCache(NamespaceName)
	//4.获取对应的key
	v, _ := cache.Get("user_srv") //替换为你的Key
	//5.打印信息
	fmt.Println("Apollo获取的配置信息:", v)
	//6.配置信息转换结构体，可以把json配置粘贴到ChatGpt让它转换
	err = json.Unmarshal([]byte(v.(string)), &cfg)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	fmt.Println(cfg) // 打印配置项的值
}
