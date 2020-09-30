package user

import (
	"bbs/app/funcs/response"
	"bbs/app/model/users"
	"bbs/app/request/User"
	"bbs/app/service/admin/user"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/user/list.html"
	createTpl string = "admin/user/create.html"
	editTpl   string = "admin/user/edit.html"
	errorTpl  string = "admin/error.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	items, err := g.DB().Table(users.Table).All()
	if err != nil {
		response.ViewExit(r, layout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl, "items": items})
	}
}

func (c *Controller) Add(r *ghttp.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": createTpl})
	}
	var data User.AddReqEntity
	if err := User.AddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err := user.Add(&data)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, "/admin/users", "添加成功")
}

func (c *Controller) Edit(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	item, err := g.DB().Table(users.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("用户未找到"))
	}
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": editTpl, "user": item})
	}
	var data User.UpdateReqEntity
	if err := User.UpdateReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err = user.Edit(&data, id)
	if err != nil {
		g.Log().Error("编辑失败:", err)
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	}
	response.RedirectToWithMessage(r, "/admin/users", "编辑成功")
}

func (c *Controller) Delete(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	err := user.Delete(id)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("删除失败"))
	}
	response.RedirectToWithMessage(r, "/admin/users", "删除成功")
}
