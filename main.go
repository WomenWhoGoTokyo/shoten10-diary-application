package main

import (
	"log"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/api"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/config"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/infrastructure/db"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/infrastructure/server"
)

func main() {
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer conn.Close()

	srv := server.NewServer(
		api.BuildRouter(conn),
	)

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
