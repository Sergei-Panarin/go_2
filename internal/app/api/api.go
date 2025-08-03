package api

// Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start http server/configure loggers, router,database connection and etc.
func (api *API) Start() error {
	return nil
}
