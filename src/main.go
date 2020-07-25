package main

import (
	"log"

	"github.com/pokemon-search-v2/src/infrastructure/echo"
)

func main() {
	if err := echo.StartServer(); err != nil {
		log.Fatalf("StartServer() is error: %v", err)
	}
}
