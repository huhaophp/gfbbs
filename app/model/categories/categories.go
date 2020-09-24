package categories

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// ListReqEntity 列表请求实体
type ListReqEntity struct {
	Name   int `p:"name"`
	Status int `p:"status"`
}

// AddReqEntity 添加文章请求实体
type AddReqEntity struct {
	Name   string `p:"name" v:"required#请输入正确的分类名称"`
	Sort   int    `p:"sort" v:"required|integer#请填写设置排序|排序必须是整数"`
	Status int    `p:"status" v:"required|in:0,1#请选择状态|状态必须在0和1之间"`
}

// List 获取栏目条目数
func List(req *ListReqEntity) (list gdb.Result, err error) {
	model := g.DB().Table(Table)
	if req.Status < -1 {
		model = model.Where("status", req.Status)
	}
	list, err = model.
		Order("sort desc").
		All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("Failed to get data.")
		return
	}
	return
}

// Add 添加栏目
func Add(data *AddReqEntity) (insId int64, err error) {
	entity := &Entity{}
	entity.Name = data.Name
	entity.Status = data.Status
	entity.Sort = data.Sort
	entity.CreatedAt = gtime.Now()
	entity.UpdatedAt = gtime.Now()
	res, err := entity.Insert()
	if err != nil {
		g.Log().Error(err)
		return
	}
	insId, err = res.LastInsertId()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}

// Edit 编辑栏目
func Edit(id int, data *AddReqEntity) (insId int64, err error) {
	entity, _ := Model.FindOne("id", id)
	entity.Name = data.Name
	entity.Sort = data.Sort
	entity.Status = data.Status
	entity.UpdatedAt = gtime.Now()
	res, err := Model.Replace(entity)
	if err != nil {
		g.Log().Error(err)
		return
	}
	insId, err = res.LastInsertId()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}

// Del 删除栏目
func Del(id int) (rows int64, err error) {
	res, err := g.DB().Table(Table).Where("id", id).Delete()
	if err != nil {
		g.Log().Error(err)
		return
	}
	rows, err = res.RowsAffected()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}

// Find 根据主键获取数据
func Find(id int) (item gdb.Record, err error) {
	item, err = g.DB().Table(Table).Where("id", id).One()
	return
}

// DoesTheNameExist 名称是否存在
func DoesTheNameExist(name string, ignoreId int) (exist bool) {
	model := g.DB().Table(Table).Where("name", name)
	if ignoreId != 0 {
		model = model.Where("id !=", ignoreId)
	}
	item, _ := model.One()
	return item != nil
}
