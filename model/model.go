package model

import (
	"fmt"
	"log"
	"mywaf-admin/setting"
	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
	err error
)
func init() {
	url:=fmt.Sprintf("%s:%s@(%s:%v)/%s",setting.Username,setting.Password,setting.Host,setting.Port,setting.Db)
	Engine,err=xorm.NewEngine("mysql",url)
	if err!=nil{
		log.Panicln("Fail to connect mysql :",err)
	}
	log.Println("success to connect mysql")

}