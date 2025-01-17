package service

import (
	"NexaForm/config"
	"NexaForm/internal/otp"
	"NexaForm/internal/rbac"
	"NexaForm/internal/survey"
	"NexaForm/internal/user"
	"NexaForm/pkg/adapters/storage"
	"context"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg           config.Config
	dbConn        *gorm.DB
	authService   *AuthService
	loggerService *LoggerService
	surveyService *SurveyService
	fileService   *FileService
	rbacSercice   *RBACService
}

// NewAppContainer initializes the app container with services
func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB() // Initialize the database and perform migrations
	app.setAuthService()
	app.setLoggerService()
	app.setSurveyService()
	app.setFileService()
	app.setRBACService()
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

// setFileService initializes the FileService
func (a *AppContainer) setFileService() {
	if a.fileService != nil {
		return
	}

	fileService, err := NewFileService(
		survey.NewOps(storage.NewSurveyRepo(a.dbConn)),
		"localhost:9000", // Endpoint
		"minioadmin",     // Access Key
		"minioadmin",     // Secret Key
		"attachments",    // Bucket Name (can be customized)
		false,            // Use SSL (false for local setup)
		a.loggerService,
	)

	if err != nil {
		log.Fatalf("Failed to initialize FileService: %v", err)
	}

	a.fileService = fileService
	log.Println("FileService initialized successfully.")
}

// FileService returns the FileService instance
func (a *AppContainer) FileService() *FileService {
	return a.fileService
}
func (a *AppContainer) setRBACService() {
	if a.rbacSercice != nil {
		return
	}

	rbacService := NewRBACService(rbac.NewOps(storage.NewRBACRepo(a.dbConn)))

	a.rbacSercice = rbacService
	log.Println("RBACService initialized successfully.")
}

// FileService returns the FileService instance
func (a *AppContainer) RBACService() *RBACService {
	return a.rbacSercice
}

// func (a *AppContainer) GetLoggerForService(serviceName string) (*zap.Logger, error) {
// 	return a.loggerService.GetLoggerForService(serviceName)
// }

// func (a *AppContainer) AttachLoggerToContext(ctx context.Context, serviceName string, fiberCtx *fiber.Ctx) context.Context {
// 	return a.loggerService.AttachLoggerToContext(ctx, serviceName, fiberCtx)
// }
