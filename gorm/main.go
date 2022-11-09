
package main

import (
	"gorm/config"
	"gorm/model"
)

func main() {
	config.InitConfig()
	model.InitMysql()

}