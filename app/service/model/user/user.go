package user

import (
	"bbs/app/model/users"
	"errors"
	"fmt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type RegisterReqEntity struct {
	Name     string `p:"name" v:"required#请输入用户昵称"`
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Password string `p:"password" v:"required|password#请输入密码|请输入长度6~18位字符的密码"`
}

type LoginReqEntity struct {
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Password string `p:"password" v:"required|password#请输入密码|请输入长度6~18位字符的密码"`
}

// Register 用户注册
func Register(entity *RegisterReqEntity) error {
	if CheckName(entity.Name) != nil {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", entity.Name))
	}
	if CheckEmail(entity.Email) != nil {
		return errors.New(fmt.Sprintf("邮箱 %s 已经存在", entity.Email))
	}
	password, _ := gmd5.Encrypt(entity.Password)
	res, err := g.DB().Table(users.Table).Insert(g.Map{
		"name":        entity.Name,
		"email":       entity.Email,
		"password":    password,
		"status":      0,
		"gender":      0,
		"register_at": gtime.Now(),
		"created_at":  gtime.Now(),
		"updated_at":  gtime.Now(),
	})
	if err != nil {
		return err
	}
	if id, err := res.LastInsertId(); err != nil || id <= 0 {
		return errors.New("用户注册失败")
	}
	return nil
}

// Login 用户登录
func Login(entity *LoginReqEntity) (gdb.Record, error) {
	user := CheckEmail(entity.Email)
	if user == nil {
		return nil, errors.New("邮箱或密码错误")
	}
	password, _ := gmd5.Encrypt(entity.Password)
	if password != user["password"].String() {
		return nil, errors.New("邮箱或密码错误")
	}
	return user, nil
}

func CheckName(name string) gdb.Record {
	record, err := g.DB().Table(users.Table).Where("name = ?", name).One()
	if err != nil {
		return nil
	}
	if record.IsEmpty() {
		return nil
	}
	return record
}

func CheckEmail(email string) gdb.Record {
	record, err := g.DB().Table(users.Table).Where("email = ?", email).One()
	if err != nil {
		return nil
	}
	if record.IsEmpty() {
		return nil
	}
	return record
}
