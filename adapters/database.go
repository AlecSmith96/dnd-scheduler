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
	db.Migrator().DropTable(&entities.Group{})
	db.Migrator().DropTable(&entities.Session{})
	db.Migrator().DropTable(&entities.Player{})
	db.Migrator().DropTable("players_sessions")
	db.Migrator().DropTable("players_groups")
}

// Creates db tables and populates with test data
func PopulateDB(db *gorm.DB) {
	db.AutoMigrate(&entities.Player{}, &entities.Session{}, &entities.Group{})

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

	session1 := entities.Session{
		ID:      uuid.New(),
		GroupID: group.ID,
		Name:    "Session 1",
		From:    time.Now(),
		To:      time.Now().Add(12 * time.Hour),
	}
	session2 := entities.Session{
		ID:      uuid.New(),
		GroupID: group.ID,
		Name:    "Session 2",
		From:    time.Now().Add(24 * time.Hour),
		To:      time.Now().Add(36 * time.Hour),
	}

	db.Create(&session1)
	db.Create(&session2)

	// Make player >-< groups connections
	for index := range players {
		db.Model(&players[index]).Association("Groups").Append(&group)
	}

	// Make player >-< session connections
	for index := range players {
		db.Model(&players[index]).Association("Sessions").Append(&session1)
		if index%2 == 0 {
			db.Model(&players[index]).Association("Sessions").Append(&session2)
		}
	}
}
