package main

import (
	"portfolio_backend/api/routes"
	"portfolio_backend/api/server"
)

func main() {
	routes.Router()
	err := server.Server()
	if err != nil {
    panic(err)
  }
}
