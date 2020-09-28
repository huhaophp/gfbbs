package comment

import (
	"bbs/app/model/comments"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type AddReqEntity struct {
	Content string `p:"content" v:"required#请输入回复内容"`
	Pid     int    `p:"pid"`
	Ruid    int    `p:"ruid"`
	Uid     int    `p:"uid"`
}

// CommentPost
func CommentPost(entity *AddReqEntity) error {
	result, err := g.DB().Table(comments.Table).Insert(g.Map{
		"pid":        entity.Pid,
		"uid":        entity.Uid,
		"ruid":       entity.Ruid,
		"content":    entity.Content,
		"is_delete":  0,
		"created_at": gtime.Now(),
		"updated_at": gtime.Now(),
	})
	if err != nil {
		return err
	}
	if id, err := result.LastInsertId(); err != nil || id <= 0 {
		return gerror.New("评论失败")
	}
	return nil
}
