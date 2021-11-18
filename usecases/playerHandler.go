package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type PlayerHandler struct {
	db *gorm.DB
}

type PlayerList struct {
	Items []*entities.Player `json:"items"`
}

func (p *PlayerList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status" example:"Resource not found."`                                         // user-level status message
	AppCode    int64  `json:"code,omitempty" example:"404"`                                                 // application-specific error code
	ErrorText  string `json:"error,omitempty" example:"The requested resource was not found on the server"` // application-level error message, for debugging
}

func (p *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (handler *PlayerHandler) PlayerListResponse() *PlayerList {
	players := &PlayerList{}
	handler.db.Preload("Sessions").Preload("Groups").Find(&players.Items)
	return players
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func (handler *PlayerHandler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	if err := render.Render(w, r, handler.PlayerListResponse()); err != nil {
		_ = render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Created player!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Hello, " + playerParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Updated player, " + playerParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Deleted player, " + playerParam + "!",
	}
	json.NewEncoder(w).Encode(message)
}

func NewPlayerHandler(dbConn *gorm.DB) *PlayerHandler {
	return &PlayerHandler{
		db: dbConn,
	}
}
