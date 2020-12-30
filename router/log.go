package router

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
	"time"
)

func SaveLog(ctx *macaron.Context) {
	res, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Printf("Fail to get post body:%s", err)
		return
	}
	wafLog := model.WafLog{}
	err = json.Unmarshal(res, &wafLog)
	if err != nil {
		log.Printf("Fail to parse post body:%s", err)
		return
	}
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", wafLog.LocalTime, time.Local)

	if err != nil {
		log.Printf("Fail to parse localtime:%s", err)
		return
	}
	wafLog.LocalTimeObj = the_time
	err = model.InsertLog(wafLog)
	if err != nil {
		log.Printf("Fail to insert log:%s", err)
	}
}

type JsonCountList struct {
	All    [25]int `json:"all"`
	Attack [25]int `json:"attack"`
}

func GetLog(ctx *macaron.Context) {
	result := map[string]interface{}{"result": "true"}
	data := [25][]model.WafLog{}
	allLogToday, err := model.GetAllLogToday()
	if err != nil {
		log.Println(err)
		result["result"] = "false"
		ctx.JSON(200, result)
		return
	}

	for _, y := range allLogToday {
		data[y.LocalTimeObj.Hour()] = append(data[y.LocalTimeObj.Hour()], y)
	}
	jsonCountList := JsonCountList{}
	for k, v := range data {
		jsonCountList.All[k] = len(v)
		for _, y := range v {
			if y.AttackType != "general" {
				jsonCountList.Attack[k] += 1
			}
		}
	}
	result["data"] = data
	result["jsonCountList"] = jsonCountList
	ctx.JSON(200, result)
}
