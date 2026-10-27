package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up|down>")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)

	if err != nil {
		log.Fatal("failed to create migrate instance:", err)
	}

	switch os.Args[1] {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("migration completed")
}
