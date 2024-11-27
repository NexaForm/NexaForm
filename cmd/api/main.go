package main

import (
	"NexaForm/config"
	"flag"
	"fmt"
	"log"
	"os"
)

var configPath = flag.String("c", "", "Pass config file")

func main() {
	config := readConfig()
	fmt.Print(config.Server.Host)
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
