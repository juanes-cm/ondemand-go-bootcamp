package viewmodels

type ResponsePokemonView struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonResponseView []ResponsePokemonView
