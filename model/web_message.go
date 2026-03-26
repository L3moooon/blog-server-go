package model

import "time"

type WebMessage struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint64    `gorm:"type:bigint;not null" json:"user_id"`
	Content    string    `gorm:"type:varchar(255);not null" json:"content"`
	CreateTime time.Time `gorm:"type:datetime;not null" json:"create_time"`
	Status     bool      `gorm:"type:tinyint;not null" json:"status"`
}

func (WebMessage) TableName() string {
	return "web_message"
}
