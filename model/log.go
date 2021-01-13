package model

import (
	"time"
	"xorm.io/builder"
)

type WafLog struct {
	Rule         string    `json:"rule"`
	ClientIp     string    `json:"client_ip"`
	AttackType   string    `json:"attack_type"`
	Data         string    `json:"data"`
	ServerName   string    `json:"server_name"`
	UserAgent    string    `json:"user_agent"`
	ReqUri       string    `json:"req_uri"`
	LocalTime    string    `json:"local_time"`
	LocalTimeObj time.Time `json:"local_time_obj"`
}


func InsertLog(data WafLog) error {
	_, err := Engine.InsertOne(&data)
	return err
}

func GetAllLogCount() (int, error) {
	n, err := Engine.Count(&WafLog{})
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func GetAttackLogCount() (int, error) {
	num, err := Engine.Where(builder.NotIn("attack_type", "general")).Count(&WafLog{})
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func GetAllLogToday() (data []WafLog, err error) {
	curDay := time.Now().Format("2006-01-02")
	err = Engine.Where("local_time REGEXP ?", curDay).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
