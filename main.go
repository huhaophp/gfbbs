package main

import (
	_ "bbs/boot"
	_ "bbs/router/web"
	_ "bbs/router/admin"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
