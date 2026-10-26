package migration

import (
	"embed"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed *.sql
var migrationFiles embed.FS

// RunMigrations executes database schema migrations
func RunMigrations(pool *pgxpool.Pool) error {
	// Convert pgxpool to standard sql.DB for migrate compatibility
	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()

	// Create PostgreSQL driver for migrate
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Errorf("Failed to create migration driver: %v", err)
		return err
	}

	// Load migration files from embedded FS
	src, err := iofs.New(migrationFiles, ".")
	if err != nil {
		log.Errorf("Failed to load migration files: %v", err)
		return err
	}

	// Create migrate instance
	m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		log.Errorf("Failed to create migrate instance: %v", err)
		return err
	}

	// Run pending migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Errorf("Migration failed: %v", err)
		return err
	}

	log.Info("Database migrations completed successfully")
	return nil
}