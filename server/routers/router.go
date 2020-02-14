package routers

import (
	"github.com/astaxie/beego"

	"github.com/imsilence/gocmdb/server/controllers"
)

func init() {

	// home
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	// Dashboard
	beego.AutoRouter(&controllers.DashboardPageController{})
	beego.AutoRouter(&controllers.DashboardController{})

	// 终端
	beego.AutoRouter(&controllers.AgentPageController{})
	beego.AutoRouter(&controllers.AgentController{})

	// 资产
	beego.AutoRouter(&controllers.AssetPageController{})
	beego.AutoRouter(&controllers.AssetController{})

	// 云平台
	beego.AutoRouter(&controllers.CloudPageController{})
	beego.AutoRouter(&controllers.CloudController{})

	// 虚拟机
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	beego.AutoRouter(&controllers.VirtualMachineController{})

	// 任务
	beego.AutoRouter(&controllers.TaskPageController{})
	beego.AutoRouter(&controllers.TaskController{})

	// 告警
	beego.AutoRouter(&controllers.AlertPageController{})
	beego.AutoRouter(&controllers.AlertController{})

	// 日志
	beego.AutoRouter(&controllers.LogPageController{})
	beego.AutoRouter(&controllers.LogController{})

	// 用户
	beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})

	// 告警设置
	beego.AutoRouter(&controllers.AlertSettingsPageController{})
	beego.AutoRouter(&controllers.AlertSettingsController{})
}
