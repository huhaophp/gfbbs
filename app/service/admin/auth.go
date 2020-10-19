package admin

import (
	"bbs/app/model/admins"
	"bbs/app/request/admin"
	"errors"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

func Login(data *admin.LoginReqEntity) (gdb.Record, error) {
	res, err := g.DB().Table(admins.Table).Where("email = ?", data.Email).One()
	if res == nil || err != nil {
		return nil, errors.New("账号不存在")
	}
	hash, _ := gmd5.Encrypt(data.Password)
	if hash != (res["password"].String()) {
		return nil, errors.New("账号或者密码错误")
	}
	if res["status"].Int() == admins.ForbiddenStatus {
		return nil, errors.New("该账号已被冻结")
	}

	return res, nil
}
