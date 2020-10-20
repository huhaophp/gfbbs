package admin

import (
	"bbs/app/controllers/admin"
	"bbs/app/controllers/admin/comment"
	"bbs/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	authController := new(admin.AuthController)
	homeController := new(admin.HomeController)
	nodeController := new(admin.NodeController)
	postController := new(admin.PostController)
	userController := new(admin.UserController)
	adminController := new(admin.AdminController)
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

		group.GET("users", userController.List)
		group.GET("users/add", userController.Add)
		group.POST("users/add", userController.Add)
		group.GET("users/{id}/edit", userController.Edit)
		group.POST("users/{id}/edit", userController.Edit)
		group.POST("users/{id}/delete", userController.Delete)

		group.GET("nodes", nodeController.List)
		group.GET("nodes/add", nodeController.Add)
		group.POST("nodes/add", nodeController.Add)
		group.GET("nodes/{id}/edit", nodeController.Edit)
		group.POST("nodes/{id}/edit", nodeController.Edit)
		group.POST("nodes/{id}/delete", nodeController.Del)

		group.GET("posts", postController.List)
		group.GET("posts/{id}/show", postController.Show)
		group.GET("posts/add", postController.Add)
		group.POST("posts/add", postController.Add)
		group.GET("posts/{id}/edit", postController.Edit)
		group.POST("posts/{id}/edit", postController.Edit)
		group.POST("posts/{id}/delete", postController.Del)

		group.GET("comments", commentController.List)
	})
}
