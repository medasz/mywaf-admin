package router

import (
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
	"mywaf-admin/utils"
)

//session会话检查
func SessionCheck(ctx *macaron.Context, sess session.Store) {
	if sess.Get("login") == nil {
		ctx.Redirect("/")
	}
}

//返回登录界面
func LoginHtml(ctx *macaron.Context,x csrf.CSRF) {
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "login")
}

//用户登录
func LoginDeal(ctx *macaron.Context, cpt *captcha.Captcha, sess session.Store) {
	if cpt.VerifyReq(ctx.Req) {
		username := ctx.Req.Form.Get("username")
		password := ctx.Req.Form.Get("password")
		password = utils.Encryption(password)
		flag, err := model.GetUser(username, password)
		if err == nil && flag {
			err := sess.Set("login", username)
			if err != nil {
				log.Println("Fail to set session :", err)
				ctx.Redirect("/")
			}
			ctx.Redirect("/index/rule/white_ip")
		}else{
			ErrorMsg(ctx,"用户名或密码错误!")
		}
	} else {
		ErrorMsg(ctx,"验证码错误")
	}
}

//返回index页面
func GetIndex(ctx *macaron.Context,sess session.Store,x csrf.CSRF) {
	ruleType:=ctx.Params(":ruleType")
	rules,err:=model.GetRule(ruleType)
	if err!=nil{
		log.Println("Fail to find rules with ruleType :",err)
		ErrorMsg(ctx,"获取规则列表失败，请查看日志!")
	}
	ctx.Data["rules"]=rules
	ctx.Data["rulesType"]=model.RuleInfo
	ctx.Data["user"]=sess.Get("login").(string)
	ctx.Data["ruleType"]=ruleType
	sess.Set("ruleType",ruleType)
	ctx.Data["csrf_token"]=x.GetToken()
	ctx.HTML(200,"index")
}

//Logout
func Logout(ctx *macaron.Context,sess session.Store){
	err:=sess.Delete("login")
	if err==nil{
		ctx.Redirect("/")
	}
	log.Printf("Fail to delete session :",err)
}
