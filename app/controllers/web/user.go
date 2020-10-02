package web

import (
	"bbs/app/funcs/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	detailTpl = "web/user/detail.html"
)

type UserController struct{}

func (c *UserController) Detail(r *ghttp.Request) {
	uid := r.GetRouterVar("uid").Int()
	data := g.Map{"mainTpl": detailTpl, "uid": uid}
	response.ViewExit(r, layout, data)
}
