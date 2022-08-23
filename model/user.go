package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 用户的唯一标识
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Email    string `gorm:"column:email;type:varchar(30);" json:"email"`
}
