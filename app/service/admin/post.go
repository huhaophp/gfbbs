package admin

import (
	commentsModel "bbs/app/model/comments"
	"bbs/app/model/posts"
	"bbs/app/request/admin"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type postService struct{}

var Post = &postService{}

func (n *postService) Add(data *admin.PostAddReqEntity) error {
	res, err := g.DB().Table(posts.Table).Insert(g.Map{
		"title":   data.Title,
		"uid":     data.Uid,
		"nid":     data.Nid,
		"content": data.Content,
		"fine":    data.Fine,
	})
	if err != nil {
		g.Log().Error("入库失败：", err)
		return errors.New("添加失败")
	}
	id, err := res.LastInsertId()
	if err != nil || id <= 0 {
		g.Log().Error("入库失败：", err)
		return errors.New("添加失败")
	}

	return nil
}

func (n *postService) Edit(data *admin.PostAddReqEntity, id int) error {
	gMap := g.Map{
		"title":    data.Title,
		"uid":      data.Uid,
		"nid":      data.Nid,
		"content":  data.Content,
		"view_num": data.ViewNum,
		"like_num": data.LikeNum,
		"fine":     data.Fine,
	}
	res, err := g.DB().Table(posts.Table).WherePri(id).Update(gMap)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}

func (n *postService) Delete(id int) error {
	res, err := g.DB().Table(posts.Table).WherePri(id).Delete()
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}

func (n *postService) Show(id int) (gdb.Record, gdb.Result, error) {
	post, _ := g.DB().Table(posts.Table+" p").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Fields("p.id,p.title,p.content,p.uid,p.nid,p.view_num,p.like_num,p.comment_num,p.create_at,u.name as user_name,u.avatar,u.sign,u.site,n.name node_name").
		Where("p.id = ?", id).
		One()
	if post == nil {
		return nil, nil, errors.New("文章不存在")
	}
	comments, _ := g.DB().Table(commentsModel.Table+" c").
		Fields("c.id,c.uid,c.ruid,c.content,u.name,u.avatar,ru.name as r_user_name,c.create_at").
		LeftJoin("users u", "u.id = c.uid").
		LeftJoin("users ru", "ru.id = c.ruid").
		Where("c.pid", id).
		Where("c.is_delete", 0).
		Order("id ASC").
		Limit(20).
		All()

	return post, comments, nil
}
