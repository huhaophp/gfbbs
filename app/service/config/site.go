package config

import (
	"bbs/app/model/categories"
	"github.com/gogf/gf/frame/g"
)

func CategoryGlobalVariableSettings() {
	var req categories.ListReqEntity
	req.Status = 1
	data, _ := categories.List(&req)
	g.View().Assign("categories", data)
}

func SiteGlobalVariableSettings() {
	//list, _ := g.DB().Table(configs.Table).Where("`key` IN(?) ", g.Slice{"name", "copyright"}).All()
}
