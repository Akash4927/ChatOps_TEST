package menu

import (
	"fmt"
	"net/http"
)

// Handler function is used to handle the menu requests
func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                     // parse arguments, you have to call this by yourself
	fmt.Fprintf(w, "Hello ICH Menu!") // send data to client side
}
