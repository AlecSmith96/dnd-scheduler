package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID   uuid.UUID `gorm:"column:session_id;primary_key"`
	Name string    `gorm:"column:name"`
	From time.Time `gorm:"column:from"`
	To   time.Time `gorm:"column:to"`
}
