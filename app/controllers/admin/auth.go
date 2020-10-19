package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/request/Auth"
	"bbs/app/service/admin/auth"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	homePage  = "/admin/home"
	LoginPage = "/admin/login"
	loginTpl  = "admin/auth/login.html"
)

type AuthController struct{}

// Login 登录页面
func (c *AuthController) Login(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, loginTpl, g.Map{})
	}
	var data Auth.LoginReqEntity
	err := Auth.LoginReqCheck(r, &data)
	if err != nil {
		response.RedirectBackWithError(r, gerror.New("请输入登录账号密码"))
	}
	res, err := auth.Login(&data)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := r.Session.Set(constants.AdminSessionKey, res); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, homePage, "登录成功")
	}
}

// Logout 退出登录
func (c *AuthController) Logout(r *ghttp.Request) {
	if err := r.Session.Remove(constants.AdminSessionKey); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, LoginPage, "退出成功")
	}
}
