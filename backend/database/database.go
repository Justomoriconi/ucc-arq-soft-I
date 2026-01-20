package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"ucc-arq-soft-I/model/hotel"
	"ucc-arq-soft-I/model/reservation"
	"ucc-arq-soft-I/model/room"
	"ucc-arq-soft-I/model/user"
)

var DB *gorm.DB

func Connect() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" {
		dbHost = "db"
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	gormLogger := logger.Default.LogMode(logger.Info)

	var err error
	for i := 1; i <= 30; i++ { // 30 intentos
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err == nil {
			log.Println("✅ Connected to MySQL with GORM")
			return
		}
		log.Printf("⏳ MySQL not ready (try %d/30): %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("❌ Error connecting to MySQL after retries: %v", err)
}

func Migrate() {
	if DB == nil {
		log.Fatal("DB is nil: First call database.Connect()")
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("migrate user: %v", err)
	}

	if err := DB.AutoMigrate(&hotel.Hotel{}); err != nil {
		log.Fatalf("migrate hotel: %v", err)
	}

	if err := DB.AutoMigrate(&room.Room{}); err != nil {
		log.Fatalf("migrate room: %v", err)
	}

	if err := DB.AutoMigrate(&reservation.Reservation{}); err != nil {
		log.Fatalf("migrate reservation: %v", err)
	}

	log.Println("✅ Migrations executed (AutoMigrate)")
}

func Close() {
	if DB == nil {
		return
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return
	}
	_ = sqlDB.Close()

}
