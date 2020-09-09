package controllers
import (
	"github.com/astaxie/beego"
	"myproject/models"

	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type SessionController struct {
	beego.Controller
}

func(this*SessionController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	logs.Info("22222222")
	log.Debug("%s",this.Data)
}

func (this*SessionController)GetSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	user := models.User{}
	//user.Name = "wyj"
//
	//resp["errno"] = 0
	//resp["errmsg"] = "OK"


	resp["errno"] =models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] =models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user

	}

}

func (this*SessionController)DeleteSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")

	resp["errno"] =models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)


}


func (this*SessionController)Login(){
//func (this*SessionController)Get(){

	//this.TplName = "index/login.html"

//1.得到用户信息
	resp := make(map[string]interface{})
	defer this.RetData(resp)
//     resp;
	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)

	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("%s",resp)
	beego.Info("199999999999")
	//println("======name = ",resp["mobile"],"=======password =",resp["password"])


	beego.Info("======name = ",resp["mobile"],"=======password =",resp["password"],"222222",resp["verity"])


//2.判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)

		beego.Info("111111name = ",resp["mobile"],"=======password =",resp["password"])
		return
	}


//3.与数据库匹配判断账号密码正确

	o := orm.NewOrm()
	user := models.User{Name:resp["mobile"].(string)}

	qs := o.QueryTable("user")
	//log.Info(string(qs.Count("mobile")))
	//err := qs.Filter("mobile", "18817595710").One(&user)
	err := qs.Filter("mobile", resp["mobile"]).One(&user)

	if err != nil{
		resp["errno"] = models.RECODE_ROLEERR
		resp["errmsg"] = models.RecodeText(models.RECODE_ROLEERR)
		beego.Info("222222=======errr =",err)
		return
	}

	if user.Password_hash != resp["password"]{
		resp["errno"] = models.RECODE_PWDERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
		beego.Info("333333name = ",resp["mobile"],"=======password =",resp["password"])
		return
	}




	//4.添加session
	//this.SetSession("name",resp["mobile"])
	//this.SetSession("mobile",resp["mobile"])
	//this.SetSession("user_id",user.Id)


    beego.Info("0000000000000000")
//5.返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}