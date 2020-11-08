package controllers

import (
	"github.com/astaxie/beego"
	"myproject/models"
	//"PPGo_Job/jobs"
	//"PPGo_Job/jobs"
	"myproject/job"

	"time"
	"bytes"
	"os/exec"
	"runtime"
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


type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {


	c.TplName = "index/detail.html"

	c.Data["pageTitle"] = "任务详细"

	id, _ := c.GetInt("id")
	beego.Info("111111111111")
	beego.Info(id)
	task, err := models.TaskGetById(id)
	beego.Info(task)
	c.Data["task"] = task
	if err != nil {
		out := make(map[string]interface{})
		out["status"] = 200
		out["message"] = "未发现id"
		c.Data["json"] = out
		//c.ajaxMsg(err.Error(), MSG_ERR)
	}

}




func (self *DetailController) Post() {

	beego.Info("AjaxRunAjaxRunvAjaxRunAjaxRunAjaxRunAjaxRunvvv")
	id, _ := self.GetInt("id")
	task, err := models.TaskGetById(id)
	if err != nil {
		out := make(map[string]interface{})

		beego.Info(out["status"])
		out["status"] = 0
		out["message"] = "未发现id"
		self.Data["json"] = out
	}
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	//cmd := exec.Command("/bin/bash", "-c", command)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("CMD", "/C", task.Command)
		beego.Info("+++++++++++++++++")
		beego.Info(task.Command)
	} else {
		cmd = exec.Command("sh", "-c", task.Command)
		beego.Info("+++++++++++++++++")
		beego.Info(task.Command)
	}
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	cmd.Start()
	job.RunCmdWithTimeout(cmd,  60000000 * time.Millisecond)
	beego.Info("-----------------\n")
	beego.Info(bufOut.String())


	out := make(map[string]interface{})

	out["status"] = 0
	beego.Info(out["status"])

	out["message"] = "success"
	self.Data["json"] = out
	//self.Data["json"] = bufOut.String()







	beego.Info(task.ServerIds,task.Description)


	//插入日志
	log := new(models.TaskLog)
	log.TaskId = id
	log.ServerId = task.ServerIds
	log.ServerName = task.TaskName
	log.Output = bufOut.String()
	log.Error = bufOut.String()
	log.ProcessTime = int(time.Now().Unix())
	log.CreateTime = int64(time.Now().Unix())
	models.TaskLogAdd(log)

	//jobresult = new(job.JobResult)
	//jobresult.OutMsg = bufOut.String()
	//jobresult.ErrMsg = bufErr.String()
	//
	//jobresult.IsOk = true
	//if err != nil {
	//	jobresult.IsOk = false
	//}
	//
	//jobresult.IsTimeout = isTimeout



	//if err != nil {
	//	self.ajaxMsg(err.Error(), MSG_ERR)
	//}
	//for _, job := range jobArr {
	//	job.Run()
	//}
	//
	//self.ajaxMsg("", MSG_OK)
	beego.Info("1111111111111")
	self.ServeJSON()
	self.StopRun()
}

