package controllers

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type LogPageController struct {
	LayoutController
}

func (c *LogPageController) Index() {
	c.Data["expand"] = "log_management"
	c.Data["menu"] = "log_analysis"
}

type LogController struct {
	auth.LoginRequiredController
}

func (c *LogController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	agent := strings.TrimSpace(c.GetString("agent"))

	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	qs := orm.NewOrm().QueryTable("log")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total
	if query != "" {
		cond = cond.And("message__icontains", query)
		qtotal, _ = qs.SetCond(cond).Count()
	}
	if agent != "" {
		cond = cond.And("agent__exact", agent)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var logs []*models.Log
	qs.SetCond(cond).Limit(length).Offset(start).All(&logs)
	for _, log := range logs {
		log.Patch()
	}

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          logs,
	}
	c.ServeJSON()
}
