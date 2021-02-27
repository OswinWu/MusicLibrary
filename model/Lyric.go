package model

import "gorm.io/gorm"

type Lyric struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	LyricHash string `gorm:"varchar(255);not null;unique"`
	SongHash  string `gorm:"varchar(255);;not null"`
}
