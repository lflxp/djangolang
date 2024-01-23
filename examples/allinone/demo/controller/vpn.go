package controller

import (
	"github.com/lflxp/djangolang/examples/allinone/demo/model"

	"github.com/lflxp/djangolang/utils/orm/sqlite"
)

// 根据id查询
func GetByUUIDVpn(uuid string) (*model.Vpn, bool, error) {
	data := new(model.Vpn)
	has, err := sqlite.NewOrm().Where("id = ?", uuid).Get(data)
	return data, has, err
}

// 全量数据
func ListVpn() (*[]model.Vpn, error) {
	data := new([]model.Vpn)
	err := sqlite.NewOrm().Find(data)
	return data, err
}

// 获取分页数据
func PageVpn(page, pageSize int) (*[]model.Vpn, error) {
	data := new([]model.Vpn)
	err := sqlite.NewOrm().Limit(pageSize, (page-1)*pageSize).Find(data)
	return data, err
}

// 添加数据
func AddVpn(data *model.Vpn) (int64, error) {
	affected, err := sqlite.NewOrm().Insert(data)
	return affected, err
}

// 删除数据
func DelVpn(id string) (int64, error) {
	data := new(model.Vpn)
	affected, err := sqlite.NewOrm().ID(id).Delete(data)
	return affected, err
}

// 更新数据
func UpdateVpn(id string, data *model.Vpn) (int64, error) {
	affected, err := sqlite.NewOrm().Table(new(model.Vpn)).ID(id).Update(data)
	return affected, err
}

// 批量导入
func MultiInsertVpn(beans ...interface{}) (int64, error) {
	return sqlite.NewOrm().Insert(beans...)
}
