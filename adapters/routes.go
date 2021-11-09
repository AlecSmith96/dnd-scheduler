package adapters

import (
	"log"
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/usecases"
)

// HandleRequests serves HTTP server and defines endpoints
func HandleRequests() {
	log.Println("Server listening on port 8080...")
	http.HandleFunc("/", usecases.GetHomepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
