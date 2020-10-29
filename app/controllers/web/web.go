package web

import (
	"bbs/app/funcs/response"
	"bbs/app/job"
	postsModel "bbs/app/model/posts"
	"bbs/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	homeTpl string = "web/home/home.html"
)

// Controller Base
type Controller struct{}

func (c *Controller) Home(r *ghttp.Request) {
	nid := r.GetQueryInt("nid")
	pageNum := r.GetQueryInt("page", 1)

	nodes := service.NodeService.Get(g.Map{"status": 0})

	// 获取动态数据
	query := g.DB().Table(postsModel.Table+" p").
		LeftJoin("users u", "u.id = p.uid").
		LeftJoin("nodes n", "n.id = p.nid").
		LeftJoin("users u1", "u1.id = p.luid")

	if nid != 0 {
		query = query.Where("p.nid", nid)
	}

	posts, _ := query.
		Fields("p.luid,p.fine,p.id,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.create_at,"+
			"u.name,u.avatar,n.name as node_name,u1.name as last_user_name").
		Order("p.fine DESC,p.id").
		Page(pageNum, 20).
		All()

	total, _ := g.DB().Table(postsModel.Table).Count()

	page := r.GetPage(total, 20)

	latestPosts := service.PostsService.GetTheLatestPosts(10)

	activeUsers := service.UserService.GetActiveUsers(job.ActiveUserJob.GetActiveUsers())

	data := g.Map{
		"nodes":       nodes,
		"nid":         nid,
		"posts":       posts,
		"mainTpl":     homeTpl,
		"latestPosts": latestPosts,
		"activeUsers": activeUsers,
		"page":        page.GetContent(2),
	}

	response.ViewExit(r, webLayout, data)
}
