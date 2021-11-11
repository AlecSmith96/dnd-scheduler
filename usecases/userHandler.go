package usecases

import (
  "encoding/json"
  "net/http"

	"github.com/go-chi/chi/v5"
  "github.com/AlecSmith96/dnd-scheduler/entities"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
  message := entities.Message{
    Message: "Created user!",
  }
  json.NewEncoder(w).Encode(message)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
  userParam := chi.URLParam(r, "userId")
  message := entities.Message{
    Message: "Hello, "+userParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
  userParam := chi.URLParam(r, "userId")
  message := entities.Message{
    Message: "Updated user, "+userParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
  userParam := chi.URLParam(r, "userId")
  message := entities.Message{
    Message: "Deleted user, "+userParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}
