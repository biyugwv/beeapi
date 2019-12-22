package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beeapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeapi/controllers:WsController"] = append(beego.GlobalControllerRouter["beeapi/controllers:WsController"],
        beego.ControllerComments{
            Method: "Join",
            Router: `/join`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeapi/controllers:WsController"] = append(beego.GlobalControllerRouter["beeapi/controllers:WsController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
