package main

import (
	"portfolio_backend/src/routes"
	"portfolio_backend/src/server"
)

func main() {
	routes.Router()
	err := server.Server()
	if err != nil {
    panic(err)
  }
}
