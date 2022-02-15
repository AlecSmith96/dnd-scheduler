package entities

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID       uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Name     string		`json:"name"`
	Sessions []Session `gorm:"foreignKey:ID" json:"sessions"`
}

type GroupCreate struct {
	Name string    `json:"username"`
}

func (g *Group) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()

	return nil
}

func (g *Group) Bind(r *http.Request) error {
	// Runs after unmarshalling is complete, do postprocessing
	if g.Name == "" {
		return errors.New("missing requried Name field")
	}
	return nil
}

func (g *Group) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before response is marshalled
	return nil
}

// swagger:parameters createNewGroup updateGroup
type GroupParamsWrapper struct {
	// in:body
	Body GroupCreate
}

// swagger:parameters getGroup updateGroup deleteGroup
type GroupId struct {
	// in:path
	GroupId string `json:"groupId"`
}

// swagger:response GroupList
type GroupList struct {
	// in:body
	Groups []Group `json:"groups"`
}

func (p *GroupList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
