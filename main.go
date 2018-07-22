package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mayadata-io/chat-go/alerts"
	"github.com/mayadata-io/chat-go/events"
	"github.com/mayadata-io/chat-go/interactiveComponent/button"
	"github.com/mayadata-io/chat-go/interactiveComponent/menu"
	"github.com/mayadata-io/chat-go/slashCommand"
)

func main() {
	port := "9090"

	// Route for handling web browser request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is Chatbot")
	})

	// Route for handling Slack requests
	http.HandleFunc("/slash", slashCommand.Handler)
	http.HandleFunc("/interactiveButton", button.Handler)
	http.HandleFunc("/interactiveMenu", menu.Handler)
	http.HandleFunc("/events", events.Handler)

	// Route for handling Alerts
	http.HandleFunc("/alerts", alerts.Handler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
