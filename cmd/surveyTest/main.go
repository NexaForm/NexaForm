package main

import (
	"NexaForm/config"
	"NexaForm/internal/survey"
	"NexaForm/internal/user"
	"NexaForm/service"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

var configPath = flag.String("c", "", "Pass config file")

func main() {
	config := readConfig()
	fmt.Print(config.Server.Host)

	app, err := service.NewAppContainer(config)
	if err != nil {
		log.Fatal(err)
	}

	err = populateSurveys(context.Background(), app.SurveyService())
	if err != nil {
		log.Fatalf("Failed to populate surveys: %v", err)
	}

	err = fetchAndPrintSurveys(context.Background(), app.SurveyService())
	if err != nil {
		log.Fatalf("Failed to fetch surveys: %v", err)
	}

	// http_server.Run(config, app)
}

func readConfig() config.Config {
	flag.Parse()

	if envConfig := os.Getenv("NEXAFORM_CONFIG_PATH"); len(envConfig) > 0 {
		*configPath = envConfig
		log.Printf("Using config path from environment variable: %s", envConfig)
	}
	if len(*configPath) < 1 {
		log.Fatal("Config path is empty")
	}

	cfg, err := config.ReadStandard(*configPath)

	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}

	return cfg
}

func populateSurveys(ctx context.Context, surveyService *service.SurveyService) error {
	ownerID := uuid.MustParse("d73db16f-64e2-4686-833c-07e52a838587")
	istrue := true
	newSurvey := &survey.Survey{
		OwnerID:       ownerID,
		Title:         "Customer Satisfaction Survey",
		Description:   "A survey to gauge customer satisfaction levels.",
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(7 * 24 * time.Hour),
		Visibility:    survey.ALL,
		AllowedMinAge: 18,
		AllowedMaxAge: 65,
		AllowedGender: user.Female,
		MaxEditTime:   time.Now().Add(1 * time.Hour),
		IsOrdered:     true,
		IsReversable:  false,
		MaxTries:      3,
		Questions: []survey.Question{
			{
				Description:   "How would you rate our service?",
				Type:          "Poll",
				Order:         1,
				IsConditional: true,
				Options: []survey.Option{
					{Text: "Excellent"},
					{Text: "Good"},
					{Text: "Fair"},
					{Text: "Poor"},
				},
			},
			{
				Description: "Would you recommend us to others?",
				Type:        "Text",
				Order:       2,
				Options: []survey.Option{
					{Text: "Yes", IsCorrect: &istrue},
					{Text: "No", IsCorrect: &istrue},
				},
			},
		},
	}

	createdSurvey, err := surveyService.CreateSurvey(ctx, newSurvey)
	if err != nil {
		return fmt.Errorf("failed to create survey: %w", err)
	}

	log.Printf("Survey created successfully: %+v\n", createdSurvey)
	return nil
}

func fetchAndPrintSurveys(ctx context.Context, surveyService *service.SurveyService) error {
	fetchedSurvey, err := surveyService.GetSurveyByID(ctx, uuid.MustParse("e15f1b91-edc1-4aab-95e7-b9c9c4d83fb4"))
	if err != nil {
		return fmt.Errorf("failed to fetch survey: %w", err)
	}

	if fetchedSurvey == nil {
		log.Println("No survey found with the provided ID.")
	} else {
		log.Printf("Fetched survey: %+v\n", fetchedSurvey)
	}
	return nil
}
