package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/admins"
	"bbs/app/request/admin"
	adminService "bbs/app/service/admin"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

type AdminController struct{}

func (c *AdminController) List(r *ghttp.Request) {
	items, err := g.DB().Table(admins.Table).All()
	if err != nil {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": constants.AdminErrorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminListTpl, "admin"), "items": items})
	}
}

func (c *AdminController) Add(r *ghttp.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminCreateTpl, "admin")})
	}
	var data admin.AdminAddReqEntity
	if err := admin.AdminAddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err := adminService.AdminAdd(&data)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, "/admin/admins", "添加成功")
}

func (c *AdminController) Edit(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	item, err := g.DB().Table(admins.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("管理员未找到"))
	}
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminEditTpl, "admin"), "admin": item})
	}
	var data admin.AdminUpdateReqEntity
	if err := admin.AdminUpdateReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err = adminService.AdminEdit(&data, id)
	if err != nil {
		g.Log().Error("编辑失败:", err)
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	}
	response.RedirectToWithMessage(r, "/admin/admins", "编辑成功")
}

func (c *AdminController) Delete(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	err := adminService.AdminDelete(id)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("删除失败"))
	}
	response.RedirectToWithMessage(r, "/admin/admins", "删除成功")
}
