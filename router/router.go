package router

import (
	"bbs/app/admin/article"
	"bbs/app/admin/auth"
	"bbs/app/admin/categroy"
	"bbs/app/admin/config"
	"bbs/app/admin/file"
	"bbs/app/admin/home"
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
	cateController := new(categroy.Controller)
	articleController := new(article.Controller)
	configController := new(config.Controller)

	// admin routes.
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.GET("login", authController.Login)
		group.POST("login", authController.Login)
		group.Middleware(middleware.AdminAuthCheck)
		group.POST("logout", authController.Logout)
		group.GET("home", homeController.Home)
		group.GET("categories", cateController.List)
		group.GET("categories/add", cateController.Add)
		group.POST("categories/add", cateController.Add)
		group.GET("categories/{id}/edit", cateController.Edit)
		group.POST("categories/{id}/edit", cateController.Edit)
		group.POST("categories/{id}/delete", cateController.Delete)
		group.GET("articles", articleController.List)
		group.GET("articles/add", articleController.Add)
		group.POST("articles/add", articleController.Add)
		group.GET("articles/{id}/edit", articleController.Edit)
		group.POST("articles/{id}/edit", articleController.Edit)
		group.POST("articles/{id}/delete", articleController.Delete)
		group.POST("file", fileController.Store)
		group.POST("markdown/file", fileController.MarkdownFileStore)
		// 通用配置
		group.GET("configs", configController.List)
		group.POST("configs/add", configController.Add)
		group.POST("configs/{id}/edit", configController.Edit)
		group.POST("configs/{id}/delete", configController.Delete)
	})

	// web routes.
	webController := new(web.Controller)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", webController.Home)
		group.GET("/articles/{id}", webController.Show)
	})

	// Handling 404 pages
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		response.NotFoundView(r)
	})
}
