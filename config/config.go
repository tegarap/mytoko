package config

import (
	"github.com/labstack/gommon/log"
	"github.com/tegarap/mytoko/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Dialector{
		DSN:        "tokobaru",
	}, &gorm.Config{})

	if err != nil {
		log.Info("Fail to connect database: ", err)
		panic(err)
	}

	DBMigrate(db)

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		models.Product{},
		models.Admin{},
		models.Cashier{},
	)
}
