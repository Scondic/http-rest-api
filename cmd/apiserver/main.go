package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Scondic/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.CreateConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}
	server := apiserver.CreateApiServer(config)

	if err := server.StartApiServer(); err != nil {
		log.Fatalln(err)
	}
}
