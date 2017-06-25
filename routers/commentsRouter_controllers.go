package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:BookController"] = append(beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:BookController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/v1/books`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:BookController"] = append(beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:BookController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/v1/book`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:BookController"] = append(beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:BookController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/v1/book/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/xuzixx/md_cs/controllers:TestController"],
		beego.ControllerComments{
			Method: "TestLogin",
			Router: `/test/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
