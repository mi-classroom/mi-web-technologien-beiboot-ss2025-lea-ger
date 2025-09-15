package migrate

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

func RunMigrations() error {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost" // Standard für lokale Entwicklung
	}
	dsn := fmt.Sprintf("host=%s port=5432 user=admin password=admin dbname=iptc-editor sslmode=disable", host)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
