package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mywaf-admin/setting"
	"mywaf-admin/utils"
)

type Rule struct {
	Id       int64
	RuleItem string `xorm:"not null"`
	RuleType string `xorm:"not null"`
}

var RuleInfo =[]string{
	"black_cookie",
	"black_file_ext",
	"black_get_args",
	"black_ip",
	"black_post",
	"black_uri",
	"black_user_agent",
	"white_ip",
	"white_uri",
}

//初始化数据库rule表
func initRule() error {
	fileList, err := utils.FileList(setting.RulePath)
	if err != nil {
		return err
	}
	rulesList := ParseRulesFile(fileList)
	for _, rule := range rulesList {
		err := NewRule(rule.RuleItem, rule.RuleType)
		if err != nil {
			log.Printf("Fail to insert Rule{RuleItem:%v,RuleType:%v}", rule.RuleItem, rule.RuleType)
			continue
		}
	}
	return nil
}

//添加规则
func NewRule(ruleItem, ruleType string) error {
	_, err := Engine.Insert(&Rule{RuleItem: ruleItem, RuleType: ruleType})
	return err
}

//删除规则
func DeleteRule(id int64) error {
	_,err:=Engine.Delete(&Rule{Id: id})
	return err
}

//获取所有规则
func GetAllRule() ([]Rule, error) {
	rules := []Rule{}
	err := Engine.Find(&rules)
	if err != nil {
		log.Println("Fail to find all rules :", err)
		return nil, err
	}
	return rules, nil
}

//根据ruleType获取rule
func GetRule(ruleType string)([]Rule,error){
	rules:=[]Rule{}
	err:=Engine.Table("rule").Where("rule_type=?",ruleType).Find(&rules)
	return rules,err
}

//根据id获取rule
func GetRuleById(id int64)(Rule,error){
	rule:=Rule{Id: id}
	_,err:=Engine.Get(&rule)
	return rule,err
}

//添加rule
func AddRule(rule,ruleType string)error{
	_,err:=Engine.Insert(&Rule{RuleItem: rule,RuleType: ruleType})
	return err
}

//根据id更新rule
func UpdateRuleById(id int64,rule string)error{
	_,err:=Engine.ID(id).Update(&Rule{RuleItem: rule})
	return err
}

//获取所有规则类型
func GetRuleType() []string {
	typeList := []string{}
	err := Engine.Table("rule").Distinct("rule_type").Find(&typeList)
	if err != nil {
		log.Printf("Fail to find rule distinct rule_type :", err)
		return nil
	}
	return typeList
}

//解析规则文件(复数)
func ParseRulesFile(filesPath []string) []Rule {
	rules := []Rule{}
	for _, filePath := range filesPath {
		tmp, err := ParseRuleFile(filePath)
		if err != nil {
			log.Printf("Fail to parse rule file %s :%s", filePath, err)
			continue
		}
		rules = append(rules, tmp...)
	}
	return rules
}

//解析规则文件
func ParseRuleFile(filePath string) ([]Rule, error) {
	rule := []Rule{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Fail to read file %s:%s", filePath, err)
		return nil, err
	}
	err = json.Unmarshal(content, &rule)
	if err != nil {
		log.Printf("Fail to parse file %s:%s", filePath, err)
		return nil, err
	}
	return rule, nil
}
