package model

import "gorm.io/gorm"

type Lyric2Song struct {
	gorm.Model
	SongHash  string `gorm:"varchar(255);not null;primaryKey"`
	LyricHash string `gorm:"varchar(255);not null;primaryKey"`
}
