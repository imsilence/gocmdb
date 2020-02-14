package models

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/cloud"
)

type Asset struct {
	Model
	Id           int    `orm:"column(id);" json:"id"`
	Name         string `orm:"column(name);size(64);" json:"name"`
	IP           string `orm:"column(ip);size(256);" json:"ip"`
	Type         string `orm:"column(type);size(32);" json:"type"`
	Application  string `orm:"column(application);size(256);" json:"application"`
	Manager      int    `orm:"column(manager);" json:"manager"`
	Remark       string `orm:"column(remark);size(1024);" json:"remark"`
	Agent        string `orm:"column(agent);size(64);" json:"uuid"`
	Instance     string `orm:"column(instance);size(64);" json:"instance"`
	InstanceDesc string `orm:"column(instance_desc); type(text);" json:"instance_desc"`
	CreateUser   int    `orm:"column(create_user);default(0);" json:"create_user"`
}

func (a *Asset) CreateOrReplace(instances []*cloud.Instance) {
	ormer := orm.NewOrm()
	for _, instance := range instances {
		obj := &Asset{Instance: instance.UUID}
		if _, _, err := ormer.ReadOrCreate(obj, "Instance"); err != nil {
			continue
		}
		obj.Name = instance.Name
		ip := ""
		if len(instance.PublicAddrs) > 0 {
			ip = instance.PublicAddrs[0]
		} else if len(instance.PrivateAddrs) > 0 {
			ip = instance.PrivateAddrs[0]
		}
		obj.IP = ip
		obj.Type = "虚拟机"
		if desc, err := json.Marshal(instance); err != nil {
			obj.InstanceDesc = string(desc)
		}
		ormer.Update(obj)
	}
}

func init() {
	orm.RegisterModel(new(Asset))
}
