package adapters

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
)

func Serve(router chi.Router, config *entities.Config) {
	port := strconv.Itoa(config.Server.Port)
	host := config.Server.Host
	log.Println("Listening on port", port)
	http.ListenAndServe(host+":"+port, router)
}
