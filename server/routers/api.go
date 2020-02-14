package routers

import (
	"github.com/astaxie/beego"
	"github.com/imsilence/gocmdb/server/controllers/api/v1"
	"github.com/imsilence/gocmdb/server/controllers/api/v2"
)

func init() {
	nsV1 := beego.NewNamespace("/v1",
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
		beego.NSRouter("api/task/:uuid/", &v1.APIController{}, "*:Task"),
	)
	nsV2 := beego.NewNamespace("/v2",
		beego.NSRouter("api/register/:uuid/", &v2.APIController{}, "*:Register"),
		beego.NSRouter("api/heartbeat/:uuid/", &v2.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/log/:uuid/", &v2.APIController{}, "*:Log"),
		beego.NSRouter("api/task/:uuid/", &v2.APIController{}, "*:Task"),
	)
	beego.AddNamespace(nsV1, nsV2)
}
