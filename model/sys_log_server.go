package model

import "time"

type SysLogServer struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CpuUsage      float64   `gorm:"type:decimal(5,2);not null" json:"cpu_usage"`
	MemoryUsage   float64   `gorm:"type:decimal(5,2);not null" json:"mem_usage"`
	DiskUsage     float64   `gorm:"type:decimal(5,2);not null" json:"disk_usage"`
	NetworkStatus bool      `gorm:"type:tinyint;not null" json:"network_status"`
	Time          time.Time `json:"time"`
}

func (SysLogServer) TableName() string {
	return "sys_log_server"
}
