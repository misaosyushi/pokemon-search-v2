package usecase

import (
	"encoding/json"

	"github.com/pokemon-search-v2/src/infrastructure/elasticsearch"
)

func SearchByPokemonName() json.RawMessage {
	repository := elasticsearch.NewPokemonRepository()
	return repository.SearchByPokemonName()
}
