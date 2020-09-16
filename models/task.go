package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)


type Task struct {
	Id            int
	GroupId       int
	ServerIds     string
	ServerType    int
	TaskName      string
	Description   string
	CronSpec      string
	Concurrent    int
	Command       string
	Timeout       int
	ExecuteTimes  int
	PrevTime      int64
	Status        int
	IsNotify      int
	NotifyType    int
	NotifyTplId   int
	NotifyUserIds string
	CreateId      int
	UpdateId      int
	CreateTime    int64
	UpdateTime    int64
	Task_type     int
	Task_team     int
	Task_parameter  string
}
//func (t *Task) TableName() string {
//	return TableName("task")
//}


func TaskAdd(task *Task) (int64, error) {
	if task.TaskName == "" {
		return 0, fmt.Errorf("任务名称不能为空")
	}

	//if task.CronSpec == "" {
	//	return 0, fmt.Errorf("时间表达式不能为空")
	//}
	if task.Command == "" {
		return 0, fmt.Errorf("命令内容不能为空")
	}
	if task.CreateTime == 0 {
		task.CreateTime = time.Now().Unix()
	}
	return orm.NewOrm().Insert(task)
}




type List struct {
	Id          int
	CreateId    int
	UpdateId    int
	GroupName   string
	Description string
	CreateTime  int64
	UpdateTime  int64
	Status      int
}

func GetList(page, pageSize int) ([]*Task, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Task, 0)
	query := orm.NewOrm().QueryTable("Task")

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}



func TaskGetById(id int) (*Task, error) {
	task := &Task{
		Id: id,
	}

	err := orm.NewOrm().Read(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}



