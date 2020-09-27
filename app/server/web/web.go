package web

import (
	commentsModel "bbs/app/model/comments"
	"bbs/app/model/nodes"
	postsModel "bbs/app/model/posts"
	response "bbs/library"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout   string = "web/layout.html"
	homeTpl  string = "web/home.html"
	postsTpl string = "web/posts.html"
)

// Controller Base
type Controller struct{}

func (c *Controller) Home(r *ghttp.Request) {
	pid := r.GetQueryInt("pid")
	pageNum := r.GetQueryInt("page", 1)
	// 获取顶级顶级节点
	tops, _ := g.DB().Table(nodes.Table).Where("pid = ? and is_top = ?", 0, 1).Order("sort DESC").All()

	var children gdb.Result
	if pid == 0 {
		children, _ = g.DB().Table(nodes.Table).Where("pid != ? and is_top = ?", 0, 1).Order("sort DESC").All()
	} else {
		children, _ = g.DB().Table(nodes.Table).Where("pid = ?", pid).Order("sort DESC").All()
	}

	var posts gdb.Result
	// 获取动态数据
	posts, _ = g.DB().Table(postsModel.Table+" p").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Fields("p.id,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.created_at,u.name,u.avatar,n.name as node_name").
		Order("created_at DESC").
		Page(pageNum, 40).
		All()

	total, _ := g.DB().Table(postsModel.Table).Count()

	page := r.GetPage(total, 40)

	response.ViewExit(r, layout, g.Map{"tops": tops, "children": children, "pid": pid, "posts": posts, "mainTpl": homeTpl, "page": page.GetContent(2)})
}

// 帖子详情
func (c *Controller) PostDetail(r *ghttp.Request) {
	pageNum := r.GetQueryInt("page", 1)
	postsId := r.GetRouterVar("postsId").Int64()

	posts, _ := g.DB().Table(postsModel.Table+" p").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Fields("p.id,p.title,p.content,p.uid,p.nid,p.view_num,p.comment_num,p.created_at,u.name as user_name,u.avatar,n.name node_name").
		Where("p.id = ?", postsId).
		One()

	comments, _ := g.DB().Table(commentsModel.Table+" c").
		Fields("c.id,c.uid,c.ruid,c.content,u.name,u.avatar,ru.name as r_user_name,c.created_at").
		LeftJoin("users u", "u.id = c.uid").
		LeftJoin("users ru", "ru.id = c.ruid").
		Where("c.pid", postsId).
		Order("id ASC").
		Page(pageNum, 40).
		All()

	total, _ := g.DB().Table(commentsModel.Table).Where("pid", postsId).Count()

	page := r.GetPage(total, 40)

	data := g.Map{"mainTpl": postsTpl, "posts": posts, "comments": comments, "page": page.GetContent(2)}

	response.ViewExit(r, layout, data)
}
