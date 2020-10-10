package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/users"
	"bbs/app/service/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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
	var reqEntity model.RegisterReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if captcha := r.Session.Get("captcha"); captcha != reqEntity.Captcha {
		response.RedirectBackWithError(r, gerror.New("验证码错误"))
	}
	_ = r.Session.Remove("captcha")
	if err := model.Register(&reqEntity); err != nil {
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
	var reqEntity model.LoginReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if record, err := model.Login(&reqEntity); err != nil {
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
	id := gconv.String(r.Session.GetMap(constants.UserSessionKey)["id"])
	if r.Method == "GET" {
		// 默认选中 info 信息
		tab := r.GetQueryString("tab", "info")
		record, err := g.DB().Table(users.Table).WherePri(id).One()
		if err != nil {
			response.RedirectBackWithError(r, err)
		}
		if record.IsEmpty() {
			response.RedirectBackWithError(r, gerror.New("用户不存在"))
		} else {
			data := g.Map{"user": record, "mainTpl": EditTpl, "tab": tab}
			response.ViewExit(r, constants.WebLayoutTplPath, data)
		}
	} else {
		tab := r.PostFormValue("tab")
		// 编辑基础信息
		if tab == "info" {
			var reqEntity model.UpdateInfoEntity
			if err := r.Parse(&reqEntity); err != nil {
				response.RedirectBackWithError(r, err)
			}
			err := model.UpdateInfo(id, &reqEntity)
			if err != nil {
				response.RedirectBackWithError(r, err)
			}
			response.RedirectToWithMessage(r, "/user/edit?tab=info", "更新成功")
		} else if tab == "avatar" {
			// 修改头像
			var reqEntity model.UpdateAvatarEntity
			if err := r.Parse(&reqEntity); err != nil {
				response.RedirectBackWithError(r, err)
			}
			err := model.UpdateAvatar(id, &reqEntity)
			if err != nil {
				response.RedirectBackWithError(r, err)
			}
			response.RedirectToWithMessage(r, "/user/edit?tab=avatar", "更新成功")
		} else {
			// 修改密码
			var reqEntity model.UpdatePasswordEntity
			if err := r.Parse(&reqEntity); err != nil {
				response.RedirectBackWithError(r, err)
			}
			err := model.UpdatePassword(id, &reqEntity)
			if err != nil {
				response.RedirectBackWithError(r, err)
			}
			response.RedirectToWithMessage(r, "/user/edit?tab=password", "更新成功")
		}
	}
}

// Center 用户中心
func (c *UserController) Center(r *ghttp.Request) {
	response.ViewExit(r, constants.WebLayoutTplPath, g.Map{"mainTpl": centerTpl})
}
