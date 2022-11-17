package controllers

import (
	"net/http"

	"api-bootcamp/controllers/translators"
	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/mediators"

	log "github.com/sirupsen/logrus"
)

type PokemonController struct {
	PokemonFactory func() mediators.PokemonMediator
	Log            log.Logger
}

// Get just one element from the csv file fitler by the ID param
func (pc *PokemonController) StoreInCSV(w http.ResponseWriter, r *http.Request) {
	pc.Log.Info("calling StoreInCSV controller")
	response := viewmodels.BaseResponse{}

	mediator := pc.PokemonFactory()
	clientDto, err := mediator.StoreInCSV()
	if err != nil {
		pc.Log.WithError(err).Error("error calling mediator StoreInCSV function")
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	response.Data = translators.ToPokemonResponseView(clientDto)
	pc.Log.Info("store in csv succeded")
	SendResponse(w, http.StatusOK, response)
}
