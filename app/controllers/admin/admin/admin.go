package admin

import (
	"bbs/app/model/admins"
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/admin/list.html"
	createTpl string = "admin/admin/create.html"
	editTpl   string = "admin/admin/edit.html"
	errorTpl  string = "admin/error.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	items, err := g.DB().Table(admins.Table).All()
	if err != nil {
		response.ViewExit(r, layout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl, "items": items})
	}
}
