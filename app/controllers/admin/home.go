package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type HomeController struct{}

// GET|后台首页
func (c *HomeController) Home(r *ghttp.Request) {
	response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{"mainTpl": constants.AdminHomeTplPath})
}
