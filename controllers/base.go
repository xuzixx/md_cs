package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// LoginPreparer ...
// 这种interface 的可以有多个
// 只要继承baseController 的结构实现接口就会在Prepare 中执行
type LoginPreparer interface {
	LoginPrepare()
}

type baseController struct {
	beego.Controller
}

// Prepare implemented Prepare method for baseController
func (c *baseController) Prepare() {

	if app, ok := c.AppController.(LoginPreparer); ok {
		app.LoginPrepare()
	}

}

func (c *baseController) GetCurrentURI() string {
	return c.Ctx.Input.URI()
}

// 重定向
func (c *baseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *baseController) loginCheck() {
	user := c.GetSession("_login_user")
	if user == nil {
		redirect := c.GetCurrentURI()
		beego.Info("user has not login, redirect to login")
		if redirect != "" {
			c.redirect(fmt.Sprintf("/login?next=%s", redirect))
		}
		c.redirect("/login")
	}
}
