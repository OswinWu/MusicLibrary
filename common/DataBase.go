package common

import (
	"MusicLibrary/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Lyric{})
	db.AutoMigrate(&model.Song{})
	db.AutoMigrate(&model.Lyric2Song{})
	db.AutoMigrate(&model.User2Song{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
