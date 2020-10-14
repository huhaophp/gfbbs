package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// MessageController Base
type MessageController struct{}

// Index 消息中心
func (c *MessageController) Index(r *ghttp.Request) {
	uid := gconv.Int(r.Session.GetMap("user")["id"])
	items := service.MessageService.List(uid, 0, 20)
	_ = service.MessageService.ReadAll(uid)
	data := g.Map{"mainTpl": "web/message/index.html", "items": items}
	response.ViewExit(r, constants.WebLayoutTplPath, data)
}
