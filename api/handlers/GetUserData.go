package handlers

import (
	"fmt"
	"net/http"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }
	fmt.Println("po")
}
