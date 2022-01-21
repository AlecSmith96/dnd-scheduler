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
	if err := render.Render(w, r, &player); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
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
	// swagger:route PUT /players/{playerId} Player updatePlayer
	// Update an existing player
	// responses:
	//	200: Player
	// TODO: Implement actual function
	var player entities.Player
	var updatedPlayerData entities.Player
	playerId := chi.URLParam(r, "playerId")
	if err := render.Bind(r, &updatedPlayerData); err != nil {
		println("Binding body failed")
		render.Render(w, r, ErrRender(err))
		return
	}
	println(player.ID.String());
	if result := handler.DB.First(&player, "id = ?", playerId); result.Error != nil {
		println("Finding player", playerId ,"failed")
		render.Status(r, http.StatusNotFound)
		return
	}
	if result := handler.DB.Model(&player).Updates(&updatedPlayerData); result.Error != nil {
		println("Update of  player", playerId ,"failed")
		render.Render(w, r, ErrRender(result.Error))
		return
	}
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
