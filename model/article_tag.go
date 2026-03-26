package model

type ArticleTag struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	TagName string `gorm:"type:varchar(45);not null" json:"tag_name"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}
