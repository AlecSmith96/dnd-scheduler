package adapters

import (
	"github.com/AlecSmith96/dnd-scheduler/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.StripSlashes)

	groupHandler := usecases.NewGroupHandler(db)

	r.Route("/api", func(r chi.Router) {
		r.Get("/players", usecases.GetAllPlayersHandler)
		r.Post("/players", usecases.CreatePlayerHandler)
		r.Get("/players/{playerId}", usecases.GetPlayerHandler)
		r.Put("/players/{playerId}", usecases.UpdatePlayerHandler)
		r.Delete("/players/{playerId}", usecases.DeletePlayerHandler)

		r.Get("/group", groupHandler.GetAllGroups)
		r.Post("/group", groupHandler.CreateGroup)
		r.Get("/group/{groupId}", groupHandler.GetGroup)
		r.Put("/group/{groupId}", groupHandler.UpdateGroup)
		r.Delete("/group/{groupId}", groupHandler.DeleteGroup)

		r.Get("/session", usecases.GetAllSessionsHandler)
		r.Post("/session", usecases.CreateSessionHandler)
		r.Get("/session/{sessionId}", usecases.GetSessionHandler)
		r.Put("/session/{sessionId}", usecases.UpdateSessionHandler)
		r.Delete("/session/{sessionId}", usecases.DeleteSessionHandler)
	})

	return r
}
