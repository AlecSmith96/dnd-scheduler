package entities

import (
	"github.com/google/uuid"
)

type Player struct {
	ID       uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Sessions []Session `gorm:"many2many:players_sessions" json:"sessions"`
	Groups   []Group   `gorm:"many2many:players_groups" json:"groups"`
	Username string    `gorm:"type:varchar;NOT NULL" json:"name"`
	Cookie   string    `json:"cookie"`
}
