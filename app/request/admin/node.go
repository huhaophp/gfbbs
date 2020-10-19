package admin

import (
	"github.com/gogf/gf/net/ghttp"
)

type NodeAddReqEntity struct {
	Name   string `p:"name" v:"required#请输入节点名称"`
	Pid    string `p:"pid" v:"required#请选择所属节点"`
	Sort   string `p:"sort" v:"required#请填写排序"`
	Status int    `p:"status"`
	Desc   string `p:"desc" v:"required#请填写节点介绍"`
}

func NodeAddReqCheck(r *ghttp.Request, data *NodeAddReqEntity) error {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}
