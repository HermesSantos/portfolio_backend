package routes

import (
	"fmt"
	"net/http"
)

func Router () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Your Go app is running.")
	})
}
