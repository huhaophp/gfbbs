package admin

import (
	"bbs/app/funcs/response"
	"bbs/app/model/admins"
	"bbs/app/request/Admin"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/admin/list.html"
	createTpl string = "admin/admin/create.html"
	editTpl   string = "admin/admin/edit.html"
	errorTpl  string = "admin/error.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	items, err := g.DB().Table(admins.Table).All()
	if err != nil {
		response.ViewExit(r, layout, g.Map{"mainTpl": errorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl, "items": items})
	}
}

func (c *Controller) Add(r *ghttp.Request)  {
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": createTpl})
	}
	var data Admin.AddReqEntity
	if err := Admin.AddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	exists, _ := g.DB().Table(admins.Table).Where("email = ?", data.Email).One()
	if exists != nil {
		response.RedirectBackWithError(r, gerror.New("邮箱已存在"))
	}
	password, _ := gmd5.Encrypt(data.Password)
	res, err := g.DB().Table(admins.Table).Insert(g.Map{
		"name":       data.Name,
		"email":       data.Email,
		"password":		password,
		"status":        data.Status,
	})
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	id, err := res.LastInsertId()
	if err != nil || id <= 0 {
		g.Log().Error("添加失败：", err)
		response.RedirectBackWithError(r, gerror.New("添加失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/admins", "添加成功")
	}
}

func (c *Controller) Edit(r *ghttp.Request)  {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	item, err := g.DB().Table(admins.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("管理员未找到"))
	}
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": editTpl, "admin": item})
	}
	var data Admin.UpdateReqEntity
	if err := Admin.UpdateReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	exists, _ := g.DB().Table(admins.Table).Where("email = ? and id != ?", data.Email, id).One()
	if exists != nil {
		response.RedirectBackWithError(r, gerror.New("邮箱已存在"))
	}
	gMap := g.Map{
		"name": data.Name,
		"email": data.Email,
		"status": data.Status,
	}
	if data.Password != "" {
		password, _ := gmd5.Encrypt(data.Password)
		gMap["password"] = password
	}
	res, err := g.DB().Table(admins.Table).WherePri(id).Update(gMap)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/admins", "编辑成功")
	}
}

func (c *Controller) Delete(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	res, err := g.DB().Table(admins.Table).Delete("id", id)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		response.RedirectBackWithError(r, gerror.New("删除失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/admins", "删除成功")
	}

}
