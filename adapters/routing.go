package adapters

import (
	"log"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

  r.Route("/api", func(r chi.Router) {
    r.Post("/users/", usecases.CreateUserHandler)
    r.Get("/users/{userId}", usecases.GetUserHandler)
    r.Put("/users/{userId}", usecases.UpdateUserHandler)
    r.Delete("/users/{userId}", usecases.DeleteUserHandler)

    r.Post("/group", usecases.CreateGroupHandler)
    r.Get("/group/{groupId}", usecases.GetGroupHandler)
    r.Put("/group/{groupId}", usecases.UpdateGroupHandler)
    r.Delete("/group/{groupId}", usecases.DeleteGroupHandler)

    r.Post("/session", usecases.CreateSessionHandler)
    r.Get("/session/{sessionId}", usecases.GetSessionHandler)
    r.Put("/session/{sessionId}", usecases.UpdateSessionHandler)
    r.Delete("/session/{sessionId}", usecases.DeleteSessionHandler)
  })

	port := "3000"
	log.Println("Listening on port", port)
	http.ListenAndServe(":"+port, r)
}
