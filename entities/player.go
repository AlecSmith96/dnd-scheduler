package entities

import (
	"github.com/google/uuid"
)

type Player struct {
	ID       uuid.UUID `gorm:"PrimaryKey"`
	Sessions []Session `gorm:"many2many:players_sessions"`
	Groups   []Group   `gorm:"many2many:players_groups"`
	Username string
	Cookie   string
}