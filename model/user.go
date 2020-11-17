package model

import (
	"log"
	"mywaf-admin/utils"
)

type User struct {
	Id int64
	Username string `xorm:"unique not null"`
	Password string	`xorm:"not null"`
}

//添加用户
func NewUser(username,password string)error{
	password=utils.Encryption(password)
	_,err:=Engine.Insert(&User{Username: username,Password: password})
	return err
}

//查询用户
func GetUser(username,password string)(bool,error){
	flag,err:=Engine.Get(&User{Username: username,Password: password})
	if err!=nil{
		log.Printf("Fail to get user User{Username:%v,Password:%v} :%s",username,password,err)
		return false,err
	}
	return flag,err
}

//获取所有用户
func GetAllUser()([]User,error){
	userList:=[]User{}
	err:=Engine.Find(&userList)
	return userList,err
}

//删除用户
func DeleteUserById(id int64)error{
	_,err:=Engine.Delete(&User{Id: id})
	return err
}

//根据id查询用户
func GetUserById(id int64)(User,error){
	user:=User{Id: id}
	_,err:=Engine.Get(&user)
	return user,err
}

//更新用户信息
func UpdateUser(id int64,username,password string)error{
	_,err:=Engine.ID(id).Update(&User{Username: username,Password: password})
	return err
}