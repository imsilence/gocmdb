package controllers

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type AgentPageController struct {
	LayoutController
}

func (c *AgentPageController) Index() {
	c.Data["menu"] = "agent"
	c.Data["expand"] = "asset_management"
}

type AgentController struct {
	auth.LoginRequiredController
}

func (c *AgentController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	qs := orm.NewOrm().QueryTable("agent")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total
	if query != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("hostname__icontains", query)
		queryCond = queryCond.Or("ip__icontains", query)
		queryCond = queryCond.Or("os__icontains", query)
		queryCond = queryCond.Or("arch__icontains", query)
		queryCond = queryCond.Or("remark__icontains", query)
		cond = cond.AndCond(queryCond)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var agents []*models.Agent
	qs.SetCond(cond).Limit(length).Offset(start).All(&agents)
	for _, agent := range agents {
		agent.Patch()
	}

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          agents,
	}
	c.ServeJSON()
}

func (c *AgentController) Detail() {
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	if pk, err := c.GetInt("pk"); err == nil {
		agent := &models.Agent{Id: pk}
		ormer := orm.NewOrm()
		if ormer.Read(agent) == nil {
			agent.Patch()
			json = map[string]interface{}{
				"code":   200,
				"text":   "获取成功",
				"result": agent,
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *AgentController) Delete() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			agent := &models.Agent{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(agent) == nil {
				agent.Delete()
				if _, err := ormer.Update(agent, "delete_time"); err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "删除成功",
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}
