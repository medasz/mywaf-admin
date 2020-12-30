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
	url := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8", setting.Username, setting.Password, setting.Host, setting.Port, setting.Db)
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

	err = Engine.Sync2(new(Flag))
	if err != nil {
		log.Panicln("Fail to create flag table :", err)
	}
	log.Println("success to create flag table!")

	err = Engine.Sync2(new(WafConfig))
	if err != nil {
		log.Panicln("Fail to create waf_config :",err)
	}
	log.Println("success to create waf_config table!")


	err = Engine.Sync2(new(WafLog))
	if err != nil {
		log.Panicln("Fail to create waf_log table :", err)
	}
	log.Println("success to create waf_log table!")

	//检测user表记录是否为空
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
	} else {
		log.Println("user is not null!")
	}

	//检测rule表记录是否为空
	has, err = Engine.IsTableEmpty(new(Rule))
	if err != nil {
		log.Panicln("Fail to isTableEmpty rule :", err)
	}
	log.Println("Success to isTableEmpty rule!")
	if has {
		log.Println("rule is null,insert rules")
		err := initRule()
		if err != nil {
			log.Panicln("Fail to init rule :", err)
		}
		log.Println("Success to init rule!")
	} else {
		log.Println("rule is not null!")
	}

	//检测flag表是否为空
	has,err=Engine.IsTableEmpty(new(Flag))
	if has{
		log.Println("flag is null,insert flag")
		err:=initFlag()
		if err!=nil{
			log.Panicln("Fail to init flag :",err)
		}
		log.Println("success to init flag!")
	}else{
		log.Println("flag is not null!")
	}

	//检测waf_config表是否为空
	has,err=Engine.IsTableEmpty(new(WafConfig))
	if has{
		log.Println("waf_config is null,insert waf config")
		err:=initWafConfig()
		if err!=nil{
			log.Panicln("Fail to init waf_config :",err)
		}
		log.Println("success to init waf_config!")
	}else{
		log.Println("waf_config is not null!")
	}
}
