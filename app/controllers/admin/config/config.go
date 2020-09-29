package config

import (
	"bbs/app/model/configs"
	response "bbs/app/funcs/response"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

const (
	layout  string = "admin/layout.html"
	listTpl string = "admin/config/list.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	list := configs.List()
	response.ViewExit(r, layout, g.Map{"mainTpl": listTpl, "list": list})
}

func (c *Controller) Add(r *ghttp.Request) {
	var data *configs.AddReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	insid := configs.Add(data)
	if insid <= 0 {
		response.RedirectBackWithError(r, gerror.New("添加失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/configs", "添加成功")
	}
}

func (c *Controller) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	var data *configs.AddReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows := configs.Edit(id, data)
	if rows <= 0 {
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/configs", "编辑成功")
	}
}

func (c *Controller) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.Json(r, -1, err.Error())
	}
	rows := configs.Del(id)
	if rows <= 0 {
		response.RedirectBackWithError(r, errors.New("删除失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/configs", "删除成功")
	}
}
