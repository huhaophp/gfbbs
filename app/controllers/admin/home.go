package admin

import (
	"bbs/app/constants"
	"bbs/app/funcs/response"
	"bbs/app/model/posts"
	"bbs/app/model/users"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type HomeController struct{}

// GET|后台首页
func (c *HomeController) Home(r *ghttp.Request) {
	userCount,_ := g.DB().Table(users.Table).Count()
	postCount,_ := g.DB().Table(posts.Table).Count()

	response.ViewExit(r, constants.AdminLayoutTplPath, g.Map{
		"mainTpl": constants.AdminHomeTplPath,
		"user": userCount,
		"post": postCount,
	})
}
