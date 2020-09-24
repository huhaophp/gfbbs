package home

import (
	response "bbs/library"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

// GET|后台首页
func (c *Controller) Home(r *ghttp.Request) {
	response.ViewExit(r, "admin/layout.html", g.Map{"mainTpl": "admin/home.html"})
}
