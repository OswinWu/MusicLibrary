package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	SongHash string `gorm:"varchar(255);not null;unique"`
	Location string `gorm:"varchar(255);not null"`
}
