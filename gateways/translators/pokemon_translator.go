package translators

import (
	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/dto"
)

func ToPokemonDTO(pk viewmodels.PokemonView) dto.PokemonListDTO {
	pokemon := make(dto.PokemonListDTO, len(pk.Pokemon))
	for value := range pokemon {
		pokemon[value] = dto.PokemonDTO{
			ID:   pk.Pokemon[value].EntryNo,
			Name: pk.Pokemon[value].Species.Name,
		}
	}
	return pokemon
}
