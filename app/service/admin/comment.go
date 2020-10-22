package admin

import (
	"bbs/app/model/comments"
	"github.com/gogf/gf/frame/g"
)

type commentService struct{}

var Comment = &commentService{}

func (n *commentService) Delete(id int) error {
	res, err := g.DB().Table(comments.Table).Where("id = ? and is_delete = ?", id, 0).Update(g.Map{"is_delete": 1})
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}
