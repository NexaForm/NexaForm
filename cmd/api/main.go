package main

import (
	http_server "NexaForm/api/http"
	"NexaForm/config"
	"NexaForm/service"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

var configPath = flag.String("c", "", "Pass config file")

func main() {
	config := readConfig()
	fmt.Print(config.Server.Host)
	app, err := service.NewAppContainer(config)
	if err != nil {
		log.Fatal(err)
	}
	go func() { app.FileService().ListenForEvents(context.Background()) }()
	http_server.Run(config, app)
}

func readConfig() config.Config {
	flag.Parse()

	if envConfig := os.Getenv("NEXAFORM_CONFIG_PATH"); len(envConfig) > 0 {
		*configPath = envConfig
		log.Printf("Using config path from environment variable: %s", envConfig)
	}
	if len(*configPath) < 0 {
		log.Fatal("Config path is empty")
	}

	cfg, err := config.ReadStandard(*configPath)

	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}

	return cfg
}
