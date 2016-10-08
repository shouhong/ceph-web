package routers

import (
	"github.com/astaxie/beego"
	"github.com/tobegit3hub/ceph-web/controllers"
	"github.com/tobegit3hub/ceph-web/controllers_tenx"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tenx", &controllers_tenx.TenxController{})
}
