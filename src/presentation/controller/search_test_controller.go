package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pokemon-search-v2/src/usecase"
	"net/http"
)

func Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, usecase.SearchByPokemonName("メッソン"))
	}
}
