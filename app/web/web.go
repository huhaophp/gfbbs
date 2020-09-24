package web

import (
	"bbs/app/model/articles"
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

const (
	loyout   string = "web/layout.html"
	homeTpl  string = "web/home.html"
	showTpl  string = "web/show.html"
	errorTpl string = "web/error.html"
)

// Controller Base
type Controller struct{}

// Home home.
func (c *Controller) Home(r *ghttp.Request) {
	var req articles.ListReqEntity
	err := r.Parse(&req)
	if err != nil {
		response.ViewExit(r, loyout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	}

	total, list, err := articles.List(&req)
	if err != nil {
		response.ViewExit(r, loyout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	}

	data := g.Map{
		"mainTpl": homeTpl,
		"list":    list,
		"cid":     req.CategoryId,
		"len":     len(list),
		"page":    r.GetPage(total, 10).GetContent(3),
	}

	response.ViewExit(r, loyout, data)
}

// Show Article details.
func (c *Controller) Show(r *ghttp.Request) {
	value, ok := r.GetRouterValue("id").(string)
	if !ok {
		response.ViewExit(r, loyout, g.Map{"mainTpl": errorTpl, "error": "Id param error"})
	}
	id, err := strconv.Atoi(value)
	if err != nil {
		response.ViewExit(r, loyout, g.Map{"mainTpl": errorTpl, "error": "Id param error1"})
	}

	article, err := articles.Find(id)
	if err != nil {
		response.ViewExit(r, loyout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	}

	data := g.Map{
		"mainTpl": showTpl,
		"article": article,
		"cid":     article["category_id"],
	}

	response.ViewExit(r, loyout, data)
}
