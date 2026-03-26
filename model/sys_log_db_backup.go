package model

import "time"

type SysLogDbBackup struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	FileName     string    `gorm:"type:varchar(255);not null" json:"file_name"`
	FileSize     int64     `gorm:"type:varchar(45)" json:"file_size"`
	Status       bool      `gorm:"type:tinyint" json:"status"`
	OssUrl       string    `gorm:"type:varchar(255)" json:"oss_url"`
	ErrorMessage string    `gorm:"type:varchar(255)" json:"error_msg"`
	Duration     int64     `gorm:"type:varchar(45)" json:"duration"`
	Time         time.Time `json:"time"`
}

func (SysLogDbBackup) TableName() string {
	return "sys_log_db_backup"
}
