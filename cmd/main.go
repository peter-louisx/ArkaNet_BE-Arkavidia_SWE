package main

import (
	"Go-Starter-Template/internal/config"
	databaseconf "Go-Starter-Template/internal/config/database_config"
	"os"
)

var addr = os.Getenv("APP_URL")

func main() {
	db, err := databaseconf.ConnectDB()
	if err != nil {
		panic(err)
	}

	app, err := config.NewApp(db)
	if err != nil {
		panic(err)
	}
	if addr == "" {
		addr = "0.0.0.0:8080"
	}
	err = app.Listen(addr)
	if err != nil {
		panic(err)
	}
}
