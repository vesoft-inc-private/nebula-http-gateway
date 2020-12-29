package routers

import (
	"nebula-http-gateway/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/db/connect", &controllers.DatabaseController{}, "POST:Connect")
	beego.Router("/api/db/exec", &controllers.DatabaseController{}, "POST:Execute")
	beego.Router("/api/db/disconnect", &controllers.DatabaseController{}, "POST:Disconnect")
}
