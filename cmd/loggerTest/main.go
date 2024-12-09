package main

import (
	"NexaForm/config"
	"NexaForm/service"
	"context"
	"log"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.ReadStandard("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	// Initialize the application container
	appContainer, err := service.NewAppContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize AppContainer: %v", err)
	}

	// Create a basic context
	ctx := context.Background()

	// Services to generate logs for
	services := []string{
		service.ServiceAuth,
		service.ServiceDatabase,
		service.ServiceAPI,
		service.ServiceLogger,
		service.ServiceUser,
	}

	// Log levels
	// Log levels
	logLevels := []string{"info", "error", "debug", "warn"}

	// Simulate log generation
	for i := 0; i < 10000; i++ {
		// Randomly pick a service and log level
		serviceName := services[rand.Intn(len(services))]
		logLevel := logLevels[rand.Intn(len(logLevels))]

		// Generate random fields for structured logs
		fields := []zap.Field{
			zap.String("request_id", randomString(10)),
			zap.String("user_id", randomString(8)),
			zap.String("action", randomAction()),
			zap.String("status", randomStatus()),
			zap.String("endpoint", randomEndpoint()),
		}

		// Log based on the level
		switch logLevel {
		case "info":
			appContainer.LoggerService().LogInfo(ctx, serviceName, "Generated info log", fields...)
		case "error":
			appContainer.LoggerService().LogError(ctx, serviceName, "Generated error log", fields...)
		case "debug":
			appContainer.LoggerService().LogDebug(ctx, serviceName, "Generated debug log", fields...)
		case "warn":
			appContainer.LoggerService().LogWarn(ctx, serviceName, "Generated warning log", fields...)
		case "fatal":
			// Skip fatal logs in bulk generation
			continue
		}

		// Sleep for a short time to simulate log intervals (optional)
		time.Sleep(1 * time.Millisecond)
	}

	// Flush logs
	appContainer.LoggerService().Sync()

	log.Println("Structured logs have been generated. Check the log files and Grafana for detailed insights.")
}

// Helper functions to generate random data

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomAction() string {
	actions := []string{"login", "logout", "update", "delete", "create"}
	return actions[rand.Intn(len(actions))]
}

func randomStatus() string {
	statuses := []string{"200 OK", "201 Created", "400 Bad Request", "401 Unauthorized", "500 Internal Server Error"}
	return statuses[rand.Intn(len(statuses))]
}

func randomEndpoint() string {
	endpoints := []string{"/api/v1/resource", "/api/v1/auth", "/api/v1/user", "/api/v1/admin"}
	return endpoints[rand.Intn(len(endpoints))]
}
