package web

import (
	"bbs/app/controllers/web"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 初始化web路由
func init() {
	webController := new(web.Controller)
	fileController := new(web.FileController)
	userController := new(web.UserController)
	nodeController := new(web.NodeController)
	commentController := new(web.CommentController)
	captchaController := new(web.CaptchaController)
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", webController.Home)
		group.GET("/captcha", captchaController.Get)
		group.GET("/posts/{postsId}", webController.PostDetail)
		group.GET("/node/{nodeId}", nodeController.Index)
		group.POST("/comment", commentController.Add)
		group.POST("/markdown/file", fileController.MdFileStore)
		group.POST("/file", fileController.FileStore)
		group.GET("/user/login", userController.Login)
		group.POST("/user/login", userController.Login)
		group.POST("/user/logout", userController.Logout)
		group.GET("/user/register", userController.Register)
		group.POST("/user/register", userController.Register)
		group.GET("/user/edit", userController.Edit)
		group.POST("/user/edit", userController.Edit)
		group.GET("/users/{id}", userController.Center)
	})
}
