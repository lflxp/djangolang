package model

import (
	"github.com/lflxp/djangolang"
)

func init() {
	djangolang.RegisterAdmin(new(Vpn))
}

type Vpn struct {
	Id   int64  `xorm:"id notnull unique pk autoincr" name:"id"`
	Vpn  string `xorm:"vpn" name:"vpn" verbose_name:"Vpn字段测试" list:"true" search:"true"`
	Name string `xorm:"name" name:"name" verbose_name:"姓名" list:"true" search:"false"`
	Ip   string `xorm:"ip" name:"ip" verbose_name:"ip信息" list:"true" search:"false"`
}
