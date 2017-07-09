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

	beego.Include(&c.BookController{})
	beego.Include(&c.TestController{})
}
