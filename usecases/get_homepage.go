package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
)

// GetHomepage returns welcome message
func GetHomepage(w http.ResponseWriter, r *http.Request) {
	message := entities.Message{
		Message: "Hello, Seb!",
	}
	json.NewEncoder(w).Encode(message)
}
