package main

import (
	"fmt"
	"net/http"
	"os"

	"api-bootcamp/api"
	"api-bootcamp/configs"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger.
var log = logrus.New()

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("error : ", err)
		os.Exit(1)
	}
	router := api.Routes(*log, config)

	srv := &http.Server{
		Addr:    config.Port,
		Handler: router,
	}
	srv.ListenAndServe()
}
