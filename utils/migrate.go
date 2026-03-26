package utils

import (
	"blog_backend_go/model"
	"log"

	"gorm.io/gorm"
)

// AutoMigrate 执行所有数据库表的自动迁移
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Article{},
		&model.ArticleComment{},
		&model.ArticleLike{},
		&model.ArticleTag{},
		&model.ArticleTagMap{},
		&model.SysLogDbBackup{},
		&model.SysLogReboot{},
		&model.SysLogServer{},
		&model.SysPermission{},
		&model.SysRole{},
		&model.SysRolePermission{},
		&model.SysUser{},
		&model.SysUserRole{},
		&model.WebFriendship{},
		&model.WebMessage{},
		&model.WebVisit{},
	)
	if err != nil {
		log.Fatalf("❌ 数据库自动迁移失败: %v", err)
	}
	log.Println("✅ 数据库表结构迁移/同步成功")
}
