package models

import (
	"github.com/astaxie/beego/orm"
)

type Event struct {
	Model
	Id      int    `orm:"column(id);" json:"id"`
	Agent   string `orm:"column(agent);size(64);" json:"agent"`
	Message string `orm:"column(message);size(1024);" json:"message"`
	Remark  string `orm:"column(remark);size(1024);" json:"remark"`
	Status  string `orm:"column(status);default(0);" json:"status"`
}

func init() {
	orm.RegisterModel(new(Event))
}
