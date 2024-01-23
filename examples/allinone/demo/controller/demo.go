package controller

import (
	"github.com/lflxp/djangolang/examples/allinone/demo/model"

	"github.com/lflxp/djangolang/utils/orm/sqlite"
)

// 根据id查询
func GetByUUIDDemo(uuid string) (*model.Demo, bool, error) {
	data := new(model.Demo)
	has, err := sqlite.NewOrm().Where("id = ?", uuid).Get(data)
	return data, has, err
}

// 全量数据
func ListDemo() (*[]model.Demo, error) {
	data := new([]model.Demo)
	err := sqlite.NewOrm().Find(data)
	return data, err
}

// 获取分页数据
func PageDemo(page, pageSize int) (*[]model.Demo, error) {
	data := new([]model.Demo)
	err := sqlite.NewOrm().Limit(pageSize, (page-1)*pageSize).Find(data)
	return data, err
}

// 添加数据
func AddDemo(data *model.Demo) (int64, error) {
	affected, err := sqlite.NewOrm().Insert(data)
	return affected, err
}

// 删除数据
func DelDemo(id string) (int64, error) {
	data := new(model.Demo)
	affected, err := sqlite.NewOrm().ID(id).Delete(data)
	return affected, err
}

// 更新数据
func UpdateDemo(id string, data *model.Demo) (int64, error) {
	affected, err := sqlite.NewOrm().Table(new(model.Demo)).ID(id).Update(data)
	return affected, err
}

// 批量导入
func MultiInsertDemo(beans ...interface{}) (int64, error) {
	return sqlite.NewOrm().Insert(beans...)
}
