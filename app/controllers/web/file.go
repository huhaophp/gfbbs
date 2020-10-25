package web

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	uploadDirPath string = "./public/uploadfile/"
)

const (
	// 图片文件大小限制单位(M)
	imageFileSizeLimit = 5
	// 支持的图片格式
	supportedImageFormat = "png,jpeg,jpg"
)

// FileController Base
type FileController struct{}

// Upload 文件上传
func (c *FileController) Upload(r *ghttp.Request) {

	file := r.GetUploadFile("file")
	if file.Size > (imageFileSizeLimit * 1024 * 1024) {
		_ = r.Response.WriteJsonExit(g.Map{"errno": 500, "msg": fmt.Sprintf("图片大小不能超过%dM", imageFileSizeLimit)})
	}

	names := strings.Split(file.Filename, ".")
	if !strings.Contains(supportedImageFormat, names[1]) {
		_ = r.Response.WriteJsonExit(g.Map{"errno": 500, "msg": fmt.Sprintf("仅支持%s格式图片", supportedImageFormat)})
	}

	name, err := file.Save(uploadDirPath, true)
	if err != nil {
		_ = r.Response.WriteJsonExit(g.Map{"errno": 500, "msg": err.Error()})
	}

	_ = r.Response.WriteJsonExit(g.Map{
		"errno":    0,
		"msg":      "上传成功",
		"data":     []string{g.Cfg().GetString("server.APPURL") + "/uploadfile/" + name},
		"filename": "/uploadfile/" + name,
	})
}
