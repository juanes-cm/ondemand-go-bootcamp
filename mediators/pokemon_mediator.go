package mediators

import (
	"os"

	"api-bootcamp/dto"
	"api-bootcamp/gateways"

	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
)

type PokemonMediator interface {
	StoreInCSV() (dto.PokemonListDTO, error)
}

type pokemonMediator struct {
	log         log.Logger
	csvFile     string
	pokeGateway gateways.PokemonGateway
}

func NewPokemonMediator(log log.Logger, pokeGateway gateways.PokemonGateway, csvFile string) PokemonMediator {
	return &pokemonMediator{
		log:         log,
		pokeGateway: pokeGateway,
		csvFile:     csvFile,
	}
}

func (pm *pokemonMediator) StoreInCSV() (dto.PokemonListDTO, error) {

	pokeDTO, err := pm.pokeGateway.GetPokemonByURl()
	if err != nil {
		pm.log.WithField("pokemon gateway", pm.pokeGateway).WithError(err).Error("error in pokemon gateway")
		return dto.PokemonListDTO{}, err
	}

	in, err := os.OpenFile(pm.csvFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		pm.log.WithField("file: ", pm.csvFile).WithError(err).Error("error opening csv file")
		return dto.PokemonListDTO{}, err
	}
	defer in.Close()

	if err := gocsv.MarshalFile(&pokeDTO, in); err != nil {
		pm.log.WithField("file: ", pm.csvFile).WithError(err).Error("error writting info into the csv file")
		return dto.PokemonListDTO{}, err
	}

	return pokeDTO, nil
}
