package main

import (
	"bbs/app/funcs/response"
	_ "bbs/boot"
	_ "bbs/router/admin"
	_ "bbs/router/web"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	g.Server().Run()
	// Handling 404 pages
	g.Server().BindStatusHandler(404, func(r *ghttp.Request) {
		response.NotFoundView(r)
	})
}
