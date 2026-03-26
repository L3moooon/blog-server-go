package model

import "time"

type WebFriendship struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(45);not null" json:"name"`
	Url         string    `gorm:"type:varchar(255);not null" json:"url"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Avatar      string    `gorm:"type:varchar(255);not null" json:"avatar"`
	Email       string    `gorm:"type:varchar(45);not null" json:"email"`
	ApplyTime   time.Time `gorm:"type:datetime;not null" json:"apply_time"`
	PassTime    time.Time `gorm:"type:datetime;not null" json:"pass_time"`
	Status      bool      `gorm:"type:tinyint;not null" json:"status"`
}

func (WebFriendship) TableName() string {
	return "web_friendship"
}
