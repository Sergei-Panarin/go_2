package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
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

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
