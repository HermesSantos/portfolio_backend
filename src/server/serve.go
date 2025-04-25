package server

import (
	"fmt"
	"net/http"
)
func Server () error {
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("Server failed: %v\n", err)
	}
	return nil
}
