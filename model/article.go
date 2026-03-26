package model

import (
	"time"
)

type Article struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title        string    `gorm:"type:varchar(45);not null" json:"title"`
	CoverImg     string    `gorm:"type:varchar(255)" json:"cover_img"`
	Abstract     string    `gorm:"type:varchar(255)" json:"abstract"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	PublishDate  time.Time `json:"publish_date"`
	LastEditDate time.Time `json:"last_edit_date"`
	AuthorName   string    `gorm:"type:varchar(45);not null" json:"author_name"`
	View         int64     `gorm:"type:bigint;not null" json:"view"`
	Like         int64     `gorm:"type:bigint;not null" json:"like"`
	Status       bool      `gorm:"type:tinyint;not null" json:"status"`
	Tag          []int64   `gorm:"type:json;not null" json:"tag"`
	Top          bool      `gorm:"type:tinyint;not null" json:"top"`
}

func (Article) TableName() string {
	return "article"
}
