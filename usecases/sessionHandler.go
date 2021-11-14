package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
)

func GetAllSessionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "All sessions here",
	}
	json.NewEncoder(w).Encode(message)
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Created session!",
	}
	json.NewEncoder(w).Encode(message)
}

func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "id")
	message := entities.Message{
		Message: "Got session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "sessionId")
	message := entities.Message{
		Message: "Updated session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "sessionId")
	message := entities.Message{
		Message: "Deleted session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}
