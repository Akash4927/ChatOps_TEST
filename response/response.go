package response

import (
	"encoding/json"
	"os"

	"github.com/mayadata-io/chat-go/environment"
)

// constants for type of response to send
const (
	ApplicationJSON = "application/json"
	ContentType     = "Content-Type"
	Ephemeral       = "ephemeral"
	InChannel       = "in_channel"
	ResponseTypeKey = "response_type"
	TextKey         = "text"
)

// ErrorMessage when chat server is not able to connect to MayaOnline
var (
	MayaErrorMessage = []byte(`{"` + ResponseTypeKey + `": "` + InChannel +
		`", "` + TextKey + `": "*Unable to reach MayaOnline. Please try again after some time.*\n*If this problem persists, feel free to contact Maya Support at mo-support@mayadata.io*"}`)

	SlashCommandErrorMessage = []byte(`{"` + ResponseTypeKey + `": "` + InChannel +
		`", "` + TextKey + `": "Sorry, I could not understand. Use ` + "`" + os.Getenv(environment.SlashCommand) + "`" + ` help to see available commands."}`)
)

var (
	jsonResponse []byte
	err          error
)

// Response is the struct which will be used in whole program to send repsonse to slack
type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

// JSONResponseOrMayaErrorMessage will convert response struct to json or
// return error related to MayaOnline
func JSONResponseOrMayaErrorMessage(response Response) []byte {
	jsonResponse, err = json.Marshal(response)
	if err != nil {
		return MayaErrorMessage
	}
	return jsonResponse
}

// JSONResponseOrSlashCommandErrorMessage will convert response struct to json or
// return error related to Slash Command
func JSONResponseOrSlashCommandErrorMessage(response Response) []byte {
	jsonResponse, err = json.Marshal(response)
	if err != nil {
		return SlashCommandErrorMessage
	}
	return jsonResponse
}
