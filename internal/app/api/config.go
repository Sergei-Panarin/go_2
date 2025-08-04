package api

// General instance for API server of the RESTful application
type Config struct {
	//Port on which the API server will listen
	BindAddr string `toml: "bind_addr"` // Bind address for the API server
	//Logger Level
	LoggerLevel string `toml: "logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080", // Default port
		LoggerLevel: "debug", //Default Logger level
	}
}
