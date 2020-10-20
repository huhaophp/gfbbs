package admin

import (
	"bbs/app/model/nodes"
	"bbs/app/model/users"
	"bbs/app/request/admin"
	"errors"
	"github.com/gogf/gf/frame/g"
)

type nodeService struct{}

var Node = &nodeService{}

func (n *nodeService) Add(data *admin.NodeAddReqEntity) error {
	node, _ := g.DB().Table(nodes.Table).Where("name = ?", data.Name).Where("is_delete = ?", 0).One()
	if node != nil {
		return errors.New("名称已存在")
	}
	res, err := g.DB().Table(nodes.Table).Insert(g.Map{
		"name":   data.Name,
		"sort":   data.Sort,
		"desc":   data.Desc,
		"status": data.Status,
		"is_top": data.IsTop,
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

func (n *nodeService) Edit(data *admin.NodeAddReqEntity, id int) error {
	node, _ := g.DB().Table(users.Table).Where("name = ? and id != ? and is_delete = ?", data.Name, id, 0).One()
	if node != nil {
		return errors.New("名称已存在")
	}
	gMap := g.Map{
		"name":   data.Name,
		"sort":   data.Sort,
		"desc":   data.Desc,
		"status": data.Status,
		"is_top": data.IsTop,
	}
	res, err := g.DB().Table(nodes.Table).WherePri(id).Update(gMap)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}

func (n *nodeService) Delete(id int) error {
	res, err := g.DB().Table(nodes.Table).Where("id = ? and is_delete = ?", id, 0).Update(g.Map{"is_delete": 1})
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}
