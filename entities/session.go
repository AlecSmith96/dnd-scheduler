package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID      uuid.UUID
	Name    string
	GroupID uuid.UUID
	From    time.Time
	To      time.Time
}
