package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	messageIndexTpl = "web/message/index.html"
)

// MessageController Base
type MessageController struct{}

// Index 消息中心
func (c *MessageController) Index(r *ghttp.Request) {
	uid := gconv.Int(r.Session.GetMap("user")["id"])

	total := service.MessageService.Total(uid)
	items := service.MessageService.List(uid, 0, 20)
	page := r.GetPage(total, 20)

	_ = service.MessageService.ReadAll(uid)

	data := g.Map{"mainTpl": messageIndexTpl, "items": items, "page": page.GetContent(2)}

	response.ViewExit(r, constants.WebLayoutTplPath, data)
}
