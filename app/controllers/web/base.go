package web

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	UserSessionKey string = "user"
	webLayout      string = "web/layout.html"
)

// 获取当前登录的用户
func GetAuthUser(r *ghttp.Request) g.Map {
	return r.Session.GetMap(UserSessionKey)
}
