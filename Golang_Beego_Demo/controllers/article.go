package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Ctx.WriteString("新聞列表")
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新聞")
}

func (c *ArticleController) EditArticle() {
	// http://localhost:8080/article/edit?id=123
	// 獲取Get傳值
	// id := c.GetString("id")
	// beego.Info(id)
	// fmt.Printf("值:%v 類型:%T\n", id, id)
	// c.Ctx.WriteString("修改新聞--" + id)
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("傳入參數錯誤")
		return
	}
	fmt.Printf("值:%v 類型:%T\n", id, id)
	c.Ctx.WriteString("修改新聞")
}
