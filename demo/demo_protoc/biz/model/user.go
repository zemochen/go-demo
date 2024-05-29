package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100) not null;unique_index"`
	Password string `gorm:"type:varchar(100) not null"`
}

func (u *User) TableName() string {
	//custom table name, default is `${struct_name}s`
	return "user"
}
