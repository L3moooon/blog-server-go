package model

type SysRolePermission struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID       uint64 `gorm:"type:int(11);not null" json:"role_id"`
	PermissionID uint64 `gorm:"type:int(11);not null" json:"permission_id"`
}

func (SysRolePermission) TableName() string {
	return "sys_role_permission"
}
