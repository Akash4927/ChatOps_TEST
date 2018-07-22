package events

import (
	"fmt"
	"net/http"
)

// Handler function to handle all the alerts
func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // parse arguments, you have to call this by yourself
	fmt.Fprintf(w, "Hello Events!") // send data to client side
}
