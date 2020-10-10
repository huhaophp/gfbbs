package web

import (
	"bbs/app/funcs/response"
	"bbs/app/service/model/comment"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// CommentController
type CommentController struct{}

// Add Comment post
func (c *CommentController) Add(r *ghttp.Request) {
	var reqEntity comment.AddReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := comment.Add(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, fmt.Sprintf("/posts/%d", reqEntity.Pid), "评论成功")
}

// Del Delete comment
func (c *CommentController) Del(r *ghttp.Request) {
	id := r.GetRouterString("id")
	if err := comment.Del(id); err != nil {
		response.RedirectBackWithError(r, err)
	}
}
