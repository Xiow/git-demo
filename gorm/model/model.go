package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	Id int `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Pssword string
	Gender string
	//PassWord string `json:"pass_word" gorm:"column:pass_word"`
}
var DB *gorm.DB

func InitMysql()*gorm.DB{
	var err error
	DB,err=gorm.Open(mysql.Open(viper.GetString("Mysql.dsn")))
	if err!=nil{
		fmt.Println("Connect Mysql Failed")
	}else{
		fmt.Println("连接数据库成功")
	}
	DB.AutoMigrate(&User{})
	return DB
}
//user1:=User{2, "简薇","123456","女"}
//user2:=User{3,"李小允","123456","女"}
//user3:=User{4, "昭阳","123456","男"}
//user4:=User{5,"方圆","123456","男"}
//user5:=User{6,"向晨","123456","男"}
//user6:=User{7,"CC","123456","女"}
