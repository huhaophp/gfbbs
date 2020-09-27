package user

import (
	"bbs/app/service/model/user"
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout      = "web/layout.html"
	loginTpl    = "web/user/login.html"
	detailTpl   = "web/user/detail.html"
	registerTpl = "web/user/register.html"
)

// Controller Base
type Controller struct{}

// Register 用户注册
func (c *Controller) Register(r *ghttp.Request) {
	if r.Method == "GET" {
		data := g.Map{"mainTpl": registerTpl}
		response.ViewExit(r, layout, data)
	}
	var reqEntity user.RegisterReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := user.Register(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/", "注册成功")
	}
}

// Login 用户登录
func (c *Controller) Login(r *ghttp.Request) {
	if r.Method == "GET" {
		data := g.Map{"mainTpl": loginTpl}
		response.ViewExit(r, layout, data)
	}
	var reqEntity user.LoginReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if record, err := user.Login(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		_ = r.Session.Set("user", record)
		response.RedirectToWithMessage(r, "/", "登录成功")
	}
}

// Logout 用户退出
func (c *Controller) Logout(r *ghttp.Request) {
	_ = r.Session.Remove("user")
	response.RedirectToWithMessage(r, "/", "退出成功")
}

// Detail 用户详情
func (c *Controller) Detail(r *ghttp.Request) {
	if r.Method == "GET" {
		data := g.Map{"mainTpl": detailTpl}
		response.ViewExit(r, layout, data)
	}
}
