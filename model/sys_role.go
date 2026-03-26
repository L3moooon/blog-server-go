package model

import "time"

type SysRole struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName    string    `gorm:"type:varchar(45);not null" json:"role_name"`
	RoleCode    string    `gorm:"type:varchar(45);not null" json:"role_code"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime;not null" json:"create_time"`
	UpdateTime  time.Time `gorm:"type:datetime;not null" json:"update_time"`
	Status      bool      `gorm:"type:tinyint;not null" json:"status"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

