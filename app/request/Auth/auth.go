package Auth

import (
	"github.com/gogf/gf/net/ghttp"
)

type LoginReqEntity struct {
	Email    string `p:"email" v:"required|length:6,16#请输入正确的邮箱|邮箱长度应当在:min到:max之间"`
	Password string `p:"password" v:"required|length:6,16#请填写密码|密码长度应当在:min到:max之间"`
}

func LoginReqCheck(r *ghttp.Request, data *LoginReqEntity) error  {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}
