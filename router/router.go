package router

import (
	"bbs/app/admin/admin"
	"bbs/app/admin/auth"
	"bbs/app/admin/cate"
	"bbs/app/admin/comment"
	"bbs/app/admin/file"
	"bbs/app/admin/home"
	"bbs/app/admin/node"
	"bbs/app/admin/post"
	"bbs/app/admin/user"
	"bbs/app/middleware"
	"bbs/app/web"
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {

	s := g.Server()

	authController := new(auth.Controller)
	homeController := new(home.Controller)
	fileController := new(file.Controller)
	nodeController := new(node.Controller)
	cateController := new(cate.Controller)
	postController := new(post.Controller)
	userController := new(user.Controller)
	adminController := new(admin.Controller)
	commentController := new(comment.Controller)

	// admin routes.
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.GET("login", authController.Login)
		group.POST("login", authController.Login)
		group.Middleware(middleware.AdminAuthCheck)
		group.POST("logout", authController.Logout)
		group.GET("home", homeController.Home)

		group.GET("admins", adminController.List)
		group.GET("nodes", nodeController.List)
		group.GET("cates", cateController.List)
		group.GET("posts", postController.List)
		group.GET("users", userController.List)
		group.GET("comments", commentController.List)
		group.POST("file", fileController.Store)
	})

	// web routes.
	webController := new(web.Controller)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", webController.Home)
	})

	// Handling 404 pages
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		response.NotFoundView(r)
	})
}
