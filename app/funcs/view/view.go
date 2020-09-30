package view

import (
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// 视图模版全局函数

// StrTime 格式化时间
func StrTime(str string) string {
	// Current timestamp
	n := gtime.Now().Timestamp()
	// Given timestamp
	t := gtime.ParseTimeFromContent(str).Timestamp()
	// Timestamp difference
	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1

	d := n - t
	rs := ""
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d秒前", int(d/ss))
	default:
		rs = "刚刚"
	}

	return rs
}

// AlertComponent 警告提示框
func AlertComponent(msg interface{}, level string) string {
	c := gconv.String(msg)
	if c == "" {
		return c
	}
	s := fmt.Sprintf("<div class='alert alert-%s alert-dismissible fade show' role='alert'>", level)
	s += fmt.Sprintf("<span>%s</span>", c)
	s += "<button type='button' class='close' data-dismiss='alert' aria-label='Close'>"
	s += "<span aria-hidden=''>&times;</span>"
	s += "</button>"
	s += "</div>"
	return s
}
