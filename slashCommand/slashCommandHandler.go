package slashCommand

import (
	"fmt"
	"net/http"

	"github.com/mayadata-io/chat-go/response"
)

// Handler function will handle slash command request and response
func Handler(w http.ResponseWriter, r *http.Request) {
	slashCommand, err := ParseSlashCommandRequest(r)
	if err != nil {
		w.Header().Set(response.ContentType, response.ApplicationJSON)
		w.Write(response.MayaErrorMessage)
		return
	}

	fmt.Printf("Slash command received %+v \n", slashCommand)

	w.Header().Set(response.ContentType, response.ApplicationJSON)
	w.Write(slashCommand.Response())
}
