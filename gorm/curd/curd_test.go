package curd

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm/model"
	"testing"
)

//查询单条记录
func TestFindOne(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	var user model.User
	result:=model.DB.Debug().First(&user)
	//result拿到受影响的函数和error
	t.Log(result.RowsAffected)
	t.Log(result.Error)
	//打印拿取道德结构体
	t.Log("根据主键拿到第一条记录",user.Id,user.Name,user.Gender,user.Pssword)
	//var user1 model.User
	//model.DB.Debug().Last(&user1)
	//t.Log("根据主键拿到最后一条数据",user1)
}
//根据主键查找记录
func TestFindFirstByPrimaryKey(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	//根据主键去获取查找到的第一条数据
	var user model.User
	model.DB.First(&user,4)
	t.Log("根据主键拿到记录:",user)
}
//查找所有记录
func TestFindAllRecords(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	userList:=make([]model.User,10)
	model.DB.Debug().Find(&userList)
	//fmt.Println(userList)
	for key,value:=range userList{
		t.Log(key,value)
	}

}
//根据条件查找记录
func TestFindRecordByContions(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	//该种方法只能查找单条数据
	var user model.User
	model.DB.Debug().Where("name=?","昭阳").First(&user)
	t.Log(user)
	//查找符合条件的多条数据(跟slice)
	userList:=make([]model.User,0)
	model.DB.Debug().Where("name = ?","简薇").Find(&userList)
	t.Log(userList)
}
//模糊查询
func TestFindByM(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	userList:=make([]model.User	,10)
	model.DB.Where("name LIKE ?","%下她%").Find(&userList)
	t.Log(userList)
}
//字段更新
func TestUpdate(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	//通过条件判断更新单个列表
	//更新单条记录的多个列额时候,先把记录拿到，再去更新
	var user model.User
	model.DB.Debug().Where("id=?",10).Find(&user)
	model.DB.Debug().Model(&user).Updates(model.User{Name:"老北京",Pssword:"999999"})
	//m:=map[string]string{"name":"老北京","password":"666666"}
	//model.DB.Debug().Model(&user).Where("id=?",10).UpdateColumns(m)
	//model.DB.Debug().Model(&user).Updates(model.User{Name:"今天番茄来炒蛋", Pssword:"今天天气很好,适合出去走走"})
	t.Log(user)
}
//删除记录

func TestDeleteRecord(t *testing.T){
	viper.AddConfigPath("C:\\\\Users\\\\james\\\\Desktop\\\\GInchat\\\\gorm\\\\config")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	model.InitMysql()
	//1.根据主键删除记录
	//var user model.User
	users:=make([]model.User,0)
	model.DB.Debug().Delete(&users,[]int{7,8,10})
	//user.Id=9
	//model.DB.Debug().Delete(&user)
	t.Log(users)

}
