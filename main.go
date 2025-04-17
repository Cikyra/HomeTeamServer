package main

import (
	"log"
	"net/http"
	"os"

	"HomeTeamServer/handlers"
	"HomeTeamServer/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Handle Migrations
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)
	sugar := logger.Sugar()

	userHandler := handlers.NewUserHandler(db, sugar)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /user", userHandler.GetUsers)
	mux.HandleFunc("GET /user/{id}", userHandler.GetUser)
	mux.HandleFunc("POST /user", userHandler.CreateUser)

	port := ":1989" // TODO: Use an env variable instead
	sugar.Infof("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux)) // TODO: Do something better here
}

/**
Go HTTP Server
	- Endpoints
		- REST API
			- GET (gets stuff), POST (creates stuff), PUT/PATCH (updates stuff), DELETE (deletes stuff)
			- GET https://localhost:8080/schools
				- Returns a list of all schools
				- GET https://localhost:8080/schools?page={0}&pageSize={100}
			- POST https://localhost:8080/schools
				- Request will have a body (usually JSON)
	- Talk to a DB
		- Have models
		- Probably do stuff with SQL (or maybe use Gorm? - Gorm is an ORM)
*/
