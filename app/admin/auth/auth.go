package auth

import (
	response "bbs/library"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

type LoginReqEntity struct {
	Email    string `p:"email" v:"required|length:6,16#请输入正确的邮箱|邮箱长度应当在:min到:max之间"`
	Password string `p:"password" v:"required|length:6,16#请填写密码|密码长度应当在:min到:max之间"`
}

// POST|提交登录
func (c *Controller) Login(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, "admin/auth/login.html", g.Map{})
	}
	var data LoginReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	authEmail := g.Config().GetString("auth.email")
	authPassword := g.Config().GetString("auth.password")
	password, _ := gmd5.Encrypt(data.Password)
	if data.Email != authEmail || password != authPassword {
		response.RedirectBackWithError(r, gerror.New("账号密码错误"))
	}
	if err := r.Session.Set("admin", "admin"); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/admin/home", "登录成功")
	}
}

// POST|退出登录
func (c *Controller) Logout(r *ghttp.Request) {
	if err := r.Session.Remove("admin"); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/admin/login", "退出成功")
	}
}
