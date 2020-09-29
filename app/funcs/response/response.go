package response

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	_ = r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}

// 返回视图并退出当前HTTP执行函数。
func ViewExit(r *ghttp.Request, layout string, data g.Map) {
	defer func() {
		if r.Session.Get("error") != nil {
			_ = r.Session.Remove("error")
		}
		if r.Session.Get("success") != nil {
			_ = r.Session.Remove("success")
		}
	}()
	_ = r.Response.WriteTpl(layout, data)
	r.Exit()
}

// 返回404视图并退出当前HTTP执行函数。
func NotFoundView(r *ghttp.Request) {
	s := strings.Split(r.RequestURI, "/")

	d := "web"

	if s[1] == "admin" {
		d = "admin"
	}

	layout := fmt.Sprintf("%s/layout.html", d)

	data := g.Map{
		"mainTpl": fmt.Sprintf("%s/error.html", d),
		"error":   "糟糕页面找不到了",
	}

	ViewExit(r, layout, data)
}

func RedirectBackWithError(r *ghttp.Request, err error) {
	_ = r.Session.Set("error", err.Error())
	r.Response.RedirectBack()
}

func RedirectToWithError(r *ghttp.Request, to string, err error) {
	_ = r.Session.Set("error", err.Error())
	r.Response.RedirectTo(to)
}

func RedirectToWithMessage(r *ghttp.Request, to string, msg string) {
	_ = r.Session.Set("success", msg)
	r.Response.RedirectTo(to)
}
