package web

import (
	"bbs/app/funcs/response"
	"bbs/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	messageTpl = "web/message/index.html"
)

// MessageController Base
type MessageController struct{}

// Index 消息中心
func (c *MessageController) Index(r *ghttp.Request) {
	authUser := GetAuthUser(r)
	uid := gconv.Int(authUser["id"])

	total := service.MessageService.Total(uid)
	items := service.MessageService.List(uid, 0, 20)
	page := r.GetPage(total, 20)

	// 阅读所有未读消息
	_ = service.MessageService.ReadAll(uid)

	data := g.Map{"mainTpl": messageTpl, "items": items, "page": page.GetContent(2)}

	response.ViewExit(r, webLayout, data)
}
