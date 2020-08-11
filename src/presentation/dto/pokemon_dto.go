package dto

type PokemonDto struct {
	Name  string `json:"name"`
	Types string `json:"types"`
	Stats struct {
		Hp        string `json:"hp"`
		Attack    string `json:"attack"`
		Defense   string `json:"defense"`
		SpAttack  string `json:"spAttack"`
		SpDefense string `json:"spDefence"`
		Speed     string `json:"speed"`
	} `json:"stats"`
	Abilities       []Ability `json:"abilities"`
	HiddenAbilities []Ability `json:"hiddenAbilities"`
}

type Ability struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
