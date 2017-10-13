package controllers

import (
	"github.com/astaxie/beego"
	"restapp/models"
)

type ProxyController struct {
	beego.Controller
}

func (p *ProxyController) GetUserInfo() {
	proxy := &models.Proxy{}
	err, users := proxy.GetAllUser()
	if err != nil {
		p.Data["error"] = err
	} else {
		p.Data["data"] = users
	}
	p.ServeJSON()
}
