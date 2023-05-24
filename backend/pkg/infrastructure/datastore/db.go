package datastore

import (
	"ReserveMate/backend/pkg/domain/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// dsn := fmt.Sprintf("host=%s user=admin password=admin2023 dbname=reserve port=5432 sslmode=disable TimeZone=Asia/Bangkok", config.C.Database.Host)
	dsn := "host=localhost user=admin password=admin2023 dbname=reserve port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Merchant{},
	)
}
