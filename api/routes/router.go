package routes

import (
	"fmt"
	"net/http"
)

func Router () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Your Go app is running.")
	})
	http.HandleFunc("/get-user-data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
		}
	})
}
