package db

import (
	"fmt"
	"os"

	"github.com/HungTP-Play/hunterxhunter/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDbDns() string {
	return "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
}

func ConnectAndMigrate() {
	dbDns := getDbDns()
	fmt.Println("Connecting to database: " + dbDns)
	db, err := gorm.Open(postgres.Open(dbDns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	err = db.AutoMigrate(&models.Character{}, &models.Image{}, &models.Arc{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrated database")
}

func GetDB() *gorm.DB {
	return db
}
