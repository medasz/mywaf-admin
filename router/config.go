package router

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
	"strconv"
)

func GetWafConfig(ctx *macaron.Context) {
	wafConfig, err := model.GetWafConfig()
	if err != nil {
		log.Println(err)
		ErrorMsg(ctx, "Fail to get waf_config!")
	}
	tmp := map[string][]model.WafConfig{}
	for _, v := range wafConfig {
		tmp[v.Type] = append(tmp[v.Type], v)
	}
	ctx.Data["wafConfig"] = tmp
	ctx.Data["rulesType"] = model.RuleInfo
	ctx.HTML(200, "config")
}

type updateForm struct {
	Id    string
	Value string
}
type updateFormIns struct {
	Id    int64
	Value string
}

func UpdateWafConfig(ctx *macaron.Context)string{
	data, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Println(err)
		//ErrorMsg(ctx, "提交数据出错!")
		//return
		return "提交数据出错!"
	}

	updateFormData := []updateForm{}
	err = json.Unmarshal(data, &updateFormData)
	if err != nil {
		println(string(data))
		log.Println(err)
		//ErrorMsg(ctx, "提交数据出错!")
		//return
		return "提交数据出错!"
	}
	updateFormDataTmp := []updateFormIns{}
	for _, v := range updateFormData {
		id, err := strconv.Atoi(v.Id)
		if err != nil {
			log.Println(err)
			continue
		}
		tmp := updateFormIns{
			Id:    int64(id),
			Value: v.Value,
		}
		updateFormDataTmp = append(updateFormDataTmp, tmp)
	}
	flag := false
	for _, v := range updateFormDataTmp {
		n, err := model.UpdateConfig(v.Id, v.Value)
		if err != nil {
			log.Println(err)
			continue
		}
		if n > 0 {
			flag = true
		}
	}
	if flag {
		err := model.UpdataFlag("config")
		if err != nil {
			log.Println(err)
			//ErrorMsg(ctx, "更新uuid出错!")
			//return
			return "更新uuid出错!"
		}
	}

	//ErrorMsg(ctx,"修改成功!")
	return "修改成功!"
}

func GetWafConfigJson(ctx *macaron.Context) {
	resData := map[string]string{"result": "true"}
	wafConfig, err := model.GetWafConfig()
	if err != nil {
		log.Println(err)
		resData["result"] = "false"
		ctx.JSON(200, resData)
	}

	for _, v := range wafConfig {
		resData[v.NameKey] = v.Value
	}
	resData["uuid"], err = model.GetFlag("config")
	if err != nil {
		log.Println(err)
		resData["result"] = "false"
		ctx.JSON(200, resData)
	}
	ctx.JSON(200, resData)
}
