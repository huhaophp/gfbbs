package web

import (
	"bbs/app/server/web"
	"bbs/app/server/web/file"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// web routes.
	webController := new(web.Controller)
	fileController := new(file.Controller)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", webController.Home)
		group.GET("/posts/{postsId}", webController.PostDetail)
		group.POST("/markdown/file", fileController.MarkdownFileStore)
	})
}
