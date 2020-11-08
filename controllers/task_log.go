package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
)

type TaskLogController struct {
	beego.Controller
}

func (self *TaskLogController) List() {
	taskId, err := self.GetInt("task_id")
	if err != nil {
		taskId = 1
	}

	task, err := models.TaskGetById(taskId)
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.Data["pageTitle"] = "日志管理 - " + task.TaskName + "(#" + strconv.Itoa(task.Id) + ")"
	self.Data["task_id"] = task.Id
	self.display()
}