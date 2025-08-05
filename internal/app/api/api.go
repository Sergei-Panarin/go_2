package api

import (
	"net/http"

	"github.com/Sergei-Panarin/go_2/StandardWebServer/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Adding field for work with storage
	storage *storage.Storage
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure loggers, router,database connection and etc.
func (api *API) Start() error {
	//Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//Configure Logger
	api.logger.Info("Starting api server at port ", api.config.BindAddr)
	//Configure router
	api.configureRouterField()
	//COnfigure Storage connection
	if err := api.configureStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
