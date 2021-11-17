package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type GroupHandler struct {
	db *gorm.DB
}

func (handler *GroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "All groups here",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Created group!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	groupParam := chi.URLParam(r, "groupId")
	message := entities.Message{
		Message: "Got group, " + groupParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	groupParam := chi.URLParam(r, "groupId")
	message := entities.Message{
		Message: "Update group, " + groupParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	groupParam := chi.URLParam(r, "groupId")
	message := entities.Message{
		Message: "Delete group, " + groupParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func NewGroupHandler(dbConn *gorm.DB) *GroupHandler {
	return &GroupHandler{
		db: dbConn,
	}
}
