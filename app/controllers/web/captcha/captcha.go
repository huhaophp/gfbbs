package captcha

import (
	"github.com/afocus/captcha"
	"github.com/gogf/gf/net/ghttp"
	"image/color"
	"image/png"
)

// Controller Base
type Controller struct{}

func (c *Controller) Get(r *ghttp.Request) {
	var ca *captcha.Captcha
	ca = captcha.New()
	_ = ca.SetFont("./public/captcha/comic.ttf")
	ca.SetSize(128, 64)
	ca.SetFrontColor(color.RGBA{201, 100, 255, 210})
	ca.SetBkgColor(color.RGBA{150,200,140, 220})
	img, str := ca.Create(4, captcha.NUM)
	_ = r.Session.Set("captcha", str)
	_ = png.Encode(r.Response.Writer, img)
}
