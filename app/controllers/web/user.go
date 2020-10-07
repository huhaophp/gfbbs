package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/users"
	"bbs/app/service/model/user"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	loginTpl    = "web/user/login.html"
	EditTpl     = "web/user/edit.html"
	centerTpl   = "web/user/center.html"
	registerTpl = "web/user/register.html"
)

// UserController Base
type UserController struct{}

// Register 用户注册
func (c *UserController) Register(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, constants.WebLayoutTplPath, g.Map{"mainTpl": registerTpl})
	}
	var reqEntity user.RegisterReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if captcha := r.Session.Get("captcha"); captcha != reqEntity.Captcha {
		response.RedirectBackWithError(r, gerror.New("验证码错误"))
	}
	_ = r.Session.Remove("captcha")
	if err := user.Register(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/", "注册成功")
	}
}

// Login 用户登录
func (c *UserController) Login(r *ghttp.Request) {
	if r.Method == "GET" {
		data := g.Map{"mainTpl": loginTpl}
		response.ViewExit(r, constants.WebLayoutTplPath, data)
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
func (c *UserController) Logout(r *ghttp.Request) {
	err := r.Session.Remove(constants.UserSessionKey)
	if err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/", "退出成功")
	}
}

// Edit 编辑用户
func (c *UserController) Edit(r *ghttp.Request) {
	if r.Method == "GET" {
		// 默认选中 info 信息
		tab := r.GetQueryString("tab", "info")
		id := r.Session.GetMap(constants.UserSessionKey)["id"]
		record, err := g.DB().Table(users.Table).WherePri(id).One()
		if err != nil {
			response.RedirectBackWithError(r, err)
		}
		if record.IsEmpty() {
			response.RedirectBackWithError(r, gerror.New("用户不存在"))
		} else {
			response.ViewExit(r, constants.WebLayoutTplPath, g.Map{"user": record, "mainTpl": EditTpl, "tab": tab})
		}
	} else {

	}
}

// Center 用户中心
func (c *UserController) Center(r *ghttp.Request) {
	g.Dump(r.GetRouterVar("id"))
	response.ViewExit(r, constants.WebLayoutTplPath, g.Map{"mainTpl": centerTpl})
}
