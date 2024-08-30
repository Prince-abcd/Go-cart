package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectdatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}
	err = DB.AutoMigrate(&Item{}, &User{}, &Cart{}, &Order{})
	if err != nil {
		log.Fatal("Failed to migrate the database!", err)
	}
	fmt.Println("Database connected ...")
}
