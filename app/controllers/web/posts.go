package web

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	commentsModel "bbs/app/model/comments"
	postsModel "bbs/app/model/posts"
	"bbs/app/service"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	postsTpl   string = "web/posts/detail.html"
	publishTpl string = "web/posts/publish.html"
)

// PostsController Base
type PostsController struct{}

// Details Post details
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
		Where("c.is_delete", 0).
		Order("id ASC").
		Page(pageNum, 20).
		All()

	total, _ := g.DB().Table(commentsModel.Table).Where("pid", postsId).Count()

	page := r.GetPage(total, 20)

	data := g.Map{"mainTpl": postsTpl, "posts": posts, "comments": comments, "page": page.GetContent(2)}

	response.ViewExit(r, constants.WebLayoutTplPath, data)
}

// Publish Post a post
func (c *PostsController) Publish(r *ghttp.Request) {
	if r.Method == "GET" {
		data := g.Map{"mainTpl": publishTpl}
		response.ViewExit(r, constants.WebLayoutTplPath, data)
	}
	var reqEntity service.PublishPostsReqEntity
	if err := r.Parse(&reqEntity); err != nil {
		response.RedirectBackWithError(r, err)
	}
	publisher := gconv.Int(r.Session.GetMap("user")["id"])
	if id := service.PostsService.Publish(publisher, &reqEntity); id == 0 {
		response.RedirectBackWithError(r, gerror.New("发布失败"))
	} else {
		response.RedirectToWithMessage(r, fmt.Sprintf("/posts/%d", id), "发布成功")
	}
}
