package service

import (
	"bbs/app/model/admins"
	"bbs/app/model/users"
	"errors"
	"fmt"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type RegisterReqEntity struct {
	Name     string `p:"name" v:"required#请输入用户昵称"`
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Password string `p:"password" v:"required|password#请输入密码|请输入长度6~18位字符的密码"`
	Captcha  string `p:"captcha" v:"required#验证码错误"`
}

type LoginReqEntity struct {
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Password string `p:"password" v:"required|password#请输入密码|请输入长度6~18位字符的密码"`
}

type UpdateInfoEntity struct {
	Name   string `p:"name" v:"required#请输入用户昵称"`
	Email  string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Gender string `p:"gender" v:"required#请输入正确的性别"`
	Sign   string `p:"sign"`
	Site   string `p:"site"`
}

type UpdateAvatarEntity struct {
	Avatar string `p:"avatar" v:"required#请上传用户头像"`
}

type UpdatePasswordEntity struct {
	Password        string `p:"password" v:"required|password#请输入新密码|新密码长度在6~18之间"`
	ConfirmPassword string `p:"confirm_password" v:"required|password#请输入确认密码|确认密码长度在6~18之间"`
}

var UserService = newUserService()

// nodeService Initialize the service
func newUserService() *userService {
	return &userService{}
}

// userService
type userService struct{}

// Register 用户注册
func (s *userService) Register(entity *RegisterReqEntity) error {
	if CheckName(entity.Name) != nil {
		return fmt.Errorf("昵称 %s 已经存在", entity.Name)
	}
	if CheckEmail(entity.Email) != nil {
		return fmt.Errorf("邮箱 %s 已经存在", entity.Email)
	}
	password, _ := gmd5.Encrypt(entity.Password)
	res, err := g.DB().Table(users.Table).Insert(g.Map{
		"name":        entity.Name,
		"email":       entity.Email,
		"password":    password,
		"status":      0,
		"gender":      0,
		"sign":        "这个人什么都没有留下",
		"register_at": gtime.Now(),
		"create_at":   gtime.Now(),
		"update_at":   gtime.Now(),
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
func (s *userService) Login(entity *LoginReqEntity) (gdb.Record, error) {
	user := CheckEmail(entity.Email)
	if user == nil {
		return nil, errors.New("邮箱或密码错误")
	}
	// 用户是否被禁用
	if user["status"].Int() == users.ForbiddenStatus {
		return nil, errors.New("用户已被禁用")
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

// UpdateInfo 更新用户基础信息
func (s *userService) UpdateInfo(id string, entity *UpdateInfoEntity) error {
	data := g.Map{
		"name":   entity.Name,
		"email":  entity.Email,
		"gender": entity.Gender,
		"sign":   entity.Sign,
		"site":   entity.Site,
	}
	res, err := g.DB().Table(users.Table).WherePri(id).Update(data)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil || row <= 0 {
		g.Log().Error("Failed to update user")
		return gerror.New("Failed to update user")
	}
	return nil
}

// UpdateAvatar 更新头像
func (s *userService) UpdateAvatar(id string, entity *UpdateAvatarEntity) error {
	res, err := g.DB().Table(users.Table).WherePri(id).Update(g.Map{
		"avatar": entity.Avatar,
	})
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil || row <= 0 {
		g.Log().Error("Failed to update user")
		return gerror.New("Failed to update user")
	}
	return nil
}

// UpdatePassword 更新用户密码
func (s *userService) UpdatePassword(id string, entity *UpdatePasswordEntity) error {
	if entity.Password != entity.ConfirmPassword {
		return gerror.New("新密码和确认输入不一致")
	}
	password, _ := gmd5.Encrypt(entity.Password)
	res, err := g.DB().Table(users.Table).WherePri(id).Update(g.Map{
		"password": password,
	})
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil || row <= 0 {
		return gerror.New("修改密码失败")
	}
	return nil
}

// CheckUserStatus 检查用户状态
func (a *userService) CheckUserStatus(uid int) gdb.Record {
	r, _ := g.DB().Table(users.Table).WherePri(uid).Where("status", admins.NormalStatus).One()
	return r
}
