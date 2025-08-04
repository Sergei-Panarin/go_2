package api

import "github.com/sirupsen/logrus"

// Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

// Start http server/configure loggers, router,database connection and etc.
func (api *API) Start() error {
	//Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting api server at port ", api.config.BindAddr)
	return nil
}
