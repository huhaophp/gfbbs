package service

import (
	"bbs/app/model/messages"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var MessageService = newMessageService()

// newMessageService Initialize the service
func newMessageService() *messageService {
	return &messageService{}
}

// messageService
type messageService struct{}

// List Get message list.
func (s *messageService) List(ruid int, page int, limit int) gdb.Result {
	items, _ := g.DB().Table(messages.Table+" m").
		LeftJoin("users u", "u.id = m.suid").
		LeftJoin("posts p", "p.id = m.tid").
		Where("m.ruid = ?", ruid).
		Fields("m.id,m.suid,m.ruid,m.tid,m.type,m.action,m.update_at,m.create_at,u.name,u.avatar,p.title").
		Order("m.create_at DESC").
		Page(page, 20).
		All()
	return items
}

// GetUnread Get unread messages.
func (s *messageService) GetUnreadNum(uid int) int {
	count, err := g.DB().Table(messages.Table).Where("ruid = ? and is_read = 0", uid).Count()
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return count
}

// Send Send message to user.
func (s *messageService) Send(p g.Map) error {
	res, err := g.DB().Table(messages.Table).Insert(p)
	if err != nil {
		g.Log().Error(err)
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	if id <= 0 {
		g.Log().Error("Failed to send message.")
		return gerror.New("Failed to send message.")
	}
	return nil
}

// ReadAll Read all unread messages.
func (s *messageService) ReadAll(userId int) error {
	res, err := g.DB().Table(messages.Table).Where("ruid =? and is_read = 0", userId).Data(g.Map{"is_read": 1}).Update()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	id, err := res.RowsAffected()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	if id <= 0 {
		g.Log().Error("Failed to read unread messages.")
		return gerror.New("Failed to read unread messages.")
	}
	return nil
}

// Del Delete a single message.
func (s *messageService) Del(mid int, uid int) error {
	res, err := g.DB().Table(messages.Table).Where("id = ? and ruid = ?", mid, uid).Data(g.Map{"is_delete": 1}).Update()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	id, err := res.RowsAffected()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	if id <= 0 {
		g.Log().Error("Delete message error.")
		return gerror.New("Delete message error.")
	}
	return nil
}
