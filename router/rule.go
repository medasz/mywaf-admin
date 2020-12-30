package router

import (
	"fmt"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
)

func EditRule(ctx *macaron.Context, x csrf.CSRF, sess session.Store) {
	id := ctx.ParamsInt64(":Id")
	rule, err := model.GetRuleById(id)
	if err != nil {
		log.Printf("Fail to get rule by id %v :%s", id, err)
		ErrorMsg(ctx, "获取规则失败!")
	}
	ctx.Data["Id"] = id
	ctx.Data["user"] = sess.Get("login").(string)
	ctx.Data["rule"] = rule.RuleItem
	ctx.Data["ruleType"] = rule.RuleType
	ctx.Data["rulesType"] = model.RuleInfo
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "editRule")
}

func DeleteRule(ctx *macaron.Context, sess session.Store) {
	ruleType := sess.Get("ruleType")
	sess.Delete("ruleType")
	id := ctx.ParamsInt64(":id")
	n, err := model.DeleteRule(id)
	if err != nil {
		log.Println("Fail to delete rule :", err)
		ErrorMsg(ctx, "删除规则失败!")
	}
	if n > 0 {
		err := model.UpdataFlag("rule")
		if err != nil {
			log.Println(err)
			//ErrorMsg(ctx, "更新uuid出错!")
			//return
			ErrorMsg(ctx, "更新uuid出错!")
		}
	}
	ctx.Redirect(fmt.Sprintf("/index/rule/%s", ruleType))
}

func EditRuleDeal(ctx *macaron.Context, sess session.Store) {
	ruleType := sess.Get("ruleType").(string)
	sess.Delete("ruleType")
	id := ctx.ParamsInt64(":Id")
	rule := ctx.Req.FormValue("rule")
	log.Printf("rule:%s;id:%v", rule, id)
	n, err := model.UpdateRuleById(id, rule)
	if err != nil {
		log.Printf("更新规则失败! id:%v;rule:%s :%s", id, rule, err)
		ErrorMsg(ctx, "更新规则失败!")
	}
	if n > 0 {
		err := model.UpdataFlag("rule")
		if err != nil {
			log.Println(err)
			//ErrorMsg(ctx, "更新uuid出错!")
			//return
			ErrorMsg(ctx, "更新uuid出错!")
		}
	}
	ctx.Redirect(fmt.Sprintf("/index/rule/%s", ruleType))
}

func AddRuleDeal(ctx *macaron.Context) {
	rule := ctx.Req.PostForm.Get("rule")
	ruleType := ctx.Req.PostForm.Get("ruleType")
	log.Printf("rule:%s;ruleType:%s", rule, ruleType)
	n, err := model.AddRule(rule, ruleType)
	if err != nil {
		log.Printf("Fail to insert rule{RuleItem:%s,RuleType:%s}", rule, ruleType)
		ErrorMsg(ctx, "添加规则失败!")
	}
	if n > 0 {
		err := model.UpdataFlag("rule")
		if err != nil {
			log.Println(err)
			//ErrorMsg(ctx, "更新uuid出错!")
			//return
			ErrorMsg(ctx, "更新uuid出错!")
		}
	}
	ctx.Redirect(fmt.Sprintf("/index/rule/%s", ruleType))
}

//func SyncRule(ctx *macaron.Context) {
//	ruleType := ctx.Params(":ruleType")
//	log.Printf(ruleType)
//	rules, err := model.GetRule(ruleType)
//	if err != nil {
//		log.Printf("Fail to find rules with ruleType:%s :%s", ruleType, err)
//		ErrorMsg(ctx, "获取规则列表失败，请查看日志!")
//	}
//	content, err := json.Marshal(rules)
//	if err != nil {
//		log.Println(err)
//		ErrorMsg(ctx, "规则json序列化失败!")
//	}
//	filename := fmt.Sprintf("%s.rule", ruleType)
//	filepath := path.Join(setting.RulePath, filename)
//	err = utils.OverFile(content, filepath)
//	if err != nil {
//		log.Printf("Fail to over rule file:%s :%s", filepath, err)
//		ErrorMsg(ctx, "同步失败!")
//	}
//
//	err = utils.ReloadRule()
//	if err != nil {
//		log.Println(err)
//		ErrorMsg(ctx, "重载规则失败!")
//	}
//
//	ErrorMsg(ctx, "成功!")
//}

type rules struct {
	Result string              `json:"result"`
	Rules  map[string][]string `json:"rules"`
	Uuid   string              `json:"uuid"`
}

func GetRuleJson(ctx *macaron.Context) {
	resData := rules{Result: "true"}
	res, err := model.GetAllRule()
	if err != nil {
		log.Printf("Fail to get rules:%s", err)
		resData.Result = "false"
		ctx.JSON(200, resData)
	}
	ruleMap := map[string][]string{}
	for _, y := range res {
		ruleMap[y.RuleType+".rule"] = append(ruleMap[y.RuleType+".rule"], y.RuleItem)
	}
	resData.Rules = ruleMap
	resData.Uuid, err = model.GetFlag("rule")
	if err != nil {
		log.Println(err)
		resData.Result = "false"
		ctx.JSON(200, resData)
	}
	ctx.JSON(200, resData)
}
