package adapters

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConn(config *entities.Config) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), 	// io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Info, 	// Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  true,          // Disable color
		},
	  )

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Port, config.Database.Dbname)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}

func TearDownDB(db *gorm.DB) {
	db.Migrator().DropTable(&entities.Group{})
	db.Migrator().DropTable(&entities.Session{})
	db.Migrator().DropTable(&entities.Player{})
	db.Migrator().DropTable("players_sessions")
	db.Migrator().DropTable("players_groups")
}

// Creates db tables and populates with test data
func PopulateDB(db *gorm.DB) {
	db.AutoMigrate(&entities.Group{}, &entities.Player{}, &entities.Session{})

	players := []entities.Player{
		{ID: uuid.New(), Username: "Player 1", Cookie: ""},
		{ID: uuid.New(), Username: "Player 2", Cookie: ""},
		{ID: uuid.New(), Username: "Player 3", Cookie: ""},
	}
	for index := range players {
		db.Create(&players[index])
	}

	group := entities.Group{ID: uuid.New(), Name: "My Group"}
	db.Create(&group)

	sessions := []entities.Session{
		{
			ID: uuid.New(), 
			GroupID: group.ID,
			Name: "Session 1", 
			From: time.Now(), 
			To: time.Now().Add(12 * time.Hour),
		},
		{
			ID:      uuid.New(),
			GroupID: group.ID,
			Name:    "Session 2",
			From:    time.Now().Add(24 * time.Hour),
			To:      time.Now().Add(36 * time.Hour),
		},
		
	}

	for index := range sessions {
		db.Create(&sessions[index])
	}

	// Make player >-< groups connections
	for index := range players {
		db.Model(&players[index]).Association("Groups").Append(&group)
	}
}
