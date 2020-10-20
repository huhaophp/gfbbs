package admin

import (
	"github.com/gogf/gf/net/ghttp"
)

type PostAddReqEntity struct {
	Title   string `p:"title" v:"required#请填写帖子标题"`
	Uid     string `p:"uid" v:"required#请选择用户"`
	Nid     string `p:"nid" v:"required#请选择节点"`
	Content string `p:"content" v:"required#请填写帖子内容"`
	ViewNum int `p:"view_num"`
	CommentNum int `p:"comment_num"`
	LikeNum int `p:"like_num"`
	Fine int `p:"fine" v:"required#请填写帖子是否加精|in:0,1#加精状态错误"`
}

func PostAddReqCheck(r *ghttp.Request, data *PostAddReqEntity) error {
	if err := r.Parse(data); err != nil {
		return err
	}
	return nil
}
