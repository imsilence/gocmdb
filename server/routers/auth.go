package routers

import (
	"github.com/astaxie/beego"
	"github.com/imsilence/gocmdb/server/controllers/auth"
)

func init() {
	beego.AutoRouter(&auth.AuthController{})
}
