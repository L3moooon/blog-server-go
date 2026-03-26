package model

import "time"

type WebVisit struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Identify      string    `gorm:"type:varchar(45);not null" json:"identify"`
	Name          string    `gorm:"type:varchar(45);not null" json:"name"`
	Avatar        string    `gorm:"type:varchar(255)" json:"avatar"`
	Ip            string    `gorm:"type:varchar(45);not null" json:"ip"`
	Location      string    `gorm:"type:varchar(255);not null" json:"location"`
	Agent         string    `gorm:"type:varchar(255);not null" json:"agent"`
	CreateTime    time.Time `gorm:"type:datetime;not null" json:"create_time"`
	LastLoginTime time.Time `gorm:"type:datetime;not null" json:"last_login_time"`
	VisitedCount  int64     `gorm:"type:bigint;not null" json:"visited_count"`
}

func (WebVisit) TableName() string {
	return "web_visit"
}
