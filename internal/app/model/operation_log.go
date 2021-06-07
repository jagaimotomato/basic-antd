package model

import (
	orm "basic-antd/init/global"
	"gorm.io/gorm"
	"time"
)

type OperationLog struct {
	gorm.Model
	By
	Title         string    `json:"title" gorm:"type:varchar(255);comment:操作模块"`
	BusinessType  string    `json:"businessType" gorm:"type:varchar(128);comment:操作类型"`
	BusinessTypes string    `json:"businessTypes" gorm:"type:varchar(128);comment:BusinessTypes"`
	Method        string    `json:"method" gorm:"type:varchar(128);comment:函数"`
	RequestMethod string    `json:"requestMethod" gorm:"type:varchar(128);comment:请求方式"`
	OperatorType  string    `json:"operatorType" gorm:"type:varchar(128);comment:操作类型"`
	OperName      string    `json:"operName" gorm:"type:varchar(128);comment:操作者"`
	OperUrl       string    `json:"operUrl" gorm:"type:varchar(255);comment:访问地址"`
	OperIp        string    `json:"operIp" gorm:"type:varchar(128);comment:客户端ip"`
	OperLocation  string    `json:"operLocation" gorm:"type:varchar(128);comment:访问位置"`
	OperParam     string    `json:"operParam" gorm:"type:varchar(255);comment:请求参数"`
	Code          string    `json:"code" gorm:"type:varchar(100);comment:状态码"`
	Status        string    `json:"status" gorm:"type:varchar(4);comment:操作状态"`
	OperTime      time.Time `json:"operTime" gorm:"type:timestamp;comment:操作时间"`
	JsonResult    string    `json:"jsonResult" gorm:"type:varchar(255);comment:返回数据"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	DurationTime  string    `json:"durationTime" gorm:"type:varchar(128);comment:耗时"`
	UserAgent     string    `json:"userAgent" gorm:"type:varchar(255);comment:ua"`
}

func (OperationLog) TableName() string {
	return "operation_log"
}

func (o *OperationLog) Insert() (err error) {
	if err = orm.Eloquent.Table(o.TableName()).Create(&o).Error; err != nil {
		return
	}
	return
}

func (o *OperationLog) GetPage(pageIndex, pageSize int) (operationLog []OperationLog, count int64, err error) {
	if err = orm.Eloquent.Table(o.TableName()).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&operationLog).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}
