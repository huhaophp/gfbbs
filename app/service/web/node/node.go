package node

import (
	"bbs/app/model/nodes"
	"github.com/gogf/gf/database/gdb"
)

// GetTopNodes 获取节点
func GetNodes(iSChild bool, pid int) gdb.Result {
	params := nodes.QueryParamEntity{IsTop: 1, IsDelete: 0, Sort: "sort DESC", ISChildNode: iSChild, Pid: pid}
	return nodes.Get(&params)
}
