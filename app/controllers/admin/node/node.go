package node

import (
	"bbs/app/model/nodes"
	"bbs/app/request/node"
	response "bbs/app/funcs/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"strconv"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/node/list.html"
	createTpl string = "admin/node/create.html"
	editTpl   string = "admin/node/edit.html"
	errorTpl  string = "admin/error.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	items, err := g.DB().Table(nodes.Table+" n1").
		LeftJoin("nodes n2", "n1.pid = n2.id").
		Fields("n1.*,n2.name as pname").
		Order("n1.sort DESC").
		All()
	if err != nil {
		response.ViewExit(r, layout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl, "items": items})
	}
}

func (c *Controller) Add(r *ghttp.Request) {
	if r.Method == "GET" {
		items, _ := g.DB().Table(nodes.Table).Where("pid = ?", 0).All()
		response.ViewExit(r, layout, g.Map{"mainTpl": createTpl, "nodes": items})
	}
	var data node.AddReqEntity
	if err := node.AddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	exists, _ := g.DB().Table(nodes.Table).Where("name = ?", data.Name).One()
	if exists != nil {
		response.RedirectBackWithError(r, gerror.New("节点名称已存在"))
	}
	res, err := g.DB().Table(nodes.Table).Insert(g.Map{
		"name":       data.Name,
		"sort":       data.Sort,
		"pid":        data.Pid,
		"desc":       data.Desc,
		"create_at": gtime.Now(),
		"update_at": gtime.Now(),
	})
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	id, err := res.LastInsertId()
	if err != nil || id <= 0 {
		response.RedirectBackWithError(r, gerror.New("添加失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/nodes", "添加成功")
	}
}

func (c *Controller) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.RedirectBackWithError(r, err)
	}

	if r.Method == "GET" {
		item, err := g.DB().Table(nodes.Table).Where("id = ?", id).One()
		if err != nil || item == nil {
			response.RedirectBackWithError(r, gerror.New("节点未找到"))
		}
		items, _ := g.DB().Table(nodes.Table).Where("pid = ? and is_delete != ?", 0, 1).All()

		response.ViewExit(r, layout, g.Map{"mainTpl": editTpl, "nodes": items, "node": item})
	}

	var data node.AddReqEntity
	if err := node.AddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	exists, _ := g.DB().Table(nodes.Table).Where("name = ? and id != ?", data.Name, id).One()
	if exists != nil {
		response.RedirectBackWithError(r, gerror.New("节点名称已存在"))
	}
	res, err := g.DB().Table(nodes.Table).WherePri(id).Update(g.Map{
		"name":       data.Name,
		"sort":       data.Sort,
		"pid":        data.Pid,
		"desc":       data.Desc,
		"status":     data.Status,
		"update_at": gtime.Now(),
	})
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/nodes", "编辑成功")
	}
}

func (c *Controller) Del(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	item, err := g.DB().Table(nodes.Table).Where("id = ? and is_delete = ?", id, 0).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("节点不存在或已被删除"))
	}
	res, err := g.DB().Table(nodes.Table).WherePri(id).Update(g.Map{
		"is_delete":  1,
		"update_at": gtime.Now(),
	})
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		response.RedirectBackWithError(r, gerror.New("删除失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/nodes", "删除成功")
	}
}
