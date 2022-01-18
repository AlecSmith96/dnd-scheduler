package entities

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type Player struct {
	ID       uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Sessions []Session `gorm:"many2many:players_sessions" json:"sessions"`
	Groups   []Group   `gorm:"many2many:players_groups" json:"groups"`
	Username string    `gorm:"type:varchar;NOT NULL" json:"username"`
	Cookie   string    `json:"cookie"`
}

func (p *Player) Bind(r *http.Request) error {
	// Runs after unmarshalling is complete, do postprocessing
	if p.Username == "" {
		return errors.New("missing requried Username field")
	}
	p.ID = uuid.New()
	return nil
}

func (p *Player) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before response is marshalled
	return nil
}

// swagger:parameters createNewPlayer
type PlayerParamsWrapper struct {
	// in:body
	Body Player
}