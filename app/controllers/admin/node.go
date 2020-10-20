package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/nodes"
	"bbs/app/request/admin"
	adminService "bbs/app/service/admin"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

type NodeController struct{}

func (c *NodeController) List(r *ghttp.Request) {
	pageNum := r.GetQueryInt("page", 1)

	items, err := g.DB().Table(nodes.Table).
		Where("is_delete = ?", 0).
		Order("sort DESC").
		Order("id DESC").
		Page(pageNum, 20).
		All()
	total, _ := g.DB().Table(nodes.Table).Where("is_delete = ?", 0).Count()
	page := r.GetPage(total, 20)

	if err != nil {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": constants.AdminErrorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
			"mainTpl": fmt.Sprintf(constants.AdminListTpl, "node"),
			"items":   items,
			"page":    page.GetContent(2),
		})
	}
}

func (c *NodeController) Add(r *ghttp.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminCreateTpl, "node")})
	}
	var data admin.NodeAddReqEntity
	if err := admin.NodeAddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err := adminService.Node.Add(&data)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, "/admin/nodes", "添加成功")
}

func (c *NodeController) Edit(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	item, err := g.DB().Table(nodes.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("节点未找到"))
	}
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminEditTpl, "node"), "node": item})
	}
	var data admin.NodeAddReqEntity
	if err := admin.NodeAddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err = adminService.Node.Edit(&data, id)
	if err != nil {
		g.Log().Error("编辑失败:", err)
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	}
	response.RedirectToWithMessage(r, "/admin/nodes", "编辑成功")
}

func (c *NodeController) Del(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	item, err := g.DB().Table(nodes.Table).Where("id = ? and is_delete = ?", id, 0).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("节点不存在或已被删除"))
	}
	err = adminService.Node.Delete(id)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("删除失败"))
	}
	response.RedirectToWithMessage(r, "/admin/nodes", "删除成功")
}
