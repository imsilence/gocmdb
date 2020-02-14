package models

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"
)

type Log struct {
	Model
	Id      int    `orm:"column(id);" json:"id"`
	Agent   string `orm:"column(agent);size(64);" json:"agent"`
	Type    int    `orm:"column(type);" json:"type"`
	Message string `orm:"column(message);size(4096);" json:"message"`

	AgentObject   *Agent                 `orm:"-" json:"agent_object"`
	MessageObject map[string]interface{} `orm:"-" json:"message_object"`
}

func (l *Log) Patch() {
	if l.Agent != "" {
		ormer := orm.NewOrm()
		a := &Agent{UUID: l.Agent}
		if err := ormer.Read(a, "UUID"); err == nil {
			a.Patch()
			l.AgentObject = a
		}
	}
	if l.Message != "" {
		json.Unmarshal([]byte(l.Message), &l.MessageObject)
	}
}

func init() {
	orm.RegisterModel(new(Log))
}
