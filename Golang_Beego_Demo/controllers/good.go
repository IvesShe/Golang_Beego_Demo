package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "您好beego"
	c.Data["num"] = 12
	c.TplName = "goods.tpl"
}
