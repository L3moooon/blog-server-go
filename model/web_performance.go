package model

import "time"

type WebPerformance struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Path           string    `gorm:"type:varchar(45);not null" json:"path"`
	UserAgent      string    `gorm:"type:varchar(255);not null" json:"user_agent"`
	ScreenSize     string    `gorm:"type:varchar(45);not null" json:"screen_size"`
	NavigationType string    `gorm:"type:varchar(45);not null" json:"navigation_type"`
	IsCache        bool      `gorm:"type:tinyint;not null" json:"is_cache"`
	FCP            int64     `gorm:"type:bigint;not null" json:"fcp"`
	LCP            int64     `gorm:"type:bigint;not null" json:"lcp"`
	INP            int64     `gorm:"type:bigint;not null" json:"inp"`
	CLS            float64   `gorm:"type:decimal(10,2);not null" json:"cls"`
	FID            int64     `gorm:"type:bigint;not null" json:"fid"`
	InitTime       time.Time `gorm:"type:datetime;not null" json:"init_time"`
	ServerTime     time.Time `gorm:"type:datetime;not null" json:"server_time"`
}

func (WebPerformance) TableName() string {
	return "web_performance"
}
