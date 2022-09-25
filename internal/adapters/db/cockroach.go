package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := "postgresql://ibrahim:Q8teKdOtIr3CubzlCsnppA@free-tier13.aws-eu-central-1.cockroachlabs.cloud:26257/gochallenge?sslmode=verify-full&options=--cluster%3Dwoozy-dryad-2598"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)
	fmt.Println(now)
	return db, err
}
