package model

import "time"

type SysLogReboot struct {
	ID      uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Action  string    `gorm:"type:varchar(45);not null" json:"action"`
	Result  string    `gorm:"type:varchar(45);not null" json:"result"`
	Message string    `gorm:"type:varchar(255);not null" json:"message"`
	Time    time.Time `json:"time"`
}

func (SysLogReboot) TableName() string {
	return "sys_log_reboot"
}
