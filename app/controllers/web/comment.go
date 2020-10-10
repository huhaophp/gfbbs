package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/service/web/comment"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// CommentController
type CommentController struct{}

// Add Comment post
func (c *CommentController) Add(r *ghttp.Request) {
	var reqEntity comment.AddCommentReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := comment.Add(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.BackWithMessage(r, "评论成功")
}

// Del Delete comment
func (c *CommentController) Del(r *ghttp.Request) {
	id := r.GetRouterString("id")
	uid := gconv.String(r.Session.GetMap(constants.UserSessionKey)["id"])
	err := comment.CheckPermissions(id, uid)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := comment.Delete(id); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.BackWithMessage(r, "删除成功")
	}
}
