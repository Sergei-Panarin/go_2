package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Sergei-Panarin/go_2/StandardWebServer/internal/app/api"
)

var (
	configPath string //= "configs/api.toml" // Path to the configuration file
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to the configuration file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("Starting API server...")
	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	log.Printf("Loaded config: %+v\n", config)
	if err != nil {
		log.Println("can not configs file. using default values for server: ", err)
	}
	server := api.New(config)
	//api server start
	log.Fatal(server.Start())
}
