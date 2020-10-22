package web

import (
	"bbs/app/controllers/web"
	"bbs/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 初始化web路由
func init() {
	webController := new(web.Controller)
	fileController := new(web.FileController)
	userController := new(web.UserController)
	LikeController := new(web.LikeController)
	PostsController := new(web.PostsController)
	commentController := new(web.CommentController)
	captchaController := new(web.CaptchaController)
	MessageController := new(web.MessageController)
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 设置 layout 全局视图变量
		group.Middleware(middleware.WebLayoutGlobalVariablesSetting)
		// 论坛首页
		group.GET("/", webController.Home)
		// 帖子详情
		group.GET("/posts/{postsId}", PostsController.Details)
		// 验证码
		group.GET("/captcha", captchaController.Get)
		// 用户操作
		group.GET("/user/login", userController.Login)
		group.POST("/user/login", userController.Login)
		group.GET("/user/register", userController.Register)
		group.POST("/user/register", userController.Register)
		group.GET("/users/{id}", userController.Center)
		// 需要授权路由
		group.Middleware(middleware.WebAuthCheck)
		group.GET("/posts/publish", PostsController.Publish)
		group.POST("/posts/publish", PostsController.Publish)
		group.POST("/comments", commentController.Add)
		group.POST("/comments/{id}/delete", commentController.Del)
		group.POST("/file", fileController.Upload)
		group.POST("/user/logout", userController.Logout)
		group.GET("/user/edit", userController.Edit)
		group.POST("/user/edit", userController.Edit)
		group.GET("/message", MessageController.Index)
		group.POST("/like/do", LikeController.Do)
		group.POST("/like/undo", LikeController.Undo)
	})
}
