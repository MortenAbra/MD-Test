package main

import (
	"log"
	"media-devoted/config"
	"media-devoted/db"
	"media-devoted/options"
	"media-devoted/routing"

	"github.com/gin-contrib/cors"
)

func main() {

	dbConfig := config.LoadConfig("config\\dbconfig.yml")

	_ = db.NewDatabase(
		db.WithHost(dbConfig.Database.Host),
		db.WithPassword(dbConfig.Database.Password),
		db.WithUser(dbConfig.Database.User),
		db.WithDBName(dbConfig.Database.DBName),
		db.WithPort(dbConfig.Database.Port),
	)

	engine := options.ServerOptions(
		options.WithHost("0.0.0.0"),
		options.WithPort("8080"),
		options.WithRoutes(routing.AddRocketGroup),
		options.WithCorsRules(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST"},
		}),
	)

	if err := engine.Start(); err != nil {
		log.Panic("Error during start of engine...")
	}

}
