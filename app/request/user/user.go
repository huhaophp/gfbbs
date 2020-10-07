package user

import (
	"github.com/gogf/gf/net/ghttp"
)

type AddReqEntity struct {
	Name    string `p:"name" v:"required|length:3,16#请输入用户名|用户名长度应当在:min到:max之间"`
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Gender	uint	`p:"gender" v:"required|in:0,1,2#请填写性别|性别只能在0,1,2之间"`
	Avatar  string	`p:"avatar"`
	Password string `p:"password" v:"required|length:6,16#请填写密码|密码长度应当在:min到:max之间"`
	Status uint `p:"status" v:"required|in:0,1#请填写状态|状态只能在0,1之间"`
	LastLoginIp	string	`p:"last_login_ip"`
	RegisterAt string	`p:"register_at"`
}

type UpdateReqEntity struct {
	Name    string `p:"name" v:"required|length:3,16#请输入用户名|用户名长度应当在:min到:max之间"`
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Gender	uint	`p:"gender" v:"required|in:0,1,2#请填写性别|性别只能在0,1,2之间"`
	Avatar  string	`p:"avatar"`
	Password string `p:"password" v:"length:6,16#密码长度应当在:min到:max之间"`
	Status uint `p:"status" v:"required|in:0,1#请填写状态|状态只能在0,1之间"`
	LastLoginIp	string	`p:"last_login_ip"`
	RegisterAt string	`p:"register_at"`
}

func AddReqCheck(r *ghttp.Request, data *AddReqEntity) error  {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}

func UpdateReqCheck(r *ghttp.Request, data *UpdateReqEntity) error  {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}
