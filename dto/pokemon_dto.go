package dto

type PokemonDTO struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
}

type PokemonListDTO []PokemonDTO
