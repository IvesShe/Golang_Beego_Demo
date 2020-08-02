package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用戶中心")
}

func (c *UserController) AddUser() {
	c.TplName = "user.html"
}

func (c *UserController) DoAddUser() {

	username := c.GetString("username")
	password := c.GetString("password")
	c.Ctx.WriteString("用戶中心--" + username + "  " + password)
}
