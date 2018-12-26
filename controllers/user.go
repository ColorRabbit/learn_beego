package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	userId, errConv := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if errConv != nil {
		c.Ctx.WriteString("ID 不存在")

		return
	}

	o := orm.NewOrm()
	user := models.User{Id: userId}

	err := o.Read(&user)
	if err == orm.ErrNoRows {
		c.Ctx.WriteString("查询不到")
	} else if err == orm.ErrMissPK {
		c.Ctx.WriteString("找不到主键")
	} else {
		fmt.Println(user.Id)
		c.Ctx.WriteString(user.Name)
	}
}

func (c *UserController) Post() {
	name := c.Input().Get("name")
	age, errConv := strconv.Atoi(c.Input().Get("age"))
	if errConv != nil {
		age = 0
	}
	o := orm.NewOrm()

	profile := new(models.Profile)
	profile.Age = age
	o.Insert(profile)

	user := new(models.User)
	user.Profile = profile
	user.Name = name
	o.Insert(user)

	c.Ctx.WriteString("添加成功")
}
