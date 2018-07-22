package button

import (
	"fmt"
	"net/http"
)

// Handler function is used to handle button requests
func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                       // parse arguments, you have to call this by yourself
	fmt.Fprintf(w, "Hello ICH Button!") // send data to client side
}
