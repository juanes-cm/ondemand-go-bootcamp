package gateways

import (
	"encoding/json"
	"net/http"

	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/dto"
	"api-bootcamp/gateways/translators"

	log "github.com/sirupsen/logrus"
)

type PokemonGateway interface {
	GetPokemonByURl() (dto.PokemonListDTO, error)
}

type pokemonGateway struct {
	log     log.Logger
	pokeUrl string
}

func NewPokemonGateway(log log.Logger, pokeURl string) PokemonGateway {
	return &pokemonGateway{
		log:     log,
		pokeUrl: pokeURl,
	}
}

func (pg *pokemonGateway) GetPokemonByURl() (dto.PokemonListDTO, error) {
	response, err := http.Get(pg.pokeUrl)
	if err != nil {
		pg.log.WithField("api: ", pg.pokeUrl).WithError(err).Error("error getting response from api")
		return dto.PokemonListDTO{}, err
	}
	pokeView := viewmodels.PokemonView{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&pokeView)
	if err != nil {
		pg.log.WithField("api: ", pg.pokeUrl).WithError(err).Error("error while attempting to parse (unmarshal) the request")
		return dto.PokemonListDTO{}, err
	}

	return translators.ToPokemonDTO(pokeView), nil
}
