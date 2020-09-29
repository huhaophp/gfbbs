package user

import (
	"bbs/app/constants"
	response "bbs/app/funcs/response"
	"bbs/app/model/users"
	"bbs/app/service/model/user"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout      = "web/layout.html"
	loginTpl    = "web/user/login.html"
	detailTpl   = "web/user/detail.html"
	EditTpl     = "web/user/edit.html"
	centerTpl   = "web/user/center.html"
	registerTpl = "web/user/register.html"
	errorTpl    = "web/error.html"
)

// Controller Base
type Controller struct{}

// Register 用户注册
func (c *Controller) Register(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": registerTpl})
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
		_ = r.Session.Set(constants.UserSessionKey, record)
		callbackUrl := r.GetQueryString("callback_url")
		if callbackUrl != "" {
			response.RedirectToWithMessage(r, callbackUrl, "登录成功")
		} else {
			response.RedirectToWithMessage(r, "/", "登录成功")
		}
	}
}

// Logout 用户退出
func (c *Controller) Logout(r *ghttp.Request) {
	err := r.Session.Remove(constants.UserSessionKey)
	if err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/", "退出成功")
	}
}

// Edit 编辑用户
func (c *Controller) Edit(r *ghttp.Request) {
	record, err := g.DB().Table(users.Table).WherePri(r.Session.GetMap(constants.UserSessionKey)["id"]).One()
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if record.IsEmpty() {
		response.RedirectBackWithError(r, gerror.New("用户不存在"))
	} else {
		response.ViewExit(r, layout, g.Map{"user": record, "mainTpl": EditTpl})
	}
}
