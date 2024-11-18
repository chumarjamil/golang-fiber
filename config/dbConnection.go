package config

import (
	"fmt"
	"os"

	models "go_lang_rest_api/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Database connection start
	godotenv.Load()
	dbname := os.Getenv("MYSQL_DBNAME")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbhost := os.Getenv("MYSQL_HOST")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
  	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("DB connection failed")
	}

	DB = db
	fmt.Println("Db connection successful!")

	AutoMigrate(db)

	// Database connection end
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Discount{},
		&models.Order{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
	)
}