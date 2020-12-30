package router

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
)

func Dashboard(ctx *macaron.Context) {
	ctx.Data["rulesType"] = model.RuleInfo
	totalRequest, err := model.GetAllLogCount()
	if err != nil {
		log.Printf("Fail to get all log count:%s", err)
		ErrorMsg(ctx, "dashboard is error")
		return
	}
	attackRequest, err := model.GetAttackLogCount()
	if err != nil {
		log.Printf("Fail to get attack log count:%s", err)
		ErrorMsg(ctx, "dashboard is error")
		return
	}
	ctx.Data["totalRequest"] = totalRequest
	ctx.Data["attackRequest"] = attackRequest
	ctx.Data["per"]=fmt.Sprintf("%.1f",float32(attackRequest)/float32(totalRequest)*100)
	ctx.HTML(200, "dashboard")
}
