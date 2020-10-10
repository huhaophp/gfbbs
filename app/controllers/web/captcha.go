package web

import (
	"github.com/afocus/captcha"
	"github.com/gogf/gf/net/ghttp"
	"image/color"
	"image/png"
)

// CaptchaController Base
type CaptchaController struct{}

func (c *CaptchaController) Get(r *ghttp.Request) {
	var ca *captcha.Captcha
	ca = captcha.New()
	_ = ca.SetFont("./public/captcha/comic.ttf")
	ca.SetSize(128, 64)
	ca.SetFrontColor(color.RGBA{211, 120, 205, 210})
	ca.SetBkgColor(color.RGBA{50,200,140, 220})
	img, str := ca.Create(4, captcha.NUM)
	_ = r.Session.Set("captcha", str)
	_ = png.Encode(r.Response.Writer, img)
}
