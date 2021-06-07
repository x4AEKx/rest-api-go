package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/x4AEKx/rest-api-go/internal/app/apiserver"

	_ "github.com/x4AEKx/rest-api-go/migrations"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
