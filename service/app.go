package service

import (
	"NexaForm/config"
	"NexaForm/internal/otp"
	"NexaForm/internal/survey"
	"NexaForm/internal/user"
	"NexaForm/internal/wallet"
	"NexaForm/pkg/adapters/storage"
	"context"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg           config.Config
	dbConn        *gorm.DB
	authService   *AuthService
	walletService *WalletService
	loggerService *LoggerService
	surveyService *SurveyService
}

// NewAppContainer initializes the app container with services
func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB() // Initialize the database and perform migrations
	app.setAuthService()
	app.setWalletService()
	app.setLoggerService()
	app.setSurveyService()
	return app, nil
}

func (a *AppContainer) RawRBConnection() *gorm.DB {
	return a.dbConn
}

// mustInitDB initializes the database and performs migrations
func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	// Initialize the database connection
	db, err := storage.NewPostgresGormConnection(a.cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database connection: ", err)
	}
	db = db.Debug()

	// Assign the connection to the app container
	a.dbConn = db

	// Add PostgreSQL extensions
	err = storage.AddExtension(a.dbConn)
	if err != nil {
		log.Fatal("Failed to add extension: ", err)
	}

	// Perform migrations
	err = storage.Migrate(a.dbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	err = storage.Seed(context.Background(), a.dbConn)
	if err != nil {
		log.Fatal("Seeding failed: ", err)
	}
	log.Println("Database initialized and migrated successfully.")
}

// setAuthService initializes the AuthService
func (a *AppContainer) setAuthService() {
	if a.authService != nil {
		return
	}

	a.authService = NewAuthService(otp.NewOps(storage.NewOtpRepo(a.dbConn)), user.NewOps(storage.NewUserRepo(a.dbConn), storage.NewRoleRepo(a.dbConn)), []byte(a.cfg.Server.TokenSecret),
		a.cfg.Server.TokenExpMinutes,
		a.cfg.Server.RefreshTokenExpMinutes)
}

// AuthService returns the AuthService instance
func (a *AppContainer) AuthService() *AuthService {
	return a.authService
}
func (a *AppContainer) setWalletService() {
	if a.WalletSrvice() != nil {
		return
	}
	walletRepo := storage.NewWalletRepo(a.dbConn)
	walletOps := wallet.NewOps(walletRepo)
	roleRepo := storage.NewRoleRepo(a.dbConn)
	userRepo := storage.NewUserRepo(a.dbConn)
	userOps := user.NewOps(userRepo, roleRepo)
	a.walletService = NewWalletService(userOps, walletOps)
}
func (a *AppContainer) WalletSrvice() *WalletService {
	return a.walletService
}

// setLoggerService initializes the LoggerService using the logging configuration
func (a *AppContainer) setLoggerService() {
	if a.loggerService != nil {
		return
	}

	loggingConfig := a.cfg.Logging.Loggers // Retrieve array of logger configs
	loggerService, err := NewLoggerService(loggingConfig, a.cfg.Logging.LokiURL)
	if err != nil {
		log.Fatalf("Failed to initialize LoggerService: %v", err)
	}

	a.loggerService = loggerService
	log.Println("LoggerService initialized successfully.")
}

// LoggerService returns the LoggerService instance
func (a *AppContainer) LoggerService() *LoggerService {
	return a.loggerService
}
func (a *AppContainer) setSurveyService() {
	if a.surveyService != nil {
		return
	}
	a.surveyService = NewSurveyService(survey.NewOps(storage.NewSurveyRepo(a.dbConn)))
}
func (a *AppContainer) SurveyService() *SurveyService {
	return a.surveyService
}
