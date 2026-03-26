package model

import "time"

type SysUser struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Account       string    `gorm:"type:varchar(45);not null" json:"account"`
	UserName      string    `gorm:"type:varchar(45);not null" json:"user_name"`
	Password      string    `gorm:"type:varchar(45);not null" json:"password"`
	Avatar        string    `gorm:"type:varchar(255)" json:"avatar"`
	Ip            string    `gorm:"type:varchar(45);not null" json:"ip"`
	Location      string    `gorm:"type:varchar(255);not null" json:"location"`
	CreateTime    time.Time `gorm:"type:datetime;not null" json:"create_time"`
	LastLoginTime time.Time `gorm:"type:datetime;not null" json:"last_login_time"`
	Status        bool      `gorm:"type:tinyint;not null" json:"status"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
