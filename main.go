package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mayadata-io/maya-chatops/chat-go/alerts"
	"github.com/mayadata-io/maya-chatops/chat-go/interactiveComponent/button"
	"github.com/mayadata-io/maya-chatops/chat-go/interactiveComponent/menu"
	"github.com/mayadata-io/maya-chatops/chat-go/slashCommand"
)

func main() {
	port := "9090"

	// Route for handling web browser request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is Chatbot")
	})

	// Route for handling Slack requests
	http.HandleFunc("/slack", slashCommand.Handler)
	http.HandleFunc("/intButton", button.Handler)
	http.HandleFunc("/intMenu", menu.Handler)

	// Route for handling Alerts
	http.HandleFunc("/maya-events", alerts.Handler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
