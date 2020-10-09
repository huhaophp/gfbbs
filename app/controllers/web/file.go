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
func (c *FileController) WangEditorFileStore(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	name, err := file.Save(uploadDirPath, true)
	if err != nil {
		_ = r.Response.WriteJsonExit(g.Map{"errno": 500, "msg": err.Error()})
	} else {
		_ = r.Response.WriteJsonExit(g.Map{
			"errno": 0,
			"msg":   "上传成功",
			"data":  g.Slice{"http://127.0.0.1:8199/uploadfile/" + name},
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
			"name": fmt.Sprintf("/uploadfile/%s", name),
		})
	}
}
