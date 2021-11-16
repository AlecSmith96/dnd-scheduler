package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
)

func GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "All groups here",
	}
	json.NewEncoder(w).Encode(message)
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Created group!",
	}
	json.NewEncoder(w).Encode(message)
}

func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	groupParam := chi.URLParam(r, "groupId")
	message := entities.Message{
		Message: "Got group, " + groupParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	groupParam := chi.URLParam(r, "groupId")
	message := entities.Message{
		Message: "Update group, " + groupParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	groupParam := chi.URLParam(r, "groupId")
	message := entities.Message{
		Message: "Delete group, " + groupParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}
