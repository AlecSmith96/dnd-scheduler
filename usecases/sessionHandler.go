package usecases

import (
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
	// swagger:route POST /sessions Session createNewSession
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
	// swagger:route GET /sessions/{sessionId} Session getSession
	//
	// Get an existing session
	//
	// responses:
	//	200: Session
	sessionId := chi.URLParam(r, "sessionId")
	var session entities.Session

	if result := handler.DB.First(&session, "id = ?", sessionId); result.Error != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	// Try to return found player
	if err := render.Render(w, r, &session); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *SessionHandler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	// swagger:route PATCH /sessions/{sessionId} Session updateSession
	//
	// Update an existing session
	//
	// responses:
	//	200: description: No content
	sessionId := chi.URLParam(r, "sessionId")
	var session entities.Session
	var updatedSessionData entities.Session

	if err := render.Bind(r, &updatedSessionData); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if result := handler.DB.First(&session, "id = ?", sessionId); result.Error != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	if result := handler.DB.Model(&session).Updates(&updatedSessionData); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}
}

func (handler *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /sessions/{sessionId} Session deleteSession
	//
	// Delete an existing session
	//
	// responses:
	//	200: description: No content
	sessionId := chi.URLParam(r, "sessionId")

	if result := handler.DB.Delete(&entities.Session{}, "id = ?", sessionId); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}
}

func NewSessionHandler(dbConn *gorm.DB) *SessionHandler {
	return &SessionHandler{
		DB: dbConn,
	}
}
