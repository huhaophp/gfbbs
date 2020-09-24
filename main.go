package main

import (
	_ "bbs/boot"
	_ "bbs/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
