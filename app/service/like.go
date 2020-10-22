package service

import (
	"bbs/app/model/likes"
	"bbs/app/model/posts"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type LikeReqEntity struct {
	Uid     int    `p:"uid" v:"required#点赞用户错误"`
	Tid     int    `p:"tid" v:"required#点赞目标错误"`
	TidType string `p:"tid_type" v:"required#点赞目标类型错误错误"`
}

var LikeService = newLikeService()

// newLikeService Initialize the service
func newLikeService() *likeService {
	return &likeService{}
}

// likeService
type likeService struct{}

// IsDo
func (s *likeService) IsDo(uid int, tid int, tidType string) int {
	res, err := g.DB().Table(likes.Table).Where(g.Map{"uid": uid, "tid": tid, "type": tidType}).One()
	if err != nil {
		g.Log().Error(err)
	}
	if res.IsEmpty() {
		return -1
	}
	return res["status"].Int()
}

// Do
func (s *likeService) Do(req *LikeReqEntity) error {
	res, err := g.DB().Table(likes.Table).Where(g.Map{"uid": req.Uid, "tid": req.Tid, "type": req.TidType}).One()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	if !res.IsEmpty() {
		if res["status"].Int() == 1 {
			return gerror.New("已点赞")
		} else {
			res, err := g.DB().Table(likes.Table).
				Data(g.Map{"status": likes.Do}).
				WherePri("id = ?", res["id"].Int()).
				Update()
			if err != nil {
				return err
			}
			rows, err := res.RowsAffected()
			if err != nil {
				return err
			}
			if rows <= 0 {
				return gerror.New("点赞失败")
			}
		}
	} else {
		res, err := g.DB().Table(likes.Table).Data(g.Map{
			"uid":    req.Uid,
			"tid":    req.Tid,
			"type":   req.TidType,
			"status": likes.Do,
		}).Insert()
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
			return gerror.New("点赞失败")
		}
	}
	_, _ = g.DB().Table(posts.Table).Data("like_num = like_num + 1").WherePri(req.Tid).Update()
	return nil
}

// Undo
func (s *likeService) Undo(req *LikeReqEntity) error {
	res, err := g.DB().Table(likes.Table).Where(g.Map{"uid": req.Uid, "tid": req.Tid, "type": req.TidType}).One()
	if err != nil {
		g.Log().Error(err)
		return err
	}
	if res.IsEmpty() || res["status"].Int() == 0 {
		return gerror.New("未点赞")
	} else {
		res, err := g.DB().Table(likes.Table).
			Data(g.Map{"status": likes.Undo}).
			WherePri("id = ?", res["id"].Int()).
			Update()
		if err != nil {
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rows > 0 {
			_, _ = g.DB().Table(posts.Table).Data("like_num = like_num - 1").WherePri(req.Tid).Update()
			return nil
		} else {
			return gerror.New("取消失败")
		}

	}
}

// GetTheLatestLikes
func (s *likeService) GetTheLatestLikes(tid int, tidtype string, limit int) gdb.Result {
	res, err := g.DB().Table(likes.Table+" l").
		LeftJoin("users u", "u.id = l.uid").
		Where(g.Map{"l.tid": tid, "l.type": tidtype, "l.status": likes.Do}).
		Fields("u.name,u.avatar,u.id").
		Order("l.id DESC").
		Limit(limit).
		All()
	if err != nil {
		g.Log().Error(err)
	}
	return res
}
