package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pokemon-search-v2/src/infrastructure/elasticsearch"
	"github.com/pokemon-search-v2/src/presentation/dto"
)

func SearchByPokemonName(searchWord string) string {
	repository := elasticsearch.NewPokemonRepository()
	result, err := repository.SearchByPokemonName(searchWord)
	if err != nil {
		return "検索に失敗しました。\nガラル地方のポケモンの名前か確認してください。"
	}

	log.Println(string(result))
	bytes, _ := json.Marshal(&result)

	var pokemonDto dto.PokemonDto
	if err := json.Unmarshal(bytes, &pokemonDto); err != nil {
		log.Fatal("Failed to parse json to DTO.")
	}
	return makeMessage(pokemonDto)
}

func makeMessage(dto dto.PokemonDto) string {
	// TODO: 特性の複数対応
	return fmt.Sprintf("【タイプ】\n%s\n【種族値】\nHP　　  : %s\nこうげき: %s\nぼうぎょ: %s\nとっこう: %s\nとくぼう: %s\nすばやさ: %s\n【特性】\n%s\n【夢特性】\n%s",
		dto.Types, dto.Stats.Hp, dto.Stats.Attack, dto.Stats.Defense, dto.Stats.SpAttack, dto.Stats.SpDefense, dto.Stats.Speed, dto.Abilities[0].Name, dto.HiddenAbilities[0].Name)
}
