package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/pokemon-search-v2/src/presentation/controller"
)

func main() {
	e := echo.New()

	// TODO: routingは別の場所で定義したい
	e.GET("/health", func(context echo.Context) error {
		return context.String(http.StatusOK, "Health is OK.")
	})
	e.POST("/pokemon", controller.Search())

	if err := e.Start(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
