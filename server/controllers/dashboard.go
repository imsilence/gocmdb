package controllers

import "github.com/imsilence/gocmdb/server/controllers/auth"

type DashboardPageController struct {
	LayoutController
}

func (c *DashboardPageController) Index() {
	c.Data["menu"] = "dashboard"
}

type DashboardController struct {
	auth.LoginRequiredController
}
