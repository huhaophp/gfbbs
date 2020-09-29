package admin

import (
	"bbs/app/controllers/admin/admin"
	"bbs/app/controllers/admin/auth"
	"bbs/app/controllers/admin/cate"
	"bbs/app/controllers/admin/comment"
	"bbs/app/controllers/admin/home"
	"bbs/app/controllers/admin/node"
	"bbs/app/controllers/admin/post"
	"bbs/app/controllers/admin/user"
	"bbs/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	authController := new(auth.Controller)
	homeController := new(home.Controller)
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
		group.GET("admins/add", adminController.Add)
		group.POST("admins/add", adminController.Add)
		group.GET("admins/{id}/edit", adminController.Edit)
		group.POST("admins/{id}/edit", adminController.Edit)
		group.POST("admins/{id}/delete", adminController.Delete)

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
	})
}
