package web

import (
	"bbs/app/funcs/response"
	postsModel "bbs/app/model/posts"
	"bbs/app/model/users"
	"bbs/app/service"
	"fmt"
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
		response.ViewExit(r, webLayout, g.Map{"mainTpl": registerTpl})
	}
	var reqEntity service.RegisterReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if captcha := r.Session.Get("captcha"); captcha != reqEntity.Captcha {
		response.RedirectBackWithError(r, gerror.New("验证码错误"))
	}
	_ = r.Session.Remove("captcha")
	if err := service.UserService.Register(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/", "注册成功")
	}
}

// Login 用户登录
func (c *UserController) Login(r *ghttp.Request) {
	if r.Method == "GET" {
		data := g.Map{"mainTpl": loginTpl}
		response.ViewExit(r, webLayout, data)
	}
	var reqEntity service.LoginReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if record, err := service.UserService.Login(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		_ = r.Session.Set(UserSessionKey, record)
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
	err := r.Session.Remove(UserSessionKey)
	if err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.RedirectToWithMessage(r, "/", "退出成功")
	}
}

// Edit 编辑用户
func (c *UserController) Edit(r *ghttp.Request) {
	authUser := GetAuthUser(r)
	id := gconv.String(authUser["id"])
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
			response.ViewExit(r, webLayout, data)
		}
	} else {
		tab := r.PostFormValue("tab")
		// 编辑基础信息
		if tab == "info" {
			var reqEntity service.UpdateInfoEntity
			if err := r.Parse(&reqEntity); err != nil {
				response.RedirectBackWithError(r, err)
			}
			err := service.UserService.UpdateInfo(id, &reqEntity)
			if err != nil {
				response.RedirectBackWithError(r, err)
			}
		} else if tab == "avatar" {
			// 修改头像
			var reqEntity service.UpdateAvatarEntity
			if err := r.Parse(&reqEntity); err != nil {
				response.RedirectBackWithError(r, err)
			}
			err := service.UserService.UpdateAvatar(id, &reqEntity)
			if err != nil {
				response.RedirectBackWithError(r, err)
			}
		} else {
			// 修改密码
			var reqEntity service.UpdatePasswordEntity
			if err := r.Parse(&reqEntity); err != nil {
				response.RedirectBackWithError(r, err)
			}
			err := service.UserService.UpdatePassword(id, &reqEntity)
			if err != nil {
				response.RedirectBackWithError(r, err)
			}
		}
		response.RedirectToWithMessage(r, fmt.Sprintf("/user/edit?tab=%s", tab), "更新成功")
	}
}

// Center 用户中心
func (c *UserController) Center(r *ghttp.Request) {
	uid := gconv.Int(r.GetRouterValue("id"))
	page := r.GetQueryInt("page", 1)

	user := service.UserService.GetUserById(uid)

	posts, _ := g.DB().Table(postsModel.Table+" p").
		LeftJoin("users u", "u.id = p.uid").
		LeftJoin("nodes n", "n.id = p.nid").
		Fields("p.id,p.fine,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.create_at,u.name,u.avatar,n.name as node_name").
		Where("p.uid = ?", uid).
		Order("p.id DESC").
		Page(page, 20).
		All()

	total, _ := g.DB().Table(postsModel.Table).Where("uid = ", uid).Count()

	response.ViewExit(r, webLayout, g.Map{
		"user":    user,
		"posts":   posts,
		"mainTpl": centerTpl,
		"page":    r.GetPage(total, 20).GetContent(2),
	})
}
