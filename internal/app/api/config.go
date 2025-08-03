package api

// General instance for API server of the RESTful application
type Config struct {
	//Port on which the API server will listen
	BindAddr string `toml: "bind_addr"` // Bind address for the API server
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080", // Default port
	}
}
