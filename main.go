package main

import (
	"book-management/app/server"
	"book-management/util/config"
	"book-management/util/dbmigrate"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	if err := dbmigrate.RunMigrations(config); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	appServer := server.NewServer(config)
	err = appServer.Run()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
