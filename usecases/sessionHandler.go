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

func (handler *SessionHandler) GetAllSessions(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /sessions Session listSessions
	//
	// List all sessions
	//
	// responses:
	//	200: SessionList
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

func (handler *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /session Session createNewSession
	//
	// Create a new session
	//
	// responses:
	//	200: Session
	var session entities.Session
	if err := render.Bind(r, &session); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if result := handler.DB.Create(&session); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}

	// Try to return created player
	if err := render.Render(w, r, &session); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "id")
	message := entities.Message{
		Message: "Got session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *SessionHandler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	sessionParam := chi.URLParam(r, "sessionId")
	message := entities.Message{
		Message: "Updated session, " + sessionParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
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
