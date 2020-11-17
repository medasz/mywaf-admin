package router

import (
	"fmt"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
)

func EditRule(ctx *macaron.Context,x csrf.CSRF,sess session.Store)  {
	id:=ctx.ParamsInt64(":Id")
	rule,err:=model.GetRuleById(id)
	if err!=nil{
		log.Printf("Fail to get rule by id %v :%s",id,err)
		ErrorMsg(ctx,"获取规则失败!")
	}
	ctx.Data["Id"]=id
	ctx.Data["user"]=sess.Get("login").(string)
	ctx.Data["rule"]=rule.RuleItem
	ctx.Data["ruleType"]=rule.RuleType
	ctx.Data["rulesType"]=model.RuleInfo
	ctx.Data["csrf_token"]=x.GetToken()
	ctx.HTML(200,"editRule")
}

func DeleteRule(ctx *macaron.Context,sess session.Store)  {
	ruleType:=sess.Get("ruleType")
	sess.Delete("ruleType")
	id:=ctx.ParamsInt64(":id")
	err:=model.DeleteRule(id)
	if err!=nil{
		log.Println("Fail to delete rule :",err)
		ErrorMsg(ctx,"删除规则失败!")
	}
	ctx.Redirect(fmt.Sprintf("/index/rule/%s",ruleType))
}

func EditRuleDeal(ctx *macaron.Context,sess session.Store){
	ruleType:=sess.Get("ruleType").(string)
	sess.Delete("ruleType")
	id:=ctx.ParamsInt64(":Id")
	rule:=ctx.Req.FormValue("rule")
	log.Printf("rule:%s;id:%v",rule,id)
	err:=model.UpdateRuleById(id,rule)
	if err!=nil{
		log.Printf("更新规则失败! id:%v;rule:%s :%s",id,rule,err)
		ErrorMsg(ctx,"更新规则失败!")
	}
	ctx.Redirect(fmt.Sprintf("/index/rule/%s",ruleType))
}

func AddRuleDeal(ctx *macaron.Context){
	rule:=ctx.Req.PostForm.Get("rule")
	ruleType:=ctx.Req.PostForm.Get("ruleType")
	log.Printf("rule:%s;ruleType:%s",rule,ruleType)
	err:=model.AddRule(rule,ruleType)
	if err!=nil{
		log.Printf("Fail to insert rule{RuleItem:%s,RuleType:%s}",rule,ruleType)
		ErrorMsg(ctx,"添加规则失败!")
	}
	ctx.Redirect(fmt.Sprintf("/index/rule/%s",ruleType))
}

func SyncRule(ctx *macaron.Context){

}