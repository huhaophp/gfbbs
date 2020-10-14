package service

import (
	"bbs/app/model/comments"
	"bbs/app/model/posts"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type AddCommentReqEntity struct {
	Content string `p:"content" v:"required#请输入回复内容"`
	Pid     int    `p:"pid"`
	Ruid    int    `p:"ruid"`
	Uid     int    `p:"uid"`
	Puid    int    `p:"puid"`
}

var CommentService = newCommentService()

// nodeService Initialize the service
func newCommentService() *commentService {
	return &commentService{}
}

// nodeService
type commentService struct{}

// Add Comment post
func (s *commentService) Add(entity *AddCommentReqEntity) error {
	res, err := g.DB().Table(comments.Table).Insert(g.Map{
		"pid":       entity.Pid,
		"uid":       entity.Uid,
		"ruid":      entity.Ruid,
		"content":   entity.Content,
		"is_delete": 0,
	})
	_, _ = g.DB().Table(posts.Table).Data("comment_num = comment_num+1").Where("id = ?", entity.Pid).Update()
	if err != nil {
		return err
	}
	if id, err := res.LastInsertId(); err != nil || id <= 0 {
		return gerror.New("评论失败")
	}
	// Judge the recipient.
	action := "comment"
	if entity.Ruid > 0 {
		entity.Puid = entity.Ruid
		action = "reply"
	}
	// Send message to recipient.
	_ = MessageService.Send(g.Map{
		"suid":      entity.Uid,
		"ruid":      entity.Puid,
		"tid":       entity.Pid,
		"type":      "posts",
		"action":    action,
		"is_read":   0,
		"is_delete": 0,
	})
	return nil
}

// Delete Delete comment
func (s *commentService) Delete(id string) error {
	res, err := g.DB().Table(comments.Table).WherePri(id).Update(g.Map{
		"is_delete": 1,
	})
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows <= 0 {
		return gerror.New("删除失败")
	}
	return nil
}

// CheckPermissions
func (s *commentService) CheckPermissions(id string, uid string) error {
	res, err := g.DB().Table(comments.Table).Where(g.Map{"uid": uid, "is_delete": 0, "id": id}).One()
	if err != nil {
		return err
	}
	if res.IsEmpty() {
		return gerror.New("评论不存在或已被删除")
	}
	return nil
}
