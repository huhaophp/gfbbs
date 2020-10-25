package middleware

import (
	"bbs/app/constants"
	"bbs/app/controllers/web"
	"bbs/app/funcs/response"
	"bbs/app/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// AdminAuthCheck Check admin authorization check.
func AdminAuthCheck(r *ghttp.Request) {
	auth := r.Session.GetMap(constants.AdminSessionKey)
	if auth == nil {
		if r.IsAjaxRequest() || r.Header.Get("Accept") == "application/json" {
			response.JsonExit(r, 401, "登录已过期")
		} else {
			response.RedirectToWithError(r, "/admin/login", gerror.New("登录已过期"))
		}
	}
	// 检查管理员状态
	admin := service.AdminService.CheckAdminStatus(gconv.Int(auth["id"]))
	if admin == nil {
		_ = r.Session.Remove(constants.AdminSessionKey)
		response.RedirectToWithError(r, "/admin/login", gerror.New("登录失效"))
	}
	r.Middleware.Next()
}

// WebAuthCheck Check web authorization check.
func WebAuthCheck(r *ghttp.Request) {
	auth := web.GetAuthUser(r)
	if auth == nil {
		if r.IsAjaxRequest() || r.Header.Get("Accept") == "application/json" {
			response.JsonExit(r, 401, "登录已过期")
		} else {
			response.RedirectToWithError(r, "/user/login", gerror.New("登录已过期"))
		}
	}
	// 检查用户状态
	user := service.UserService.CheckUserStatus(gconv.Int(auth["id"]))
	if user == nil {
		_ = r.Session.Remove(web.UserSessionKey)
		response.RedirectToWithError(r, "/admin/login", gerror.New("登录失效"))
	}
	r.Middleware.Next()
}

// WebLayoutGlobalVariablesSetting
func WebLayoutGlobalVariablesSetting(r *ghttp.Request) {
	user := web.GetAuthUser(r)
	if user != nil {
		r.GetView().Assigns(g.Map{
			"unread_num": service.MessageService.GetUnreadNum(gconv.Int(user["id"])),
		})
	}
	r.Middleware.Next()
}
