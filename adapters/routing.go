package adapters

import (
	"github.com/AlecSmith96/dnd-scheduler/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.StripSlashes)

	r.Route("/api", func(r chi.Router) {
		r.Get("/players", usecases.GetAllPlayersHandler)
		r.Post("/players", usecases.CreatePlayerHandler)
		r.Get("/players/{playerId}", usecases.GetPlayerHandler)
		r.Put("/players/{playerId}", usecases.UpdatePlayerHandler)
		r.Delete("/players/{playerId}", usecases.DeletePlayerHandler)

		r.Get("/group", usecases.GetAllGroupsHandler)
		r.Post("/group", usecases.CreateGroupHandler)
		r.Get("/group/{groupId}", usecases.GetGroupHandler)
		r.Put("/group/{groupId}", usecases.UpdateGroupHandler)
		r.Delete("/group/{groupId}", usecases.DeleteGroupHandler)

		r.Get("/session", usecases.GetAllSessionsHandler)
		r.Post("/session", usecases.CreateSessionHandler)
		r.Get("/session/{sessionId}", usecases.GetSessionHandler)
		r.Put("/session/{sessionId}", usecases.UpdateSessionHandler)
		r.Delete("/session/{sessionId}", usecases.DeleteSessionHandler)
	})

	return r
}
