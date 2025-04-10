package main

import (
	"fmt"
	"log"
	"net/http"

	"HomeTeamServer/handlers"
)

func main() {
	userHandler := handlers.NewUserHandler()

	http.Handle("/user", userHandler)

	port := ":1989" // TODO: Use an env variable instead
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // TODO: Do something better here
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
