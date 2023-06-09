package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/miceremwirigi/go-sales-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	dbhost := os.Getenv("PSQL_HOST")
	dbuser := os.Getenv("PSQL_USER")
	dbpassword := os.Getenv("PSQL_PASSWORD")
	dbname := os.Getenv("PSQL_DBNAME")
	dbport := os.Getenv("PSQL_PORT")
	dbsslmode := os.Getenv("PSQL_SSL_MODE")
	dbtimezone := os.Getenv("PSQL_TIMEZONE")

	fmt.Println("connecting to db...")
	dsn := fmt.Sprintf("host = %s user = %s password = %s dbname= %s port = %s sslmode= %s timezone = %s", dbhost, dbuser, dbpassword, dbname, dbport, dbsslmode, dbtimezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("db connetion failed")
		// return
	}

	DB = db // connection variable
	fmt.Println("db connected successfully")
	// AutoMigrate(DB) //pass connection into automigrate function
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)
}
