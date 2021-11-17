package usecases

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type GroupHandler struct {
	db *gorm.DB
}

// Return all groups with their sessions information
func (handler *GroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups := make([]entities.Group, 0)
	err := handler.db.Preload("Sessions").Find(&groups).Error

	if len(groups) == 0 || err != nil {
		message := entities.Message{
			Message: "Unable to return groups from db.",
		}
		log.Printf("Unable to return groups from db: %v", err)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(groups)
}

func (handler *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group entities.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		message := entities.Message{
			Message: "Error: Bad request body.",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	dbErr := handler.db.Create(&group).Error
	if dbErr != nil {
		log.Printf("Unable to create new record: %v", dbErr)
		message := entities.Message{
			Message: "Error: Unable to create record in db.",
		}
		json.NewEncoder(w).Encode(message)
		return
	}
	json.NewEncoder(w).Encode(group)
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
