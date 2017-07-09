package controllers

// TestController ...
type TestController struct {
	baseController
}

// TestLogin ...
// @router /login [get]
func (c *TestController) TestLogin() {
	c.Ctx.WriteString("TestLogin")
}

// LoginPrepare ...
func (c *TestController) LoginPrepare() {
	c.loginCheck()
}
