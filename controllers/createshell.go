package controllers

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"github.com/astaxie/beego"
	"myproject/models"
	"strconv"
	"fmt"
)

func(this*CreateController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	logs.Info("22222222")
	log.Debug("%s",this.Data)
}




func (this*CreateController) Post() {
	//func (this*SessionController)Get(){

	//this.TplName = "index/login.html"

	//1.得到用户信息
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	//     resp;
	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)



	beego.Info(this.GetString("command"))
	beego.Info(resp["command"])
	beego.Info(resp["parameter"])

	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("%s", resp)
	beego.Info("199999999999")
	task := new(models.Task)
	//task.Command := resp["command"]
	task.Command = resp["command"].(string)
	//beego.Info(task.Command)
	task.TaskName = resp["title"].(string)
	//task.Task_type = int(resp["type"].(int))
	//task.Task_team = int(resp["team"].(int))

	Task_type,err := strconv.Atoi(resp["type"].(string))
	if err == nil {
		beego.Info(fmt.Printf("i: %v\n",Task_type))
	}
	task.Task_type = Task_type

	Task_team,err := strconv.Atoi(resp["type"].(string))
	if err == nil {
		beego.Info(fmt.Printf("i: %v\n",Task_team))
	}
	task.Task_team = Task_team

	task.Task_parameter = resp["parameter"].(string)
	beego.Info(task.Task_parameter)
	beego.Info(task.Command)
	beego.Info(task.Task_type)

	if _, err := models.TaskAdd(task); err != nil {
		//this.ajaxMsg(err.Error(), -1)
		beego.Info("nihao")
	}







	//println("======name = ",resp["mobile"],"=======password =",resp["password"])

	//beego.Info("======name = ", resp["mobile"], "=======password =", resp["password"], "222222", resp["verity"])

	////2.判断是否合法
	//if resp["mobile"] == nil || resp["password"] == nil {
	//	resp["errno"] = models.RECODE_DATAERR
	//	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	//
	//	beego.Info("111111name = ", resp["mobile"], "=======password =", resp["password"])
	//	return
	//}
	//
	////3.与数据库匹配判断账号密码正确
	//
	//o := orm.NewOrm()
	//user := models.User{Name: resp["mobile"].(string)}
	//
	//qs := o.QueryTable("user")
	////log.Info(string(qs.Count("mobile")))
	////err := qs.Filter("mobile", "18817595710").One(&user)
	//err := qs.Filter("mobile", resp["mobile"]).One(&user)
	//
	//if err != nil {
	//	resp["errno"] = models.RECODE_ROLEERR
	//	resp["errmsg"] = models.RecodeText(models.RECODE_ROLEERR)
	//	beego.Info("222222=======errr =", err)
	//	return
	//}
	//
	//if user.Password_hash != resp["password"] {
	//	resp["errno"] = models.RECODE_PWDERR
	//	resp["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
	//	beego.Info("333333name = ", resp["mobile"], "=======password =", resp["password"])
	//	return
	//}

	//4.添加session
	//this.SetSession("name",resp["mobile"])
	//this.SetSession("mobile",resp["mobile"])
	//this.SetSession("user_id",user.Id)

	beego.Info("0000000000000000")
	//5.返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}