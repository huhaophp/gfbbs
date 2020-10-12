package web

import (
	response "bbs/app/funcs/response"
	postsModel "bbs/app/model/posts"
	"bbs/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout  string = "web/layout.html"
	homeTpl string = "web/home/home.html"
)

// Controller Base
type Controller struct{}

func (c *Controller) Home(r *ghttp.Request) {
	nid := r.GetQueryInt("nid")
	pageNum := r.GetQueryInt("page", 1)

	nodes := service.NodeService.Get(g.Map{"status": 0})

	// 获取动态数据
	posts, _ := g.DB().Table(postsModel.Table+" p").
		LeftJoin("users u", "u.id = p.uid").
		LeftJoin("nodes n", "n.id = p.nid").
		LeftJoin("users u1", "u1.id = p.luid").
		Fields("p.luid,p.id,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.create_at,u.name,u.avatar,n.name as node_name,u1.name as last_user_name").
		Order("create_at DESC").
		Page(pageNum, 20).
		All()
	g.Dump(posts)
	total, _ := g.DB().Table(postsModel.Table).Count()

	page := r.GetPage(total, 20)

	data := g.Map{"nodes": nodes, "nid": nid, "posts": posts, "mainTpl": homeTpl, "page": page.GetContent(4)}

	response.ViewExit(r, layout, data)
}
