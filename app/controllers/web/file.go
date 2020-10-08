package web

import (
	"bbs/app/funcs/response"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	uploadDirPath string = "./public/uploadfile/"
)

// Controller Base
type FileController struct{}

// Markdown File Store Upload uploads files to /tmp .
func (c *FileController) MdFileStore(r *ghttp.Request) {
	file := r.GetUploadFile("editormd-image-file")
	name, err := file.Save(uploadDirPath, true)
	if err != nil {
		_ = r.Response.WriteJsonExit(g.Map{"success": 0, "message": err.Error()})
	} else {
		_ = r.Response.WriteJsonExit(g.Map{
			"success": 1,
			"path":    "/uploadfile/" + name,
			"message": "上传成功",
			"url":     "/uploadfile/" + name,
		})
	}
}

// Upload uploads files to /tmp .
func (c *FileController) FileStore(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	name, err := file.Save(uploadDirPath, true)
	if err != nil {
		response.Json(r, 0, "上传失败")
	} else {
		response.Json(r, 1, "上传成功", g.Map{
			"name":   fmt.Sprintf("/uploadfile/%s", name),
		})
	}
}
