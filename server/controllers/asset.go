package controllers

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/forms"
	"github.com/imsilence/gocmdb/server/models"
)

type AssetPageController struct {
	LayoutController
}

func (c *AssetPageController) Index() {
	c.Data["menu"] = "asset"
	c.Data["expand"] = "asset_management"
}

type AssetController struct {
	auth.LoginRequiredController
}

func (c *AssetController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	qs := orm.NewOrm().QueryTable("asset")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total
	if query != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("name__icontains", query)
		queryCond = queryCond.Or("remark__icontains", query)
		cond = cond.AndCond(queryCond)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var assets []*models.Asset
	qs.SetCond(cond).Limit(length).Offset(start).All(&assets)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          assets,
	}
	c.ServeJSON()
}

func (c *AssetController) Create() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.PlatformCreateForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()

				platform := &models.Platform{
					Name:       form.Name,
					Type:       form.Type,
					Addr:       form.Addr,
					Region:     form.Region,
					Key:        form.Key,
					Secrect:    form.Secrect,
					Remark:     form.Remark,
					CreateUser: c.User.Id,
				}
				if _, err := ormer.Insert(platform); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": platform,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "cloud/create.html"
		c.Data["types"] = models.PlatformTypes
	}
}

func (c *AssetController) Detail() {
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	if pk, err := c.GetInt("pk"); err == nil {
		platform := &models.Platform{Id: pk}
		ormer := orm.NewOrm()
		if ormer.Read(platform) == nil {
			json = map[string]interface{}{
				"code":   200,
				"text":   "获取成功",
				"result": platform,
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *AssetController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.PlatformModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()

				platform := &models.Platform{Id: form.Id}
				if ormer.Read(platform) == nil {
					platform.Name = form.Name
					platform.Type = form.Type
					platform.Addr = form.Addr
					platform.Region = form.Region
					platform.Key = form.Key
					platform.Secrect = form.Secrect
					platform.Remark = form.Remark

					if _, err := ormer.Update(platform); err == nil {
						json = map[string]interface{}{
							"code":   200,
							"text":   "更新成功",
							"result": platform,
						}
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		platform := &models.Platform{Id: -1}
		if pk, err := c.GetInt("pk"); err == nil {
			platform.Id = pk
			orm.NewOrm().Read(platform)
		}
		c.TplName = "cloud/modify.html"
		c.Data["platform"] = platform
		c.Data["types"] = models.PlatformTypes
	}
}

func (c *AssetController) Delete() {
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
			platform := &models.Platform{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(platform) == nil {
				platform.Delete()
				if _, err := ormer.Update(platform, "delete_time"); err == nil {
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
