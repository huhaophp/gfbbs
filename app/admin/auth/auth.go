package auth

import (
	"bbs/app/model/admins"
	"bbs/app/request/admin"
	response "bbs/library"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	sessionKey = "admin"
	homePage   = "/admin/home"
	LoginPage  = "/admin/login"
	loginTpl   = "admin/auth/login.html"
)

type Controller struct{}

// 登录页面
func (c *Controller) Login(r *ghttp.Request) {
	isAuth := r.Session.Get(sessionKey)
	if isAuth != nil {
		response.RedirectToWithMessage(r, homePage, "")
	}
	if r.Method == "GET" {
		response.ViewExit(r, loginTpl, g.Map{})
	}
	var data admin.LoginReqEntity
	err := admin.LoginReqCheck(r, &data)
	if err != nil {
		response.RedirectBackWithError(r, gerror.New("请输入登录账号密码"))
	}
	res, err := g.DB().Table(admins.Table).Where("email = ?", data.Email).One()
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if res == nil {
		response.RedirectBackWithError(r, gerror.New("账号或密码错误"))
	}
	hash, _ := gmd5.Encrypt(data.Password)
	if hash != (res["password"].String()) {
		response.RedirectBackWithError(r, gerror.New("账号或者密码错误"))
	}
	if err := r.Session.Set(sessionKey, res["name"].String()); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, homePage, "登录成功")
	}
}

// 退出登录
func (c *Controller) Logout(r *ghttp.Request) {
	if err := r.Session.Remove(sessionKey); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, LoginPage, "退出成功")
	}
}
