package web

import (
	"bbs/app/funcs/response"
	"bbs/app/service"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// Controller Base
type LikeController struct{}

// Do 点赞
func (c *LikeController) Do(r *ghttp.Request) {
	var reqEntity service.LikeReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	authUer := r.Session.GetMap("user")
	if reqEntity.Uid != gconv.Int(authUer["id"]) {
		response.JsonExit(r, 0, "无权限操作")
	}
	if err := service.LikeService.Do(&reqEntity); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		likers := service.LikeService.GetTheLatestLikes(reqEntity.Tid, reqEntity.TidType, 20)
		response.JsonExit(r, 1, "点赞成功", g.Map{"likes": RenderLikesHtml(likers)})
	}
}

// RenderLikesHtml
func RenderLikesHtml(likes gdb.Result) string {
	html := ""
	for _, v := range likes {
		html += fmt.Sprintf("<a href='/users/%d'><img src='%s' class='img-thumbnail like-user'></a>", v["id"].Int(), v["avatar"])
	}
	return html
}

// Undo 取消点赞
func (c *LikeController) Undo(r *ghttp.Request) {
	var reqEntity service.LikeReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	if reqEntity.Uid != gconv.Int(r.Session.GetMap("user")["id"]) {
		response.JsonExit(r, 0, "无权限操作")
	}
	if err := service.LikeService.Undo(&reqEntity); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		likers := service.LikeService.GetTheLatestLikes(reqEntity.Tid, reqEntity.TidType, 20)
		response.JsonExit(r, 1, "取消成功", g.Map{"likes": RenderLikesHtml(likers)})
	}
}
