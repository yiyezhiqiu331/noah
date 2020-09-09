package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"myproject/models"
	"strings"
	"github.com/astaxie/beego/logs"
)

//type SessionControllern struct {
//	beego.Controller
//}

func(this*LoginController)RetDatan(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	logs.Info("22222222")
	log.Debug("%s",this.Data)
}

type RegisterController struct {
	beego.Controller
}

// 注册展示页面
func (this *RegisterController) ShowRegister() {
	this.TplName = "index/register.html"
}

// 注册获取数据页面
func (this *RegisterController) HandleRegister() {


	// 获取浏览器传递的值，并去除两边的空格

	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))

	beego.Info("账号:", Name, "密码:", Pwd)
	// beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if Name == "" || Pwd == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "index/register.html"
		this.Data["errmsg"] = "用户名或密码不能为空 ！"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象
	o := orm.NewOrm()
	//   获取插入对象
	user := models.User{}
	//   插入数值
	user.Name = Name
	user.Password_hash = Pwd
	user.Mobile = Pwd

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		this.TplName = "index/register.html"
		this.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
		beego.Info(err)
		return
	}
	// 测试返回视图
	this.Ctx.WriteString("注册成功！！！")
	// 实际情况注册成功返回到登录页面
	this.Redirect("login", 302)
}


type LoginController struct {
	beego.Controller
}

// 登录页面 get
func (this *LoginController) ShowLogin() {
	this.TplName = "index/login.html"
}

// 登录页面 post

func (this *LoginController) HandleLogin() {

	resp := make(map[string]interface{})
	defer this.RetDatan(resp)

	beego.Info("99999999999")
	// 拿到浏览器数据，并去除两边的空格
	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	beego.Info("账号:", Name, "密码:", Pwd)

	//数据处理
	if Name == "" || Pwd == "" {
		beego.Info("登录失败！！")
		this.TplName = "index/login.html"
		this.Data["errmsg"] = "登录失败！！！！！"
		return
	}
	// 查找数据库
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var user models.User
	// 查询
	user.Name = Name
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名登录失败！！！")
		this.TplName = "index/login.html"
		this.Data["errmsg"] = "用户名登录失败！！！！！"
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)

		return
	}
	// 判断密码是否一致
	if user.Password_hash != Pwd {
		beego.Info("密码登录失败！！！")
		this.TplName = "index/login.html"
		this.Data["errmsg"] = "密码登录失败！！"
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	// 测试返回视图
	//this.Ctx.WriteString("登录成功")
	// 实际情况注册成功返回到登录页面
	this.Redirect("/", 302)
}