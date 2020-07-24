package dto

type PokemonDto struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	BaseStats struct {
		Hp             string `json:"hp"`
		Attack         string `json:"attack"`
		Defense        string `json:"defense"`
		SpecialAttack  string `json:"specialAttack"`
		SpecialDefense string `json:"specialDefense"`
		Speed          string `json:"speed"`
	} `json:"baseStats"`
}
