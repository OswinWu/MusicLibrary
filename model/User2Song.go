package model

import "gorm.io/gorm"

type User2Song struct {
	gorm.Model
	SongHash string `gorm:"varchar(255);not null;primaryKey"`
	Uid      string `gorm:"varchar(255);not null;primaryKey"`
}
