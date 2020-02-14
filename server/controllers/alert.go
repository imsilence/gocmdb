package controllers

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type AlertPageController struct {
	LayoutController
}

func (c *AlertPageController) Index() {
	c.Data["expand"] = "log_management"
	c.Data["menu"] = "alert_event"
}

type AlertController struct {
	auth.LoginRequiredController
}

func (c *AlertController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	status, err := c.GetInt("status")
	if err != nil {
		status = -1
	}

	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	qs := orm.NewOrm().QueryTable("event")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total
	if query != "" || status >= 0 {
		if query != "" {
			queryCond := orm.NewCondition()
			queryCond = queryCond.Or("message__icontains", query)
			queryCond = queryCond.Or("remark__icontains", query)
			cond = cond.AndCond(queryCond)
		}
		if status >= 0 {
			cond = cond.And("status__exact", status)
		}
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var events []*models.Event
	qs.SetCond(cond).Limit(length).Offset(start).All(&events)
	for _, event := range events {
		event.Patch()
	}

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          events,
	}
	c.ServeJSON()
}

func (c *AlertController) Detail() {
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	if pk, err := c.GetInt("pk"); err == nil {
		event := &models.Event{Id: pk}
		ormer := orm.NewOrm()
		if ormer.Read(event) == nil {
			event.Patch()
			json = map[string]interface{}{
				"code":   200,
				"text":   "获取成功",
				"result": event,
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

type AlertSettingsPageController struct {
	LayoutController
}

func (c *AlertSettingsPageController) Index() {
	c.Data["expand"] = "system_management"
	c.Data["menu"] = "alert_settings"
}

type AlertSettingsController struct {
	auth.LoginRequiredController
}
