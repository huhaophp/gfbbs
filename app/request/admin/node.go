package admin

import (
	"github.com/gogf/gf/net/ghttp"
)

type NodeAddReqEntity struct {
	Name   string `p:"name" v:"required#请输入节点名称"`
	Sort   string `p:"sort" v:"required#请填写排序"`
	Desc   string `p:"desc" v:"required#请填写节点介绍"`
	Status int    `p:"status" v:"required#请填写节点状态|in:0,1#节点状态只能在0,1之间"`
	IsTop int    `p:"is_top" v:"required#请填写节点置顶状态|in:0,1#节点置顶状态只能在0,1之间"`
}

func NodeAddReqCheck(r *ghttp.Request, data *NodeAddReqEntity) error {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}
