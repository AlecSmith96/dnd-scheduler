package usecases

import (
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type PlayerHandler struct {
	DB *gorm.DB
}

func (handler *PlayerHandler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /players Player listPlayers
	//
	// List all players
	//
	// responses:
	//	200: PlayerList
	var players entities.PlayerList
	if result := handler.DB.Find(&players.Players); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}

	// Try to return players
	if err := render.Render(w, r, &players); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /players Player createNewPlayer
	//
	// Create a new player
	//
	// responses:
	//	200: Player
	var player entities.Player
	if err := render.Bind(r, &player); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if result := handler.DB.Create(&player); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}

	// Try to return created player
	if err := render.Render(w, r, &player); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /players/{playerId} Player getPlayer
	//
	// Get a specific player
	//
	// responses:
	//	200: Player
	playerId := chi.URLParam(r, "playerId")
	var player entities.Player

	if result := handler.DB.First(&player, "id = ?", playerId); result.Error != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	// Try to return found player
	if err := render.Render(w, r, &player); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	// swagger:route PATCH /players/{playerId} Player updatePlayer
	//
	// Update an existing player
	//
	// responses:
	//	200: description: No content
	playerId := chi.URLParam(r, "playerId")
	var player entities.Player
	var updatedPlayerData entities.Player

	if err := render.Bind(r, &updatedPlayerData); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if result := handler.DB.First(&player, "id = ?", playerId); result.Error != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	if result := handler.DB.Model(&player).Updates(&updatedPlayerData); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}
}

func (handler *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /players/{playerId} Player deletePlayer
	//
	// Delete an existing player
	//
	// responses:
	//	200: description: No content
	playerId := chi.URLParam(r, "playerId")

	if result := handler.DB.Delete(&entities.Player{}, "id = ?", playerId); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}
}

func NewPlayerHandler(dbConn *gorm.DB) *PlayerHandler {
	return &PlayerHandler{
		DB: dbConn,
	}
}
