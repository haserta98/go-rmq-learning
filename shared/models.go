package shared

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Age      int
}

type Order struct {
	gorm.Model
	ID     uint `gorm:"primarykey"`
	User   User `gorm:"foreignkey:UserID"`
	UserID uint
	Price  float32
}
