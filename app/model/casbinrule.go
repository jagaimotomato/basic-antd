package model

type CasbinRule struct {
	PType string `gorm:"size:100" json:"ptype"`
	V0    string `gorm:"size:100" json:"v0"`
	V1    string `gorm:"size:100" json:"v1"`
	V2    string `gorm:"size:100" json:"v2"`
	V3    string `gorm:"size:100" json:"v3"`
	V4    string `gorm:"size:100" json:"v4"`
	V5    string `gorm:"size:100" json:"v5"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
