package model

type ArticleTagMap struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	ArticleID uint64 `gorm:"type:bigint;not null" json:"article_id"`
	TagID     uint64 `gorm:"type:bigint;not null" json:"tag_id"`
}

func (ArticleTagMap) TableName() string {
	return "article_tag_map"
}
