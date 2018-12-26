package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello/controllers:AdminController"] = append(beego.GlobalControllerRouter["hello/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/admin/?:key([a-zA-Z]+)`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:AdminController"] = append(beego.GlobalControllerRouter["hello/controllers:AdminController"],
        beego.ControllerComments{
            Method: "AllBlock",
            Router: `/all/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:AdminController"] = append(beego.GlobalControllerRouter["hello/controllers:AdminController"],
        beego.ControllerComments{
            Method: "StaticBlock",
            Router: `/staticblock/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
