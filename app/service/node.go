package service

import (
	"bbs/app/model/nodes"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var NodeService = newNodeService()

// nodeService Initialize the service
func newNodeService() *nodeService {
	return &nodeService{}
}

// nodeService
type nodeService struct{}

// GetAllNormalNodes
func (s *nodeService) Get(where g.Map) gdb.Result {
	res, err := g.DB().Table(nodes.Table).Where(where).Order("sort").All()
	if err != nil {
		g.Log().Error("GetAllNormalNodes error")
		return nil
	}
	return res
}
