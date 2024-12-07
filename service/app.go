package service

import (
	"NexaForm/config"
	"NexaForm/internal/otp"
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
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB() // Initialize the database and perform migrations
	app.setAuthService()
	app.setWalletService()
	return app, nil
}

func (a *AppContainer) RawRBConnection() *gorm.DB {
	return a.dbConn
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	// Initialize the database connection
	db, err := storage.NewPostgresGormConnection(a.cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database connection: ", err)
	}

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
func (a *AppContainer) setAuthService() {
	if a.authService != nil {
		return
	}

	a.authService = NewAuthService(otp.NewOps(storage.NewOtpRepo(a.dbConn)), user.NewOps(storage.NewUserRepo(a.dbConn), storage.NewRoleRepo(a.dbConn)), []byte(a.cfg.Server.TokenSecret),
		a.cfg.Server.TokenExpMinutes,
		a.cfg.Server.RefreshTokenExpMinutes)
}
func (a *AppContainer) AuthService() *AuthService {
	return a.authService
}
func (a *AppContainer) setWalletService() {
	if a.WalletSrvice() != nil {
		return
	}
	a.walletService = NewWalletService(otp.NewOps(storage.NewOtpRepo(a.dbConn)), wallet.NewOps(storage.NewWalletRepo(a.dbConn)))
}
func (a *AppContainer) WalletSrvice() *WalletService {
	return a.walletService
}
