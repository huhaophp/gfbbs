package cate

import (
	response "bbs/app/funcs/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/cate/list.html"
	createTpl string = "admin/cate/create.html"
	editTpl   string = "admin/cate/edit.html"
	errorTpl  string = "admin/error.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl})
	}
}
