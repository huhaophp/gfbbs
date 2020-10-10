package comment

import (
	"bbs/app/model/comments"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type AddCommentReqEntity struct {
	Content string `p:"content" v:"required#请输入回复内容"`
	Pid     int    `p:"pid"`
	Ruid    int    `p:"ruid"`
	Uid     int    `p:"uid"`
}

// Add Comment post
func Add(entity *AddCommentReqEntity) error {
	res, err := g.DB().Table(comments.Table).Insert(g.Map{
		"pid":       entity.Pid,
		"uid":       entity.Uid,
		"ruid":      entity.Ruid,
		"content":   entity.Content,
		"is_delete": 0,
	})
	if err != nil {
		return err
	}
	if id, err := res.LastInsertId(); err != nil || id <= 0 {
		return gerror.New("评论失败")
	}
	return nil
}

// Delete Delete comment
func Delete(id string) error {
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
func CheckPermissions(id string, uid string) error {
	res, err := g.DB().Table(comments.Table).Where(g.Map{"uid": uid, "is_delete": 0, "id": id}).One()
	if err != nil {
		return err
	}
	if res.IsEmpty() {
		return gerror.New("评论不存在或已被删除")
	}
	return nil
}
