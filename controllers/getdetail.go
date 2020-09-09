package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

type GetDetailController struct {
	beego.Controller
}

func(this*GetDetailController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	logs.Info("22222222")
	log.Debug("%s",this.Data)
}




func (this*GetDetailController) Get() {
	//func (this*SessionController)Get(){

	//this.TplName = "index/login.html"

	//1.得到用户信息
	//resp := make(map[string]interface{})
	//defer this.RetData(resp)
	//
	//groupResult, n := models.GetList(1, 1000)
	//
	//fmt.Print("~~~~~~~~~~~~~~~~\n")
	//fmt.Print(groupResult)
	//fmt.Print("~~~~~~~~~~~~~~~~\n")
	//
	//out := make(map[string]interface{})
	//out["code"] = 0
	//out["msg"] = "成功"
	//out["count"] = n
	//out["data"] = groupResult
	//this.Data["json"] = out
	//this.ServeJSON()
	//this.StopRun()
	//beego.Info(groupResult)
	//this.ajaxList("成功", 0, n, list)
	this.TplName = "index/detail.html"


	beego.Info("0000000000000000")

}