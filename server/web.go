package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/imsilence/gocmdb/server/utils"
	"github.com/imsilence/gocmdb/server/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/imsilence/gocmdb/server/routers"
)

func main() {
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	init := flag.Bool("init", false, "init server")
	syncdb := flag.Bool("syncdb", false, "sync db")
	force := flag.Bool("force", false, "force sync db(drop table)")
	verbose := flag.Bool("verbose", false, "verbose")

	flag.Usage = func() {
		fmt.Println(`usage: server -h`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		log.Fatal("数据库连接错误")
	}

	if *verbose {
		orm.Debug = true
	}

	switch {
	case *init:
		orm.RunSyncdb("default", *force, *verbose)
		ormer := orm.NewOrm()
		password := utils.RandString(6)
		admin := &models.User{Name: "admin", IsSuperuser: true}
		admin.SetPassword(password)
		if ormer.Read(admin, "Name") == orm.ErrNoRows {
			if _, err := ormer.Insert(admin); err == nil {
				fmt.Printf("初始化admin账号成功, 密码: %s\n", password)
			}
		} else {
			fmt.Println("admin账号已存在")
		}
	case *syncdb:
		orm.RunSyncdb("default", *force, *verbose)
	default:
		beego.Run()
	}
}
