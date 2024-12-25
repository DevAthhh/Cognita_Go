package main

import (
	"Cognita_Go/internal/handlers"
)

func main() {
	router := handlers.NewRouter()
	router.Run(":8000")
}
