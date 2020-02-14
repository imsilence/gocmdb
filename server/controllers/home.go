package controllers

import (
	"net/http"
	"github.com/astaxie/beego"
	"github.com/imsilence/gocmdb/server/controllers/auth"
)

type HomeController struct {
	auth.LoginRequiredController
}

func (c *HomeController) Index() {
	c.Redirect(beego.URLFor("DashboardPageController.Index"), http.StatusFound)
}