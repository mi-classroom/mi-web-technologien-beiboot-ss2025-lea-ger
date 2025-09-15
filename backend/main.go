package main

import (
	"backend/db"
	"backend/migrate"
	"backend/router"
	"flag"
	"fmt"
	"os"
)

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run DB migrations")
	flag.Parse()

	if *migrateFlag {
		if err := migrate.RunMigrations(); err != nil {
			fmt.Fprintf(os.Stderr, "Migration failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migration erfolgreich ausgeführt.")
		return
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost" // Standard für lokale Entwicklung
	}
	dsn := fmt.Sprintf("host=%s port=5432 user=admin password=admin dbname=iptc-editor sslmode=disable", host)
	db.InitDB(dsn)

	server := router.SetupRouter()

	server.Run(":8080")
}
