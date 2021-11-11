package usecases

import (
  "encoding/json"
  "net/http"

	"github.com/go-chi/chi/v5"
  "github.com/AlecSmith96/dnd-scheduler/entities"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
  message := entities.Message{
    Message: "Created group!",
  }
  json.NewEncoder(w).Encode(message)
}

func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
  groupParam := chi.URLParam(r, "groupId")
  message := entities.Message{
    Message: "Got group, "+groupParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}

func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
  groupParam := chi.URLParam(r, "groupId")
  message := entities.Message{
    Message: "Update group, "+groupParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
  groupParam := chi.URLParam(r, "groupId")
  message := entities.Message{
    Message: "Delete group, "+groupParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}
