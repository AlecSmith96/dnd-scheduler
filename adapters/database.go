package adapters

import (
	"fmt"
	"time"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConn(config *entities.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Port, config.Database.Dbname)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func TearDownDB(db *gorm.DB) {
	db.Migrator().DropTable(&entities.Session{})
}

func PopulateDB(db *gorm.DB) {
	db.Migrator().CreateTable(&entities.Session{})

	session := entities.Session{
		ID:      uuid.New(),
		GroupID: uuid.New(),
		Name:    "Session 2",
		From:    time.Now(),
		To:      time.Now().Add(72 * time.Hour),
	}
	session2 := entities.Session{
		ID:      uuid.New(),
		GroupID: uuid.New(),
		Name:    "Session 3",
		From:    time.Now(),
		To:  
		    time.Now().Add(72 * time.Hour),
	}
	db.Create(&session)
	db.Create(&session2)
}
