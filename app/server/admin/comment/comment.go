package comment

import (
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/comment/list.html"
)

// Controller Base
type Controller struct{}

// Category List
func (c *Controller) List(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl})
	}
}
