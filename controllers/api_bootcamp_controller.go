package controllers

import (
	"net/http"
	"strconv"

	"api-bootcamp/controllers/translators"
	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/mediators"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ApiBootcampController struct {
	ApiBootcampFactory func() mediators.ApiMediator
	Log                log.Logger
}

// Get just one element from the csv file fitler by the ID param
func (ac *ApiBootcampController) GetCSVElementByID(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info("calling GetCSVElementByID controller")
	response := viewmodels.BaseResponse{}
	vars := mux.Vars(r)
	csvID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ac.Log.WithError(err).Error("error converting elementID to int")
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	mediator := ac.ApiBootcampFactory()
	clientDto, err := mediator.GetCSVElementByID(csvID)
	if err != nil {
		ac.Log.WithError(err).Error("error calling mediator GetCSVElementByID function")
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	response.Data = translators.ToResponseView(clientDto)
	ac.Log.WithField("id_element: ", csvID).Info("getting element succeded")
	SendResponse(w, http.StatusOK, response)
}
