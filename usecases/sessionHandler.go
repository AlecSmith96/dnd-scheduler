package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type SessionHandler struct {
	DB *gorm.DB
}

func (handler *SessionHandler) GetAllSessionsHandler(w http.ResponseWriter, r *http.Request) {
	var sessions entities.SessionList
	if result := handler.DB.Find(&sessions.Sessions); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}

	// Try to return players
	if err := render.Render(w, r, &sessions); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *SessionHandler) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Created session!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *SessionHandler) GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "id")
	message := entities.Message{
		Message: "Got session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *SessionHandler) UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "sessionId")
	message := entities.Message{
		Message: "Updated session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *SessionHandler) DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "sessionId")
	message := entities.Message{
		Message: "Deleted session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func NewSessionHandler(dbConn *gorm.DB) *SessionHandler {
	return &SessionHandler{
		DB: dbConn,
	}
}
