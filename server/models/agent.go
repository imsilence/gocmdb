package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

type Agent struct {
	Model
	Id            int        `orm:"column(id);" json:"id"`
	UUID          string     `orm:"column(uuid);size(64);" json:"uuid"`
	Hostname      string     `orm:"column(hostname);size(64);" json:"hostname"`
	IP            string     `orm:"column(ip);size(4096);" json:"ip"`
	OS            string     `orm:"column(os);size(128);" json:"os"`
	Arch          string     `orm:"column(arch);size(128);" json:"arch"`
	Ram           int64      `orm:"column(ram);" json:"ram"`
	CPU           int        `orm:"column(cpu);" json:"cpu"`
	Disk          string     `orm:"column(disk);size(4096);" json:"disk"`
	Remark        string     `orm:"column(remark);size(1024);" json:"remark"`
	CreateUser    int        `orm:"column(create_user);default(0);" json:"create_user"`
	HeartbeatTime *time.Time `orm:"column(heartbeat_time);null;" json:"heartbeat_time"`

	IsOnline bool           `orm:"-" json:"is_online"`
	IPList   []string       `orm:"-" json:"ip_list"`
	DiskList map[string]int `orm:"-" json:"disk_list" `
}

func (a *Agent) Patch() {
	a.IsOnline = false
	if a.HeartbeatTime != nil && time.Since(*a.HeartbeatTime) < 5*time.Minute {
		a.IsOnline = true
	}
	if a.IP != "" {
		json.Unmarshal([]byte(a.IP), &a.IPList)
	}
	if a.Disk != "" {
		json.Unmarshal([]byte(a.Disk), &a.DiskList)
	}
}

func init() {
	orm.RegisterModel(new(Agent))
}
