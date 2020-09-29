package main

import (
	_ "bbs/boot"
	_ "bbs/router/admin"
	_ "bbs/router/web"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
