package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/pokemon-search-v2/src/infrastructure/elasticsearch"
	"github.com/pokemon-search-v2/src/infrastructure/line"
	"github.com/pokemon-search-v2/src/presentation/dto"
)

func Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		pokemonInfo := elasticsearch.SearchByPokemonName()
		fmt.Println(string(pokemonInfo))
		bytes, _ := json.Marshal(&pokemonInfo)

		var pokemonDto dto.PokemonDto
		if err := json.Unmarshal(bytes, &pokemonDto); err != nil {
			log.Fatal("Failed to parse json to DTO.")
		}
		return c.String(http.StatusOK, line.MakeMessage(pokemonDto))
	}
}
