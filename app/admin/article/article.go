package article

import (
	"bbs/app/model/articles"
	"bbs/app/model/categories"
	response "bbs/library"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/article/list.html"
	createTpl string = "admin/article/create.html"
	editTpl   string = "admin/article/edit.html"
	errorTpl  string = "admin/error.html"
)

// Controller Base
type Controller struct{}

// List Article list
func (c *Controller) List(r *ghttp.Request) {
	var req articles.ListReqEntity
	err := r.Parse(&req)
	if err != nil {
		response.ViewExit(r, layout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	}
	total, list, err := articles.List(&req)
	if err != nil {
		response.ViewExit(r, layout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	}

	page := r.GetPage(total, 10)

	response.ViewExit(r, layout, g.Map{"list": list, "mainTpl": listTpl, "page": page.GetContent(3), "req": req})
}

// Add article added
func (c *Controller) Add(r *ghttp.Request) {
	if r.Method == "GET" {
		var req categories.ListReqEntity
		cates, _ := categories.List(&req)
		response.ViewExit(r, layout, g.Map{"mainTpl": createTpl, "categories": cates})
	}
	var data *articles.AddReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	insid := articles.Add(data)
	if insid <= 0 {
		response.RedirectBackWithError(r, gerror.New("添加失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/articles", "添加成功")
	}
}

// Edit edit article.
func (c *Controller) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if r.Method == "GET" {
		item, _ := articles.Find(id)
		var req categories.ListReqEntity
		cates, _ := categories.List(&req)
		response.ViewExit(r, layout, g.Map{"mainTpl": editTpl, "item": item, "categories": cates})
	}
	var data *articles.AddReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows := articles.Edit(id, data)
	if rows <= 0 {
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/articles", "编辑成功")
	}
}

// Delete delete article
func (c *Controller) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows := articles.Del(id)
	if rows <= 0 {
		response.RedirectBackWithError(r, errors.New("删除失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/articles", "删除成功")
	}
}
