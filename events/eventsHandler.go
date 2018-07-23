package events

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mayadata-io/chat-go/response"
)

type Events struct {
	Type      string `json:"type,omitempty"`
	Token     string `json:"token,omitempty"`
	Challenge string `json:"challenge,omitempty"`
}

// Handler function to handle all the alerts
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var events Events
	err = json.Unmarshal(body, &events)
	if err != nil {
		panic(err)
	}
	log.Printf("Events %+v", events)
	w.Header().Set(response.ContentType, response.ApplicationJSON)
	w.Write([]byte(events.Challenge))
}
