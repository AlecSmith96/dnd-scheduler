package entities

import (
	"github.com/google/uuid"
)

type Group struct {
	ID       uuid.UUID `gorm:"column:group_id;primary_key"`
	Sessions []Session `gorm:foreignKey:session_id;"`
	Players  []Player  `gorm:foreignKey:player_id;"`
}
