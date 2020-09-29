package middleware

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/admins"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// AdminAuthCheck Check admin authorization check.
func AdminAuthCheck(r *ghttp.Request) {
	auth := r.Session.GetMap(constants.AdminSessionKey)
	if auth == nil {
		if r.IsAjaxRequest() || r.Header.Get("Accept") == "application/json" {
			response.Json(r, 401, "Authorization failed")
		} else {
			response.RedirectToWithError(r, "/admin/login", gerror.New("登录已过期"))
		}
	} else {
		admin, err := g.DB().Table(admins.Table).WherePri(auth["id"]).Where("status", admins.NormalStatus).One()
		if err != nil || admin == nil {
			//r.Session.Remove(constants.AdminSessionKey)
			response.RedirectToWithError(r, "/admin/login", gerror.New("登录失效"))
		}
		r.Middleware.Next()
	}
}

// WebAuthCheck Check web authorization check.
func WebAuthCheck(r *ghttp.Request) {
	auth := r.Session.Get(constants.UserSessionKey)
	if auth == nil {
		if r.IsAjaxRequest() || r.Header.Get("Accept") == "application/json" {
			response.Json(r, 401, "Authorization failed")
		} else {
			response.RedirectToWithError(r, "/user/login", gerror.New("登录已过期"))
		}
	} else {
		r.Middleware.Next()
	}
}
