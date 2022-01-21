package adapters

import (
	"github.com/AlecSmith96/dnd-scheduler/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.StripSlashes, cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	groupHandler := usecases.NewGroupHandler(db)
	playerHandler := usecases.NewPlayerHandler(db)

	r.Route("/api", func(r chi.Router) {
		r.Get("/players", playerHandler.GetAllPlayers)
		r.Post("/players", playerHandler.CreatePlayer)
		r.Get("/players/{playerId}", playerHandler.GetPlayer)
		r.Put("/players/{playerId}", playerHandler.UpdatePlayer)
		r.Delete("/players/{playerId}", playerHandler.DeletePlayer)

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
