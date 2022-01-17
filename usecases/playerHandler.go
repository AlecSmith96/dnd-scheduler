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
	DB *gorm.DB
}

// List of players
// swagger:response PlayerList
type PlayerList struct {
	// in:body
	Players []*entities.Player `json:"players"`
}

func (p *PlayerList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (handler *PlayerHandler) PlayerListResponse() *PlayerList {
	players := &PlayerList{}
	handler.DB.Preload("Sessions").Preload("Groups").Find(&players.Players)
	return players
}

// swagger:route GET /players Player listPlayers
//
// List all players
//
// responses:
//	200: PlayerList
func (handler *PlayerHandler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	if err := render.Render(w, r, handler.PlayerListResponse()); err != nil {
		_ = render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	message := entities.Message{
		Message: "Hey",
	}
	json.NewEncoder(w).Encode(message)
}

func (handler *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual function
	playerParam := chi.URLParam(r, "playerId")
	message := entities.Message{
		Message: "Hello " + playerParam + "!",
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
		DB: dbConn,
	}
}
