package routers

import (
	"Golang_Beego_Demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/article", &controllers.ArticleController{}) // 訪問Get方法
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")
	beego.Router("/user", &controllers.UserController{}) // 訪問Get方法
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")
}
