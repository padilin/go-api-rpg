package charDB

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(dbPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schemas
	err = DB.AutoMigrate(
		&Character{},
		&Stats{},
		&Currency{},
		&Class{},
		&Ability{},
		&Item{},
		&Equipment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func OpenDB(dbPath string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful: true,
			LogLevel: logger.Info,
		},
	)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
