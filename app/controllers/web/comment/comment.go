package comment

import (
	"bbs/app/service/model/comment"
	response "bbs/app/funcs/response"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// Controller Base
type Controller struct{}

func (c *Controller) Add(r *ghttp.Request) {
	var reqEntity comment.AddReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if err := comment.CommentPost(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, fmt.Sprintf("/posts/%d", reqEntity.Pid), "评论成功")
}
