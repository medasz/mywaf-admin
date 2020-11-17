package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mywaf-admin/setting"
	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
	err    error
)

func init() {
	//连接mysql
	url := fmt.Sprintf("%s:%s@(%s:%v)/%s", setting.Username, setting.Password, setting.Host, setting.Port, setting.Db)
	Engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		log.Panicln("Fail to connect mysql :", err)
	}
	log.Println("success to connect mysql")

	//同步数据结构到数据库
	err = Engine.Sync2(new(User))
	if err != nil {
		log.Panicln("Fail to create user table :", err)
	}
	log.Println("success to create user table!")

	err = Engine.Sync2(new(Rule))
	if err != nil {
		log.Panicln("Fail to create rule table :", err)
	}
	log.Println("success to create rule table!")

	//检测user表记录是否存在
	has, err := Engine.IsTableEmpty(new(User))
	if err != nil {
		log.Panicln("Fail to isTableEmpty user :", err)
	}
	log.Println("success to isTableEmpty user!")
	if has {
		log.Println("user is null,and insert admin/password")
		err = NewUser("admin", "password")
		if err != nil {
			log.Panicln("Fail to insert user admin/password :", err)
		}
		log.Println("Success to insert user admin/password!")
	}else{
		log.Println("user is not null!")
	}

	//检测rule表记录是否存在
	has, err = Engine.IsTableEmpty(new(Rule))
	if err != nil {
		log.Panicln("Fail to isTableEmpty rule :", err)
	}
	log.Println("Success to isTableEmpty rule!")
	if has {
		log.Println("rule is null,insert rules")
		err:=initRule()
		if err!=nil{
			log.Panicln("Fail to init rule :",err)
		}
		log.Println("Success to init rule!")
	}else {
		log.Println("rule is not null!")
	}
}
