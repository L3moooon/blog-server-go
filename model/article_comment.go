package model

import "time"

type ArticleComment struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ArticleID   uint64    `gorm:"type:bigint;not null" json:"article_id"`
	UserID      uint64    `gorm:"type:bigint;not null" json:"user_id"`
	ParentID    uint64    `gorm:"type:bigint;not null" json:"parent_id"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	CommentDate time.Time `json:"comment_date"`
	EditDate    time.Time `json:"edit_date"`
	Status      bool      `gorm:"type:tinyint;not null" json:"status"`
	LikeCount   int64     `gorm:"type:bigint;not null" json:"like_count"`
}

func (ArticleComment) TableName() string {
	return "article_comment"
}
