package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Mail     string `gorm:"varchar(255);not null;unique"`
	Uid      string `gorm:"varchar(255);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
