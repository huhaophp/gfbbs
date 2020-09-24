package articles

import (
	"bbs/app/model/categories"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// pageSize 默认条目数
const pageSize = 20

// ListReqEntity 列表请求实体
type ListReqEntity struct {
	CategoryId int    `p:"category_id"`
	PageNum    int    `p:"page"`
	PageSize   int    `p:"page_size"`
	Keywords   string `p:"keywords"`
}

// AddReqEntity 添加文章请求实体
type AddReqEntity struct {
	Title      string `p:"title" v:"required#请填写文章标题"`
	Summary    string `p:"summary"`
	Cover      string `p:"cover" v:"required#请上传文章封面"`
	CategoryId int    `p:"category_id" v:"required|min:1#请选择栏目|请选择栏目"`
	Tags       string `p:"tags" v:"required#请选择标签"`
	From       int    `p:"from" v:"required#请选择文章来源"`
	Status     int    `p:"status" v:"required#请设置文章状态"`
	Content    string `p:"content" v:"required#请填写文章内容"`
	MdContent  string `p:"md_content" v:"required#Markdown内容错误"`
}

// List 获取文章条目数
func List(req *ListReqEntity) (total int, list gdb.Result, err error) {
	model := g.DB().Table(Table + " a")
	if req.CategoryId != 0 {
		model = model.Where("a.category_id", req.CategoryId)
	}
	if req.Keywords != "" {
		model = model.Where("a.title LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	total, err = model.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("Failed to get the total number of rows.")
		return
	}
	if req.PageSize == 0 {
		req.PageSize = pageSize
	}
	list, err = model.
		InnerJoin(categories.Table+" c", "a.category_id=c.id").
		Fields("a.*,c.name category_name").
		Page(req.PageNum, req.PageSize).
		Order("a.created_at desc").
		All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("Failed to get data.")
		return
	}
	return
}

// Find 根据主键获取数据
func Find(id int) (item gdb.Record, err error) {
	item, err = g.DB().Table(Table+" a").
		Where("a.id", id).
		InnerJoin(categories.Table+" c", "a.category_id=c.id").
		Fields("a.*,c.name category_name").
		One()
	return
}

// Add 添加文章
func Add(data *AddReqEntity) (insId int64) {
	entity := &Entity{}
	entity.Title = data.Title
	entity.Summary = data.Summary
	entity.CategoryId = data.CategoryId
	entity.Cover = data.Cover
	entity.Tags = data.Tags
	entity.From = data.From
	entity.Status = data.Status
	entity.Content = data.Content
	entity.MdContent = data.MdContent
	entity.CreatedAt = gtime.Now()
	entity.UpdatedAt = gtime.Now()
	res, err := entity.Insert()
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	insId, err = res.LastInsertId()
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return insId
}

// Edit 编辑栏目
func Edit(id int, data *AddReqEntity) (rows int64) {
	entity, _ := Model.FindOne("id", id)
	entity.Title = data.Title
	entity.Summary = data.Summary
	entity.CategoryId = data.CategoryId
	entity.Cover = data.Cover
	entity.Tags = data.Tags
	entity.From = data.From
	entity.Status = data.Status
	entity.Content = data.Content
	entity.MdContent = data.MdContent
	entity.UpdatedAt = gtime.Now()
	res, err := Model.Replace(entity)
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	rows, err = res.RowsAffected()
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return rows
}

// Del 删除栏目
func Del(id int) (rows int64) {
	res, err := g.DB().Table(Table).Where("id", id).Delete()
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	rows, err = res.RowsAffected()
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return rows
}
