package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
}