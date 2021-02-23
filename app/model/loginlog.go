package model

import (
	orm "basic-antd/init/global"
	"gorm.io/gorm"
	"time"
)

type LoginLog struct {
	gorm.Model
	By
	Username  string    `gorm:"size:128" json:"username"`
	Status    string    `gorm:"size:2" json:"status"`
	Mobile    string    `gorm:"size:2" json:"mobile"`
	Ip        string    `gorm:"size:64" json:"ip"`
	Location  string    `gorm:"size:255" json:"location"`
	Browser   string    `gorm:"size:128" json:"browser"`
	Os        string    `gorm:"size:128" json:"os"`
	Platform  string    `gorm:"size:128" json:"platform"`
	LoginTime time.Time `gorm:"type:timestamp" json:"loginTime"`
	Remark    string    `gorm:"size:255" json:"remark"`
	Msg       string    `gorm:"size:255" json:"msg"`
}

func (LoginLog) TableName() string {
	return "login_log"
}

func (l *LoginLog) Insert() (err error) {
	if err = orm.Eloquent.Table(l.TableName()).Create(&l).Error; err != nil {
		return
	}
	return
}

func (l *LoginLog) GetPage(pageIndex, pageSize int) (loginLog []LoginLog, count int64, err error) {
	if err = orm.Eloquent.Table(l.TableName()).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&loginLog).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}
