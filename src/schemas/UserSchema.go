package schemas

import "gorm.io/gorm"

type User struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255);"`
	Password string `gorm:"type:varchar(255); null null;" form:"password"`
	Email    string `gorm:"type:varchar(255); not null;" form:"email"`
	gorm.Model
}
