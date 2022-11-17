package api

import (
	"net/http"

	"api-bootcamp/configs"
	"api-bootcamp/controllers"
	"api-bootcamp/gateways"
	"api-bootcamp/mediators"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Routes(log log.Logger, config configs.AppConfig) http.Handler {

	apiBootcampController, pokemonController := generateController(log, config)

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/{id:[0-9]+}", apiBootcampController.GetCSVElementByID).Methods(http.MethodGet)
	router.HandleFunc("/", pokemonController.StoreInCSV).Methods(http.MethodPost)

	return router
}

func generateController(log log.Logger, config configs.AppConfig) (controllers.ApiBootcampController, controllers.PokemonController) {

	apiBootcampFactory := func() mediators.ApiMediator {
		return mediators.NewApiMediator(config.CsvFile)
	}
	apiController := controllers.ApiBootcampController{
		ApiBootcampFactory: apiBootcampFactory,
	}

	pokemonFactory := func() mediators.PokemonMediator {
		pokeGateway := gateways.NewPokemonGateway(log, config.PokeUrl)
		return mediators.NewPokemonMediator(log, pokeGateway, config.CsvFile)
	}
	pokeController := controllers.PokemonController{
		PokemonFactory: pokemonFactory,
		Log:            log,
	}

	return apiController, pokeController
}
