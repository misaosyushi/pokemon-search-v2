package elasticsearch

import (
	"encoding/json"
)

type IPokemonRepository interface {
	SearchByPokemonName(searchWord string) (json.RawMessage, error)
}
