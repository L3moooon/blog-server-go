package model

import "time"

type ArticleLike struct {
	UserID     uint64    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	TargetID   uint64    `gorm:"type:int(11);not null" json:"target_id"`
	TargetType string    `gorm:"type:varchar(45);not null" json:"target_type"`
	Time       time.Time `json:"like_date"`
}

func (ArticleLike) TableName() string {
	return "article_like"
}
