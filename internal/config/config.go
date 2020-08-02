package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

// Config ...
type Config struct {
	RootPath string
}

// New ...
func New() Config {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	v := viper.New()
 
	v.AddConfigPath(path + "/config")     //设置读取的文件路径
	v.SetConfigName("config") //设置读取的文件名
	v.SetConfigType("yml")   //设置文件的类型

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
 
	var config Config
	if err := v.Unmarshal(&config) ; err != nil{
		log.Println("err: ", err)
	}
	log.Println(config.RootPath)

	return config
}