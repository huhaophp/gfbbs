package admin

import (
	"github.com/gogf/gf/net/ghttp"
)

type AdminAddReqEntity struct {
	Name    string `p:"name" v:"required|length:3,16#请输入用户名|用户名长度应当在:min到:max之间"`
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Password string `p:"password" v:"required|length:6,16#请填写密码|密码长度应当在:min到:max之间"`
	Status string `p:"status" v:"required|in:0,1#请填写状态|状态只能在0,1之间"`
}

type AdminUpdateReqEntity struct {
	Name    string `p:"name" v:"required|length:3,16#请输入用户名|用户名长度应当在:min到:max之间"`
	Email    string `p:"email" v:"required|email#请输入邮箱|请输入正确的邮箱"`
	Password string `p:"password" v:"length:6,16#|密码长度应当在:min到:max之间"`
	Status string `p:"status" v:"required|in:0,1#请填写状态|状态只能在0,1之间"`
}

func AdminAddReqCheck(r *ghttp.Request, data *AdminAddReqEntity) error  {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}

func AdminUpdateReqCheck(r *ghttp.Request, data *AdminUpdateReqEntity) error  {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}
