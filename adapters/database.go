package adapters

import (
	"fmt"
	// "time"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	// "github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConn(config *entities.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Port, config.Database.Dbname)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func TearDownDB(db *gorm.DB) {
	// db.Exec("DELETE FROM players")
	// db.Exec("DELETE FROM groups")
	// db.Exec("DELETE FROM sessions")

	db.Migrator().DropTable(&entities.Player{})
	db.Migrator().DropTable(&entities.Group{})
	db.Migrator().DropTable(&entities.Session{})
}

// Creates db tables and populates with test data
func PopulateDB(db *gorm.DB) {
	db.Migrator().CreateTable(&entities.Session{})
	db.Migrator().CreateTable(&entities.Player{})
	db.Migrator().CreateTable(&entities.Group{})

	// db.AutoMigrate(&entities.Player{},&entities.Session{}, &entities.Group{})

	// player1 := entities.Player{
	// 	ID: uuid.New(),
	// 	Username: "Player 1",
	// 	Cookie: "",
	// }
	// player2 := entities.Player{
	// 	ID: uuid.New(),
	// 	Username: "Player 1",
	// 	Cookie: "",
	// }
	// player3 := entities.Player{
	// 	ID: uuid.New(),
	// 	Username: "Player 1",
	// 	Cookie: "",
	// }

	// session1 := entities.Session{
	// 	ID: uuid.New(),
	// 	Name: "Session 1",
	// 	From: time.Now(),
	// 	To: time.Now().Add(12 * time.Hour),
	// }

	// session2 := entities.Session{
	// 	ID: uuid.New(),
	// 	Name: "Session 2",
	// 	From: time.Now(),
	// 	To: time.Now().Add(12 * time.Hour),
	// }

	// group := entities.Group{
	// 	ID: uuid.New(),
	// 	Players: []entities.Player{player1, player2, player3},
	// 	Sessions: []entities.Session{session1, session2},
	// }

	// db.Create(player1)
	// db.Create(player2)
	// db.Create(player3)
	// db.Create(group)
	// db.Create(session1)
	// db.Create(session2)
}

