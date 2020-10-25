package admin

import (
	"bbs/app/funcs/response"
	"bbs/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	SensitiveListTpl = "admin/sensitive/list.html"
)

type SensitiveController struct{}

func (s *SensitiveController) List(r *ghttp.Request) {
	pageNum := r.GetQueryInt("page", 1)

	total := service.SensitiveService.Total()
	items := service.SensitiveService.List(pageNum, 20)

	page := r.GetPage(total, 20)

	data := g.Map{"items": items, "page": page.GetContent(2), "mainTpl": SensitiveListTpl}

	response.ViewExit(r, Layout, data)
}

func (s *SensitiveController) Add() {

}

func (s *SensitiveController) Edit() {

}

func (s *SensitiveController) Del() {

}
