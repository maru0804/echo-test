package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id int64 `gorm:"primary_key;AUTO_INCREMENT;unique;not null" binding:"required"`
}
