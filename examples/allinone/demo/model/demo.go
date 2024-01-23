package model

import (
	"time"

	"github.com/lflxp/djangolang"
)

func init() {
	djangolang.RegisterAdmin(new(Demo))
}

type Demo struct {
	Id         int64     `xorm:"id pk not null autoincr" name:"id" search:"true"`
	Country    string    `json:"country" xorm:"varchar(255) not null" search:"true"`
	Zoom       string    `json:"zoom" xorm:"varchar(255) not null"`
	Company    string    `json:"company" xorm:"varchar(255) not null"`
	Items      string    `json:"items" xorm:"varchar(255) not null"`
	Production string    `json:"production" xorm:"varchar(255) not null"`
	Count      string    `json:"count" xorm:"varchar(255) not null"`
	Serial     string    `json:"serial" xorm:"varchar(255) not null" search:"true"`
	Extend     string    `json:"extend" xorm:"varchar(255) not null"`
	Files      string    `xorm:"file" name:"file" verbose_name:"上传文件" colType:"file"`
	File2      string    `xorm:"file2" name:"file2" verbose_name:"上传文件2" colType:"file"`
	Type       string    `xorm:"type" name:"type" verbose_name:"类型" search:"false" colType:"textarea"`
	Detail     string    `xorm:"detail" name:"detail" verbose_name:"VPN信息" list:"false" search:"false" o2m:"vpn|id,vpn" colType:"o2m"`
	Times      time.Time `xorm:"times" name:"times" verbose_name:"时间" colType:"time" list:"true" search:"true"`
}
