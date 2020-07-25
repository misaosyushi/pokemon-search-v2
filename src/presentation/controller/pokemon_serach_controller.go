package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/pokemon-search-v2/src/infrastructure/line"
	"github.com/pokemon-search-v2/src/presentation/dto"
	"github.com/pokemon-search-v2/src/usecase"
)

func Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		pokemonInfo := usecase.SearchByPokemonName()
		fmt.Println(string(pokemonInfo))
		bytes, _ := json.Marshal(&pokemonInfo)

		// TODO: 検索結果がなかったとき
		var pokemonDto dto.PokemonDto
		if err := json.Unmarshal(bytes, &pokemonDto); err != nil {
			log.Fatal("Failed to parse json to DTO.")
		}
		return c.String(http.StatusOK, line.MakeMessage(pokemonDto))
	}
}
