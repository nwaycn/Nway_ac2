package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Application"] = "上海宁卫FreeSwitch呼叫平台,基于FreeSWITCH的大并发，强劲能力提供更优性能和功能的呼叫平台"
	c.Data["Title"]="上海宁卫呼叫平台"
	c.Data["Website"] = "freeswitch.net.cn"
	c.Data["Email"] = "lihao@nway.com.cn"
	c.TplNames = "login.tpl"
}
