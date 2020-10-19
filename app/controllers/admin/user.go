package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/users"
	"bbs/app/request/admin"
	adminService "bbs/app/service/admin"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

type UserController struct{}

func (c *UserController) List(r *ghttp.Request) {
	items, err := g.DB().Table(users.Table).All()
	if err != nil {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": constants.AdminErrorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminListTpl, "user"), "items": items})
	}
}

func (c *UserController) Add(r *ghttp.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminCreateTpl, "user")})
	}
	var data admin.UserAddReqEntity
	if err := admin.UserAddReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err := adminService.UserAdd(&data)
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	response.RedirectToWithMessage(r, "/admin/users", "添加成功")
}

func (c *UserController) Edit(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, gerror.New("id错误"))
	}
	item, err := g.DB().Table(users.Table).Where("id = ?", id).One()
	if err != nil || item == nil {
		response.RedirectBackWithError(r, gerror.New("用户未找到"))
	}
	if strings.ToUpper(r.Method) == "GET" {
		response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": fmt.Sprintf(constants.AdminEditTpl, "user"), "user": item})
	}
	var data admin.UserUpdateReqEntity
	if err := admin.UserUpdateReqCheck(r, &data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	err = adminService.UserEdit(&data, id)
	if err != nil {
		g.Log().Error("编辑失败:", err)
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	}
	response.RedirectToWithMessage(r, "/admin/users", "编辑成功")
}

func (c *UserController) Delete(r *ghttp.Request) {
	id := r.GetRouterVar("id").Int()
	if id <= 0 {
		response.RedirectBackWithError(r, errors.New("id错误"))
	}
	err := adminService.UserDelete(id)
	if err != nil {
		g.Log().Error("删除失败:", err)
		response.RedirectBackWithError(r, errors.New("删除失败"))
	}
	response.RedirectToWithMessage(r, "/admin/users", "删除成功")
}
