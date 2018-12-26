package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (c *AdminController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /admin/?:key([a-zA-Z]+) [*]
func (c *AdminController) Index() {
	c.Ctx.WriteString("ceshi" + c.Ctx.Input.Param(":key"))
}

// @router /staticblock/:key [get]
func (c *AdminController) StaticBlock() {
	c.Ctx.WriteString("staticblock" + c.Ctx.Input.Param(":key"))
}

// @router /all/:key [get]
func (c *AdminController) AllBlock()  {
	c.Ctx.WriteString("all" + c.Ctx.Input.Param(":key"))
}

