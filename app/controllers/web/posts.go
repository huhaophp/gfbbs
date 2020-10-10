package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	commentsModel "bbs/app/model/comments"
	postsModel "bbs/app/model/posts"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	postsTpl string = "web/posts/detail.html"
)

// Controller Base
type PostsController struct{}

func (c *PostsController) Details(r *ghttp.Request) {
	pageNum := r.GetQueryInt("page", 1)
	postsId := r.GetRouterVar("postsId").Int64()

	posts, _ := g.DB().Table(postsModel.Table+" p").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Fields("p.id,p.title,p.content,p.uid,p.nid,p.view_num,p.comment_num,p.create_at,u.name as user_name,u.avatar,n.name node_name").
		Where("p.id = ?", postsId).
		One()

	comments, _ := g.DB().Table(commentsModel.Table+" c").
		Fields("c.id,c.uid,c.ruid,c.content,u.name,u.avatar,ru.name as r_user_name,c.create_at").
		LeftJoin("users u", "u.id = c.uid").
		LeftJoin("users ru", "ru.id = c.ruid").
		Where("c.pid", postsId).
		Order("id ASC").
		Page(pageNum, 20).
		All()

	total, _ := g.DB().Table(commentsModel.Table).Where("pid", postsId).Count()

	page := r.GetPage(total, 20)

	data := g.Map{"mainTpl": postsTpl, "posts": posts, "comments": comments, "page": page.GetContent(2)}

	response.ViewExit(r, constants.WebLayoutTplPath, data)
}
