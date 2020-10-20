package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/nodes"
	"bbs/app/model/posts"
	"bbs/app/request/admin"
	adminService "bbs/app/service/admin"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

type PostController struct{}

func (c *PostController) List(r *ghttp.Request) {
	pageNum := r.GetQueryInt("page", 1)
	title := r.GetQueryString("title")
	uname := r.GetQueryString("uname")

	query := g.DB().Table(posts.Table+" p").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Fields("p.id,p.title,p.uid,p.nid,p.view_num,p.like_num,p.comment_num,p.fine,p.create_at,u.name as user_name,u.avatar,n.name node_name")

	if title != "" {
		query.Where("title like ?", "%"+title+"%")
	}
	if uname != "" {
		query.Where("u.name like ?", "%"+uname+"%")
	}
	items, err := query.Order("id DESC").Page(pageNum, 20).All()

	total, _ := g.DB().Table(nodes.Table).Count()
	page := r.GetPage(total, 20)

	if err != nil {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": constants.AdminErrorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
			"mainTpl": fmt.Sprintf(constants.AdminListTpl, "post"),
			"items":   items,
			"page":    page.GetContent(2),
		})
	}
}

func (c *PostController) Add(r *ghttp.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		nodesData,_ := g.DB().Table(nodes.Table).Where("is_delete = ?", 0).All()
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
			"mainTpl": fmt.Sprintf(constants.AdminCreateTpl, "post"),
			"nodes": nodesData,
		})
	}
	var data admin.PostAddReqEntity
	if err := admin.PostAddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err := adminService.Post.Add(&data)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, "/admin/posts", "添加成功")
}

func (c *PostController) Edit(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	item, err := g.DB().Table(posts.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("文章不存在"))
	}
	if strings.ToUpper(r.Method) == "GET" {
		nodesData,_ := g.DB().Table(nodes.Table).Where("is_delete = ?", 0).All()
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
			"mainTpl": fmt.Sprintf(constants.AdminEditTpl, "post"),
			"post": item,
			"nodes": nodesData,
		})
	}
	var data admin.PostAddReqEntity
	if err := admin.PostAddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err = adminService.Post.Edit(&data, id)
	if err != nil {
		g.Log().Error("编辑失败:", err)
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	}
	response.RedirectToWithMessage(r, "/admin/posts", "编辑成功")
}

func (c *PostController) Del(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	item, err := g.DB().Table(posts.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("文章不存在或已被删除"))
	}
	err = adminService.Post.Delete(id)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("删除失败"))
	}
	response.RedirectToWithMessage(r, "/admin/posts", "删除成功")
}

func (c *PostController) Show(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	item, err := g.DB().Table(posts.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("文章不存在或已被删除"))
	}
	post, comments, err := adminService.Post.Show(id)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("获取失败"))
	}
	response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
		"mainTpl":  fmt.Sprintf(constants.AdminShowTpl, "post"),
		"post":     post,
		"comments": comments,
	})
}
