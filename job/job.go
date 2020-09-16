
package job

import (
	"fmt"
	"strings"
	"strconv"
	"myproject/models"
	"github.com/astaxie/beego"
	"os/exec"
	//"PPGo_Job/libs"
	"time"
	"bytes"
	"runtime"
	"PPGo_Job/libs"
)



func RunCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		beego.Warn(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			beego.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}

func NewJobFromTask(task *models.Task) ([]*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}

	if task.ServerIds == "" {
		return nil, fmt.Errorf("任务执行失败，找不到执行的服务器")
	}

	TaskServerIdsArr := strings.Split(task.ServerIds, ",")
	jobArr := make([]*Job, 0)
	for _, server_id := range TaskServerIdsArr {
		if server_id == "0" {
			//本地执行
			job := NewCommandJob(task.Id, 0, task.TaskName, task.Command)
			job.Task = task
			job.Concurrent = false
			if task.Concurrent == 1 {
				job.Concurrent = true
			}
			//job.Concurrent = task.Concurrent == 1
			job.ServerId = 0
			job.ServerName = "本地服务器"
			jobArr = append(jobArr, job)
		} else {
			server_id_int, _ := strconv.Atoi(server_id)
			//远程执行
			server, _ := models.TaskServerGetById(server_id_int)

			if server.Status == 2 {
				fmt.Println("服务器已禁用")
				continue
			}

			if server.ConnectionType == 0 {
				if server.Type == 0 {
					//密码验证登录服务器
					job := RemoteCommandJobByPassword(task.Id, server_id_int, task.TaskName, task.Command, server)
					job.Task = task
					job.Concurrent = false
					if task.Concurrent == 1 {
						job.Concurrent = true
					}
					//job.Concurrent = task.Concurrent == 1
					job.ServerId = server_id_int
					job.ServerName = server.ServerName
					jobArr = append(jobArr, job)
				} else {
					job := RemoteCommandJob(task.Id, server_id_int, task.TaskName, task.Command, server)
					job.Task = task
					job.Concurrent = false
					if task.Concurrent == 1 {
						job.Concurrent = true
					}
					//job.Concurrent = task.Concurrent == 1
					job.ServerId = server_id_int
					job.ServerName = server.ServerName
					jobArr = append(jobArr, job)
				}
			} else if server.ConnectionType == 1 {
				if server.Type == 0 {
					//密码验证登录服务器
					job := RemoteCommandJobByTelnetPassword(task.Id, server_id_int, task.TaskName, task.Command, server)
					job.Task = task
					job.Concurrent = false
					if task.Concurrent == 1 {
						job.Concurrent = true
					}
					//job.Concurrent = task.Concurrent == 1
					job.ServerId = server_id_int
					job.ServerName = server.ServerName
					jobArr = append(jobArr, job)
				}
			} else if server.ConnectionType == 2 {
				//密码验证登录服务器
				job := RemoteCommandJobByAgentPassword(task.Id, server_id_int, task.TaskName, task.Command, server)
				job.Task = task
				job.Concurrent = false
				if task.Concurrent == 1 {
					job.Concurrent = true
				}
				//job.Concurrent = task.Concurrent == 1
				job.ServerId = server_id_int
				job.ServerName = server.ServerName
				jobArr = append(jobArr, job)

			}
		}
	}

	return jobArr, nil
}




func NewCommandJob(id int, serverId int, name string, command string) *Job {
	job := &Job{
		Id:   id,
		Name: name,
	}

	job.JobKey = libs.JobKey(id, serverId)
	job.RunFunc = func(timeout time.Duration) (jobresult *JobResult) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		//cmd := exec.Command("/bin/bash", "-c", command)
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("CMD", "/C", command)
			beego.Info("+++++++++++++++++")
			beego.Info(command)
		} else {
			cmd = exec.Command("sh", "-c", command)
			beego.Info("+++++++++++++++++")
			beego.Info(command)
		}
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		cmd.Start()
		err, isTimeout := runCmdWithTimeout(cmd, timeout)
		jobresult = new(JobResult)
		jobresult.OutMsg = bufOut.String()
		jobresult.ErrMsg = bufErr.String()

		jobresult.IsOk = true
		if err != nil {
			jobresult.IsOk = false
		}

		jobresult.IsTimeout = isTimeout

		return jobresult
	}
	return job
}