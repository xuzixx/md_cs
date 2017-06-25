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
