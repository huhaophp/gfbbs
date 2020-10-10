package web

import (
	response "bbs/app/funcs/response"
	"bbs/app/model/nodes"
	postsModel "bbs/app/model/posts"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout      string = "web/layout.html"
	homeTpl     string = "web/home/home.html"
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
		Fields("p.id,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.create_at,u.name,u.avatar,n.name as node_name").
		Order("create_at DESC").
		Page(pageNum, 20).
		All()

	total, _ := g.DB().Table(postsModel.Table).Count()

	page := r.GetPage(total, 20)

	data := g.Map{"tops": tops, "children": children, "pid": pid, "posts": posts, "mainTpl": homeTpl, "page": page.GetContent(4)}

	response.ViewExit(r, layout, data)
}
