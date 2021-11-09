package entities

import (
	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID `gorm:"column:player_id;primary_key"`
	Username  string    `gorm:"column:username"`
	Cookie    string    `gorm:"column:cookie"`
	SessionID uuid.UUID `gorm:"foreign_key:session_id`
	GroupID   uuid.UUID `gorm:foreign_key:group_id`
}
