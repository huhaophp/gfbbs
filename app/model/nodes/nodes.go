// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package nodes

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// Fill with you ideas below.
type QueryParamEntity struct {
	IsTop       int
	IsDelete    int
	Sort        string
	ISChildNode bool
	Pid         int
}

func Get(filter *QueryParamEntity) (items gdb.Result) {
	query := g.DB().Table(Table)
	if isTop := filter.IsTop; isTop >= 0 {
		query = query.Where("is_top = ?", isTop)
	}
	if isDel := filter.IsDelete; isDel >= 0 {
		query = query.Where("is_delete = ?", isDel)
	}
	if filter.ISChildNode {
		query = query.Where("pid != ?", 0)
	}
	if pid := filter.Pid; pid > 0 {
		query = query.Where("pid = ?", pid)
	}
	if sort := filter.Sort; sort != "" {
		query = query.Order(sort)
	}
	items, _ = query.All()
	return
}