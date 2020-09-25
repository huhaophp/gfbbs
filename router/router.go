package router

import (
	"bbs/app/middleware"
	"bbs/app/server/admin/admin"
	"bbs/app/server/admin/auth"
	"bbs/app/server/admin/cate"
	"bbs/app/server/admin/comment"
	"bbs/app/server/admin/file"
	"bbs/app/server/admin/home"
	"bbs/app/server/admin/node"
	"bbs/app/server/admin/post"
	"bbs/app/server/admin/user"
	"bbs/app/server/web"
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
		group.GET("nodes/add", nodeController.Add)
		group.POST("nodes/add", nodeController.Add)
		group.GET("nodes/{id}/edit", nodeController.Edit)
		group.POST("nodes/{id}/edit", nodeController.Edit)
		group.POST("nodes/{id}/delete", nodeController.Del)

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
