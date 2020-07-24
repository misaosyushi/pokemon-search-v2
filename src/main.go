package main

import (
	"encoding/json"
	"fmt"
	"github.com/pokemon-search-v2/src/infrastructure/elasticsearch"
	"github.com/pokemon-search-v2/src/presentation/dto"
)

func main() {
	pokemonInfo := elasticsearch.SearchByPokemonName()
	fmt.Println(string(pokemonInfo))
	bytes, _ := json.Marshal(&pokemonInfo)

	var pokemonDto dto.PokemonDto
	json.Unmarshal(bytes, &pokemonDto)
	fmt.Println(pokemonDto.Name)
	fmt.Println(pokemonDto.Type)
	fmt.Println(pokemonDto.BaseStats.Speed)
}
