// @APIVersion 1.0.0
// @Title MD_CS API
// @Description some test
// @Contact wxyahoo@gmail.com
package routers

import (
	"github.com/astaxie/beego"

	c "github.com/xuzixx/md_cs/controllers"
)

func init() {
	beego.Router("/", &c.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/books", beego.NSInclude(&c.BookController{})),
		beego.NSNamespace("/test", beego.NSInclude(&c.TestController{})),
	)

	beego.AddNamespace(ns)
}
