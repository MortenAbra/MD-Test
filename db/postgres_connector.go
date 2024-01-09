package db

import (
	"fmt"
	"log"
	"media-devoted/config"
	dbmodels "media-devoted/types/db_models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db     *gorm.DB
	config *config.DatabaseConfig
}

var (
	instance *Database
	once     sync.Once
)

// NewDatabase creates and returns a new Database instance with the given options applied.
func NewDatabase(options ...func(*config.DatabaseConfig)) *Database {
	once.Do(func() {
		cfg := &config.DatabaseConfig{}
		for _, option := range options {
			option(cfg)
		}

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		log.Println("Connection established")

		instance = &Database{db: db, config: cfg}
		instance.migrate()
	})
	return instance
}

// WithHost sets the database host.
func WithHost(host string) func(*config.DatabaseConfig) {
	return func(cfg *config.DatabaseConfig) {
		cfg.Host = host
	}
}

// WithPort sets the database port.
func WithPort(port string) func(*config.DatabaseConfig) {
	return func(cfg *config.DatabaseConfig) {
		cfg.Port = port
	}
}

// WithUser sets the database user.
func WithUser(user string) func(*config.DatabaseConfig) {
	return func(cfg *config.DatabaseConfig) {
		cfg.User = user
	}
}

// WithPassword sets the database password.
func WithPassword(password string) func(*config.DatabaseConfig) {
	return func(cfg *config.DatabaseConfig) {
		cfg.Password = password
	}
}

// WithDBName sets the database name.
func WithDBName(dbName string) func(*config.DatabaseConfig) {
	return func(cfg *config.DatabaseConfig) {
		cfg.DBName = dbName
	}
}

// migrate handles the database migration.
func (d *Database) migrate() {
	// AutoMigrate models
	if err := d.db.AutoMigrate(&dbmodels.RocketModel{}); err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}
	log.Println("Migrated successfully")
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	if instance == nil || instance.db == nil {
		log.Fatal("Database not initialized")
	}
	return instance.db
}
