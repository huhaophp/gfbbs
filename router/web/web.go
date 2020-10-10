package web

import (
	"bbs/app/controllers/web"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 初始化web路由
func init() {
	webController := new(web.Controller)
	PostsController := new(web.PostsController)
	fileController := new(web.FileController)
	userController := new(web.UserController)
	nodeController := new(web.NodeController)
	commentController := new(web.CommentController)
	captchaController := new(web.CaptchaController)
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", webController.Home)
		group.GET("/captcha", captchaController.Get)
		group.GET("/posts/{postsId}", PostsController.Details)
		group.GET("/node/{nodeId}", nodeController.Index)
		group.POST("/comments", commentController.Add)
		group.POST("/comments/{id}", commentController.Del)
		group.POST("/editor/file", fileController.WangEditorFileStore)
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
