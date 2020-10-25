package boot

import (
	"bbs/app/funcs/view"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gview"
)

func init() {
	initSystemSetting()
	initViewFunctions()
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
		"StrTime":          view.StrTime,
		"StrLimit":         view.StrLimit,
		"AlertComponent":   view.AlertComponent,
	})
}
