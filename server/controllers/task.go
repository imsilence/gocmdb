package controllers

import "github.com/imsilence/gocmdb/server/controllers/auth"

type TaskPageController struct {
	LayoutController
}

func (c *TaskPageController) Index() {
	c.Data["menu"] = "task_management"
}

type TaskController struct {
	auth.LoginRequiredController
}
