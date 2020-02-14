package models

import "github.com/astaxie/beego/orm"

type Task struct {
	Id     int    `orm:"column(id);" json:"id"`
	Agent  string `orm:"column(agent);size(64);" json:"agent"`
	Status int    `orm:"column(status);default(0);" json:"status"`
}

func GetTasksForAgent(agent string) []*Task {
	var tasks []*Task
	ormer := orm.NewOrm()
	ormer.QueryTable("task").Filter("agent__exact", agent).Filter("status__in", []int{statusNew}).All(&tasks)
	return tasks
}

func init() {
	orm.RegisterModel(new(Task))
}
