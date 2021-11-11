package usecases

import (
  "encoding/json"
  "net/http"

	"github.com/go-chi/chi/v5"
  "github.com/AlecSmith96/dnd-scheduler/entities"
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
  message := entities.Message{
    Message: "Created session!",
  }
  json.NewEncoder(w).Encode(message)
}

func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
  sessionParam := chi.URLParam(r, "id")
  message := entities.Message{
    Message: "Got session, "+sessionParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}

func UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
  sessionParam := chi.URLParam(r, "sessionId")
  message := entities.Message{
    Message: "Updated session, "+sessionParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}

func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
  sessionParam := chi.URLParam(r, "sessionId")
  message := entities.Message{
    Message: "Deleted session, "+sessionParam+"!",
  }
  json.NewEncoder(w).Encode(message)
}
