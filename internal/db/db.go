package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ev-n-er/jarvis_co_bot/internal/model"
)

type BotDB struct {
	gorm *gorm.DB
}

func Initialize() {
	db := createGorm()

	if db != nil {
		db.AutoMigrate(&model.Visit{})
	}
}

func Create() *BotDB {
	return &BotDB{
		gorm: createGorm(),
	}
}

func createGorm() *gorm.DB {

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil
	}

	return db
}

func (botDB *BotDB) CreateVisit(userId uint, date time.Time, visitType int8) bool {
	visit := model.Visit{
		UserID:    userId,
		Date:      date,
		Type:      visitType,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()}

	result := botDB.gorm.Create(&visit)

	if result.Error != nil {
		log.Fatalf("Could not connect to database: %v", result.Error)
		return false
	}

	return true
}
