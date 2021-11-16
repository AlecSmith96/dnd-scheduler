package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
)

func GetAllPlayersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "All players here",
	}
	json.NewEncoder(w).Encode(message)
}

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Created player!",
	}
	json.NewEncoder(w).Encode(message)
}

func GetPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Hello, " + playerParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func UpdatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Updated player, " + playerParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Deleted player, " + playerParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}
