package main

import (
	"log"
	"time"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/imsilence/gocmdb/server/models"
	"github.com/imsilence/gocmdb/server/cloud"
)

func main() {
	conf, err := config.NewConfig("ini", "./conf/db.conf")
	if err != nil {
		log.Fatal(err)
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conf.String("dsn"))
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		log.Fatal("数据库连接错误")
	}

	platform := new(models.Platform)

	for now := range time.Tick(10 * time.Second) {
		for _, pf := range platform.AllEnabled() {
			go func(pf *models.Platform) {
				log.Println(pf.Name)
				ormer := orm.NewOrm()
				sdk, err := cloud.DefaultManager.Cloud(pf.Type)
				if err != nil {
					log.Println(err)
					return
				}

				sdk.Init(pf.Addr, pf.Key, pf.Secrect, pf.Region)
				if err := sdk.TestConnect(); err != nil {
					log.Println(err)
					return
				}

				for _, instance := range sdk.GetInstances() {
					obj := &models.VirtualMachine{UUID: instance.UUID, Platform: pf}
					if _, _, err := ormer.ReadOrCreate(obj, "UUID", "Platform"); err != nil {
						log.Println(err)
						continue
					}
					obj.Key = instance.Key
					obj.Name = instance.Name
					obj.OS = instance.OS
					obj.CPU = instance.CPU
					obj.Memory = instance.Memory
					obj.PublicAddrs = strings.Join(instance.PublicAddrs, ",")
					obj.PrivateAddrs = strings.Join(instance.PrivateAddrs, ",")
					obj.Status = instance.Status
					obj.VmCreatedTime = instance.CreatedTime
					obj.VmExpiredTime = instance.ExpiredTime
					ormer.Update(obj)
				}
				ormer.QueryTable(new(models.VirtualMachine)).Filter("platform__exact", pf).Filter("update_time__gte", now).Update(orm.Params{"delete_time": nil})
				ormer.QueryTable(new(models.VirtualMachine)).Filter("platform__exact", pf).Filter("update_time__lt", now).Update(orm.Params{"delete_time": now})
				pf.SyncTime = &now
				ormer.Update(pf, "sync_time")
			}(pf)
		}

	}
}