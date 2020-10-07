package web

import (
	response "bbs/app/funcs/response"
	"bbs/app/service/model/comment"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// Controller Base
type CommentController struct{}

func (c *CommentController) Add(r *ghttp.Request) {
	var reqEntity comment.AddReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := comment.CommentPost(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, fmt.Sprintf("/posts/%d", reqEntity.Pid), "评论成功")
}
