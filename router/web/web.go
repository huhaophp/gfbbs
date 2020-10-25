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
		// 用户登录
		group.GET("/user/login", userController.Login)
		group.POST("/user/login", userController.Login)
		// 用户注册
		group.GET("/user/register", userController.Register)
		group.POST("/user/register", userController.Register)
		// 用户中心
		group.GET("/users/{id}", userController.Center)
		// 需要授权路由
		group.Middleware(middleware.WebAuthCheck)
		// 发布帖子
		group.GET("/posts/publish", PostsController.Publish)
		group.POST("/posts/publish", PostsController.Publish)
		// 发布评论
		group.POST("/comments", commentController.Add)
		// 删除评论
		group.POST("/comments/{id}/delete", commentController.Del)
		// 文件上传
		group.POST("/file", fileController.Upload)
		// 用户退出登录
		group.POST("/user/logout", userController.Logout)
		// 用户编辑
		group.GET("/user/edit", userController.Edit)
		group.POST("/user/edit", userController.Edit)
		// 消息中心
		group.GET("/message", MessageController.Index)
		// 帖子点赞
		group.POST("/like/do", LikeController.Do)
		group.POST("/like/undo", LikeController.Undo)
	})
}
