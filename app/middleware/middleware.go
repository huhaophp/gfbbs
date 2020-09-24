package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 检查登录
func AdminAuthCheck(r *ghttp.Request) {
	auth := r.Session.Get("admin", "")
	if auth == "" || auth == nil {
		if r.IsAjaxRequest() || r.Header.Get("Accept") == "application/json" {
			r.Response.WriteJson(g.Map{"error": "登录已过期"})
			r.Exit()
		} else {
			r.Session.Set("error", "登录已过期")
			r.Response.RedirectTo("/admin/login")
		}
	} else {
		r.Middleware.Next()
	}
}
