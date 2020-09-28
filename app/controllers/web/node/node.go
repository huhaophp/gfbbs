package node

import (
	"bbs/app/model/nodes"
	postsModel "bbs/app/model/posts"
	response "bbs/library"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout   string = "web/layout.html"
	indexTpl string = "web/node/index.html"
)

// Controller Base
type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	id := r.GetRouterVar("nodeId").Int()
	pageNum := r.GetQueryInt("page", 1)
	if id == 0 {
		response.RedirectBackWithError(r, gerror.New("节点未找到"))
	}

	item, err := g.DB().Table(nodes.Table).Where("id = ? and is_delete = ?", id, 0).One()
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if item.IsEmpty() {
		response.RedirectBackWithError(r, gerror.New("节点未找到"))
	}

	items, _ := g.DB().Table(postsModel.Table+" p").
		Fields("p.id,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.created_at,u.name,u.avatar,n.name as node_name").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Where("p.nid = ?", id).
		Order("created_at DESC").
		Page(pageNum, 40).
		All()

	response.ViewExit(r, layout, g.Map{"mainTpl": indexTpl, "node": item, "posts": items})
}
