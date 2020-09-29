package boot

import (
	"bbs/app/funcs/response"
	"bbs/app/funcs/view"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gview"
)

func init() {
	initSystemSetting()
	initViewFunctions()
	// Handling routing 404 errors
	g.Server().BindStatusHandler(404, func(r *ghttp.Request) {
		response.PageNotFound(r)
	})
}

// initSystemSetting 初始化系统设置
func initSystemSetting() {
	_ = gtime.SetTimeZone("Asia/Shanghai")
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	g.Server().AddStaticPath("/public", g.Cfg().GetString("server.ServerRoot"))
}

// initViewFunctions 初始化模版全局函数
func initViewFunctions() {
	g.View().BindFuncMap(gview.FuncMap{
		"StrTime": view.StrTime,
	})
}
