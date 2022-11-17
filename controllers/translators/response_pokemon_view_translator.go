package translators

import (
	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/dto"
)

func ToPokemonResponseView(list dto.PokemonListDTO) viewmodels.PokemonResponseView {
	result := make(viewmodels.PokemonResponseView, len(list))
	for value := range result {
		result[value] = viewmodels.ResponsePokemonView{
			ID:   list[value].ID,
			Name: list[value].Name,
		}
	}
	return result
}
