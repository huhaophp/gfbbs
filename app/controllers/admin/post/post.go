package post

import (
	response "bbs/app/funcs/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/post/list.html"
	createTpl string = "admin/post/create.html"
	editTpl   string = "admin/post/edit.html"
	errorTpl  string = "admin/error.html"
)

// Controller Base
type Controller struct{}

// Category List
func (c *Controller) List(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl})
	}
}
