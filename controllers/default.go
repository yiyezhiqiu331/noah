package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index/index.html"
}


type WilController struct {
	beego.Controller
}

func (c *WilController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index/welcome.html"
}


type Menu1Controller struct {
	beego.Controller
}

func (c *Menu1Controller) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	beego.Info("menu11111111111")
	c.TplName = "index/menu1.html"
}


type PublicController struct {
	beego.Controller
}

func (c *PublicController) Get() {

	c.TplName = "index/public.html"
}


type CreateController struct {
	beego.Controller
}

func (c *CreateController) Get() {

	c.TplName = "index/create.html"
}



type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {


	c.TplName = "index/list.html"
}


type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {


	c.TplName = "index/test.html"
}


