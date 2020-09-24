package boot

import (
	"bbs/app/service/config"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

func init() {
	_ = gtime.SetTimeZone("Asia/Shanghai")
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	g.Server().AddStaticPath("/public", g.Cfg().GetString("server.ServerRoot"))

	config.CategoryGlobalVariableSettings()
	config.SiteGlobalVariableSettings()
}
