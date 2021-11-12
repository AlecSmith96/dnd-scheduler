package entities

import (
	"github.com/google/uuid"
)

type Group struct {
	ID       uuid.UUID `gorm:"PrimaryKey"`
	Name     string
	Sessions []Session
}
