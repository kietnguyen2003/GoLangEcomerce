package po

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"column:uuid; type:varchar(225); unique; not null; auto_increment;comment:UUID"`
	UserName string `gorm:"column:user_name; type:varchar(225); unique; not null; comment:User name"`
	Role     []Role `gorm:"many2many:users_roles;"`
	IsActive bool   `gorm:"column:is_active; default:1; type:boolean"`
}

func (u *User) TableName() string {
	return "users"
}
