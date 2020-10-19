package admin

import (
	"bbs/app/model/users"
	"bbs/app/request/admin"
	"errors"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

func UserAdd(data *admin.UserAddReqEntity) error {
	user, _ := g.DB().Table(users.Table).Where("email = ?", data.Email).One()
	if user != nil {
		return errors.New("邮箱已存在")
	}
	password, _ := gmd5.Encrypt(data.Password)
	res, err := g.DB().Table(users.Table).Insert(g.Map{
		"name":        data.Name,
		"email":       data.Email,
		"gender":      data.Gender,
		"avatar":      data.Avatar,
		"password":    password,
		"status":      data.Status,
		"register_at": gtime.Now(),
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

func UserEdit(data *admin.UserUpdateReqEntity, id int) error {
	user, _ := g.DB().Table(users.Table).Where("email = ? and id != ?", data.Email, id).One()
	if user != nil {
		return errors.New("邮箱已存在")
	}
	gMap := g.Map{
		"name":        data.Name,
		"email":       data.Email,
		"gender":      data.Gender,
		"avatar":      data.Avatar,
		"status":      data.Status,
	}
	if data.Password != "" {
		password, _ := gmd5.Encrypt(data.Password)
		gMap["password"] = password
	}
	res, err := g.DB().Table(users.Table).WherePri(id).Update(gMap)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}

func UserDelete(id int) error {
	res, err := g.DB().Table(users.Table).Delete("id", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}
