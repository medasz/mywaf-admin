package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	//server
	WebInfo string
	Mode string
	//NginxBin string

	//mysql
	Host string
	Port int
	Db string
	Username string
	Password string

	//rules
	//RulePath string
)

func init() {
	cfg,err:=ini.Load("conf/app.ini")
	if err!=nil{
		log.Panicln("Fail to read file: %v",err)
	}
	WebInfo=cfg.Section("server").Key("web").MustString("127.0.0.1:5000")
	Host=cfg.Section("mysql").Key("host").MustString("127.0.0.1")
	Port=cfg.Section("mysql").Key("port").MustInt(3306)
	Db=cfg.Section("mysql").Key("db").MustString("mywaf")
	Username=cfg.Section("mysql").Key("user").MustString("admin")
	Password=cfg.Section("mysql").Key("pass").MustString("password")
	//RulePath=cfg.Section("mywaf").Key("rulePath").MustString("./rules")

	Mode=cfg.Section("server").Key("mode").MustString("prod")
	//NginxBin=cfg.Section("server").Key("nginx_bin").MustString("/opt/openresty/nginx/sbin/nginx")
}