package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"] = append(beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"] = append(beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"] = append(beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:storeId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"] = append(beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:storeId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"] = append(beego.GlobalControllerRouter["manageFile/controllers:StoreHouseController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manageFile/controllers:UserController"] = append(beego.GlobalControllerRouter["manageFile/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
