package api

import (
	"net/http"

	"github.com/Sergei-Panarin/go_2/StandardWebServer/storage"
	"github.com/sirupsen/logrus"
)

// Configure the field Logger in the API
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// COnfigure the router in the API
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})
}

// Configure the storage in the API
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//Trying to establish connection with the storage
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
