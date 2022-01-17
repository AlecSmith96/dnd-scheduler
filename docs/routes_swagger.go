// Package docs D&D Scheduler API.
// REST APIs for scheduling D&D sessions
//
// Schemes: http, https
// Host: localhost:8080
// BasePath: /api
// Version: v1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package docs

import (
	"github.com/AlecSmith96/dnd-scheduler/entities"
)

// swagger:route GET / GetHomepage request_1
// Get Homepage returns JSON welcome message.
// responses:
//   200: Message

// Message field with the welcome message.
// swagger:response Message
type GetHomepageResponseWrapper struct {
	// in:body
	Body entities.Message
}

// swagger:parameters 1
type GetHomepageParamsWrapper struct {
	// No params in request body.
	//in:body
}
