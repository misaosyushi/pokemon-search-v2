package elasticsearch

import (
	"encoding/json"
)

type IPokemonRepository interface {
	SearchByPokemonName() json.RawMessage
}
