package web

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	webLayout      string = "web/layout.html"
	UserSessionKey string = "user"
)

func GetAuthUser(r *ghttp.Request) g.Map {
	return r.Session.GetMap(UserSessionKey)
}
