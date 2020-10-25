package service

import (
	"bbs/app/model/sensitives"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var SensitiveService = newSensitiveService()

// newSensitiveService Initialize the service
func newSensitiveService() *sensitiveService {
	return &sensitiveService{}
}

// sensitiveService
type sensitiveService struct{}

func (s *sensitiveService) List(page int, limit int) gdb.Result {
	items, _ := g.DB().Table(sensitives.Table).
		Page(page, limit).
		All()
	return items
}

func (s *sensitiveService) Total() (total int) {
	total, _ = g.DB().Table(sensitives.Table).Count()
	return
}
