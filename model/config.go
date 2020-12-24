package model

type WafConfig struct {
	Id      int64
	Type    string `xorm:"notnull varchar(10)"`
	Name    string `xorm:"notnull"`
	Value   string `xorm:"notnull text"`
	NameKey string `xorm:"notnull"`
	Button  bool   `xorm:"notnull"`
}

var wafConfigInit = []*WafConfig{
	{
		Type:    "waf",
		Name:    "waf状态",
		Value:   "on",
		NameKey: "config_waf_status",
		Button:  true,
	}, {
		Type:    "waf",
		Name:    "waf规则目录",
		Value:   "/opt/openresty/nginx/conf/mywaf/rules",
		NameKey: "config_rule_dir",
		Button:  false,
	}, {
		Type:    "waf",
		Name:    "waf日志记录目录",
		Value:   "/tmp",
		NameKey: "config_log_dir",
		Button:  false,
	}, {
		Type:    "waf",
		Name:    "waf拦截模式(redirect/html)",
		Value:   "html",
		NameKey: "config_waf_mode",
		Button:  false,
	}, {
		Type:    "waf",
		Name:    "waf拦截后跳转网页",
		Value:   "http://github.com/medasz",
		NameKey: "config_redirect_uri",
		Button:  false,
	}, {
		Type: "waf",
		Name: "waf拦截返回页面",
		Value: `<html>
    <head>
    <meta charset="UTF-8">
    <title>MIDUN WAF</title>
    </head>
      <body>
        <div>
      <div class="table">
        <div>
          <div class="cell">
            您的IP为: %s
          </div>
          <div class="cell">
            欢迎在遵守白帽子道德准则的情况下进行安全测试。
          </div>
          <div class="cell">
            联系方式：http://xsec.io
          </div>
        </div>
      </div>
    </div>
      </body>
    </html>`,
		NameKey: "config_output_html",
		Button:  false,
	}, {
		Type:    "ip",
		Name:    "ip白名单检测开关",
		Value:   "on",
		NameKey: "config_white_ip",
		Button:  true,
	}, {
		Type:    "ip",
		Name:    "ip黑名单检测开关",
		Value:   "on",
		NameKey: "config_black_ip",
		Button:  true,
	}, {
		Type:    "head",
		Name:    "user_agent黑名单检测开关",
		Value:   "on",
		NameKey: "config_user_agent",
		Button:  true,
	}, {
		Type:    "head",
		Name:    "cookie黑名单检测开关",
		Value:   "on",
		NameKey: "config_black_cookie",
		Button:  true,
	}, {
		Type:    "uri",
		Name:    "uri白名单检测开关",
		Value:   "on",
		NameKey: "config_white_uri",
		Button:  true,
	}, {
		Type:    "uri",
		Name:    "uri黑名单检测开关",
		Value:   "on",
		NameKey: "config_black_uri",
		Button:  true,
	}, {
		Type:    "data",
		Name:    "get参数黑名单检测开关",
		Value:   "on",
		NameKey: "config_black_get_args",
		Button:  true,
	}, {
		Type:    "data",
		Name:    "post请求体黑名单检测开关",
		Value:   "on",
		NameKey: "config_black_post",
		Button:  true,
	}, {
		Type:    "cc",
		Name:    "CC攻击检测开关",
		Value:   "on",
		NameKey: "config_cc",
		Button:  true,
	}, {
		Type:    "cc",
		Name:    "CC攻击频率设置",
		Value:   "10/60",
		NameKey: "config_cc_rate",
		Button:  false,
	},
}

//初始化waf_config表
func initWafConfig() error {
	err := AddWafConfigMul(wafConfigInit)
	return err
}

//添加数据
func AddWafConfig(wafConfig *WafConfig) error {
	_, err := Engine.InsertOne(wafConfig)
	if err != nil {
		return err
	}
	return nil
}

func AddWafConfigMul(wafConfig []*WafConfig) error {
	_, err := Engine.Insert(wafConfig)
	if err != nil {
		return err
	}
	return nil
}

//获取waf_config数据
func GetWafConfig() ([]WafConfig, error) {
	wafConfig := []WafConfig{}
	err := Engine.Find(&wafConfig)
	if err != nil {
		return wafConfig, err
	}
	return wafConfig, nil
}

//更新数据根据id
func UpdateConfig(id int64, value string) (int, error) {
	n, err := Engine.ID(id).Cols("value").Update(WafConfig{Value: value})
	return int(n), err
}