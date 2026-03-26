package model

type SysUserRole struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint64 `gorm:"type:bigint;not null" json:"user_id"`
	RoleID uint64 `gorm:"type:bigint;not null" json:"role_id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
