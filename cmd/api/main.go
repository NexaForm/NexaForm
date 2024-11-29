package main

import (
	"NexaForm/api/http/middlerwares/logger"
	"NexaForm/config"
	"flag"
	"log"
	"os"
)

var configPath = flag.String("c", "", "Pass config file")

func main() {
	// TODO : remove this part
	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	config := readConfig()
	for range 3 {
		logger.New("Test Log").SetUser("Shahryar").Debug("foo/p1", lorem, config)
		logger.New("Test Log").SetUser("Shahryar").Info("foo/p1", lorem, config)
		logger.New("Test Log").SetUser("Shahryar").Warning("foo/p1", lorem, config)
		logger.New("Test Log").SetUser("Shahryar").Error("foo/p1", lorem, config)
		logger.New("Test Log").SetUser("Shahryar").Fatal("foo/p1", lorem, config)
		logger.New("Test Log").Debug("foo/p1", lorem, config)
		logger.New("Test Log").Info("foo/p1", lorem, config)
		logger.New("Test Log").Warning("foo/p1", lorem, config)
		logger.New("Test Log").Error("foo/p1", lorem, config)
		logger.New("Test Log").Fatal("foo/p1", lorem, config)
	}
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
