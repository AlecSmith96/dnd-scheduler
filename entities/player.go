package entities

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	ID       uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Sessions []Session `gorm:"many2many:players_sessions" json:"sessions"`
	Groups   []Group   `gorm:"many2many:players_groups" json:"groups"`
	Username string    `gorm:"type:varchar;NOT NULL" json:"username"`
	Cookie   string    `json:"cookie"`
}

type PlayerCreate struct {
	Username string    `json:"username"`
}

func (u *Player) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return nil
}

func (p *Player) Bind(r *http.Request) error {
	// Runs after unmarshalling is complete, do postprocessing
	if p.Username == "" {
		return errors.New("missing requried Username field")
	}
	return nil
}

func (p *Player) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before response is marshalled
	return nil
}

// swagger:parameters createNewPlayer updatePlayer
type PlayerParamsWrapper struct {
	// in:body
	Body PlayerCreate
}

// swagger:parameters updatePlayer deletePlayer
type PlayerId struct {
	// in:path
	PlayerId string `json:"playerId"`
}

// swagger:response PlayerList
type PlayerList struct {
	// in:body
	Players []Player `json:"players"`
}

func (p *PlayerList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
