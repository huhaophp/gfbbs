package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// CommentController
type CommentController struct{}

// Add Comment post
func (c *CommentController) Add(r *ghttp.Request) {
	var reqEntity service.AddCommentReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if reqEntity.Ruid == reqEntity.Uid {
		response.RedirectBackWithError(r, gerror.New("不能自己回复自己哦"))
	}
	if err := service.CommentService.Add(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.BackWithMessage(r, "评论成功")
}

// Del Delete comment
func (c *CommentController) Del(r *ghttp.Request) {
	id := r.GetRouterString("id")
	uid := gconv.String(r.Session.GetMap(constants.UserSessionKey)["id"])
	err := service.CommentService.CheckPermissions(id, uid)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := service.CommentService.Delete(id); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.BackWithMessage(r, "删除成功")
	}
}
