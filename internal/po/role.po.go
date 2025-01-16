package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; primary_key; auto_increment;comment:ID"`
	RoleName string `gorm:"column:role_name; type:varchar(225); not null; comment:Role name"`
	RoleDesc string `gorm:"column:role_desc; type:varchar(225); not null; comment:Role description"`
}

func (r *Role) TableName() string {
	return "roles"
}
