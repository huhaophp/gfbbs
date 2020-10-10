package web

import (
	"bbs/app/funcs/response"
	"bbs/app/service/model"
	"github.com/gogf/gf/net/ghttp"
)

// CommentController
type CommentController struct{}

// Add Comment post
func (c *CommentController) Add(r *ghttp.Request) {
	var reqEntity model.AddCommentReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := model.CommentAdd(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.BackWithMessage(r, "评论成功")
}

// Del Delete comment
func (c *CommentController) Del(r *ghttp.Request) {
	id := r.GetRouterString("id")
	if err := model.CommentDelete(id); err != nil {
		response.RedirectBackWithError(r, err)
	} else {
		response.BackWithMessage(r, "删除成功")
	}
}
