package entities

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID      uuid.UUID `gorm:"PrimaryKey"`
	GroupID uuid.UUID `gorm:"foreignKey:ID"`
	Name    string
	From    time.Time
	To      time.Time
}

type SessionCreate struct {
	Name string `json:"name"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return nil
}

func (s *Session) Bind(r *http.Request) error  {
	// Runs after unmarshalling is complete, do postprocessing
	if s.Name == "" {
		return errors.New("missing requried Name field")
	}
	return nil
}

func (s *Session) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before response is marshalled
	return nil
}

// swagger:parameters createNewSession updateSession
type SessionParamsWrapper struct {
	// in:body
	Body SessionCreate
}

// swagger:parameters getSession updateSession deleteSession
type SessionId struct {
	// in:path
	SessionId string `json:"sessionId"`
}

// swagger:response SessionList
type SessionList struct {
	// in:body
	Sessions []Session `json:"sessions"`
}

func (s *SessionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
