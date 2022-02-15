package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID      uuid.UUID `gorm:"PrimaryKey"`
	Name    string
	From    time.Time
	To      time.Time
}
