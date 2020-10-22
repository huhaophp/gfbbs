package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/comments"
	adminService "bbs/app/service/admin"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type CommentController struct{}

func (c *CommentController) List(r *ghttp.Request) {
	pageNum := r.GetQueryInt("page", 1)
	postId := r.GetRouterVar("post_id").Int()

	items, err := g.DB().Table(comments.Table).
		Where("is_delete = ?", 0).
		Where("pid = ?", postId).
		Page(pageNum, 20).
		All()
	total, _ := g.DB().Table(comments.Table).Where("is_delete = ?", 0).Count()
	page := r.GetPage(total, 20)

	if err != nil {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": constants.AdminErrorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
			"mainTpl": fmt.Sprintf(constants.AdminListTpl, "comment"),
			"items":   items,
			"page":    page.GetContent(2),
		})
	}
}

func (c *CommentController) Del(r *ghttp.Request) {
	postId := r.GetRouterVar("post_id").Int()
	commentId := r.GetRouterVar("comment_id").Int()
	if postId <= 0 || commentId <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	item, err := g.DB().Table(comments.Table).Where("id = ? and is_delete = ?", postId, 0).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("评论不存在或已被删除"))
	}
	err = adminService.Comment.Delete(postId)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("删除失败"))
	}
	response.RedirectToWithMessage(r, fmt.Sprintf("/admin/posts/%d", postId), "删除成功")
}
