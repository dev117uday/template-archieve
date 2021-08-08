package db

import (
	"go-postgres/conf"
	"go-postgres/logrus"
	"go-postgres/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	db, err := gorm.Open(postgres.Open(conf.PGCS), &gorm.Config{})
	if err != nil {
		logrus.LogIt(0, "ConnectTODB", "Failed to connect to DB", err)
	}
	Db = db
	db.AutoMigrate(&model.User{})
}
