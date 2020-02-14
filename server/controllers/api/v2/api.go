package v2

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/controllers/api"
	"github.com/imsilence/gocmdb/server/models"
)

type APIController struct {
	api.BaseController
}

func (c *APIController) Prepare() {
	c.BaseController.Prepare()
	if beego.AppConfig.String("token") != c.Ctx.Input.Header("Token") {
		c.Data["json"] = map[string]interface{}{
			"code": 403,
			"text": "token不正确",
		}
		c.ServeJSON()
		c.StopRun()
	}
}

func (c *APIController) Heartbeat() {
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	uuid := c.Ctx.Input.Param(":uuid")
	agent := &models.Agent{UUID: uuid}
	ormer := orm.NewOrm()
	if ormer.Read(agent, "UUID") == nil {
		now := time.Now()
		agent.HeartbeatTime = &now
		ormer.Update(agent, "HeartbeatTime")
		json = map[string]interface{}{
			"code": 200,
			"text": "心跳发送成功",
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *APIController) Register() {
	rtJson := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	body := new(models.Agent)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, body); err == nil {
		ormer := orm.NewOrm()
		agent := &models.Agent{UUID: body.UUID}
		if created, id, err := ormer.ReadOrCreate(agent, "UUID"); err == nil {
			agent.Hostname = body.Hostname
			agent.IP = body.IP
			agent.OS = body.OS
			agent.Arch = body.Arch
			agent.Ram = body.Ram
			agent.CPU = body.CPU
			agent.Disk = body.Disk
			agent.DeleteTime = nil
			ormer.Update(agent)
			rtJson = map[string]interface{}{
				"code": 200,
				"text": "注册成功",
				"result": map[string]interface{}{
					"created": created,
					"id":      id,
				},
			}
		}
	}

	c.Data["json"] = rtJson
	c.ServeJSON()
}

func (c *APIController) Log() {
	rtJson := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	log := new(models.Log)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, log); err == nil {
		ormer := orm.NewOrm()
		if _, err := ormer.Insert(log); err == nil {
			rtJson = map[string]interface{}{
				"code": 200,
				"text": "上传成功",
			}
		}
	}

	c.Data["json"] = rtJson
	c.ServeJSON()
}

func (c *APIController) Task() {
	uuid := c.Ctx.Input.Param(":uuid")
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "获取任务成功",
		"result": models.GetTasksForAgent(uuid),
	}
	c.ServeJSON()
}
