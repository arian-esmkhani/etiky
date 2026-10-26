package config

import (
    "context"
    "fmt"
    "time"
    
	"github.com/gofiber/fiber/v2/log"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    Database string
    SSLMode  string
}

// NewConnection creates and validates a new PostgreSQL connection pool with default settings.
// Returns an error if connection fails or ping timeout occurs.
func NewConnection(cfg *Config) (*pgxpool.Pool, error) {

    connString := fmt.Sprintf(
        "postgres://%s:%s@%s:%d/%s?sslmode=%s",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode,
    )
    
    dbConfig, err := pgxpool.ParseConfig(connString)
    if err != nil {
		log.Errorf("Failed to parse database config: %v", err)
        return nil, err
    }
    
    // Connection pool settings
    dbConfig.MaxConns = 20
    dbConfig.MinConns = 5
    dbConfig.MaxConnLifetime = time.Hour
    dbConfig.MaxConnIdleTime = 30 * time.Minute
    
    pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
    if err != nil {
		log.Errorf("Failed to create connection pool: %v", err)
        return nil, err
    }
    
    // Verify database connectivity
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := pool.Ping(ctx); err != nil {
        log.Errorf("Database ping failed: %v", err)
        return nil, err
    }
    
	log.Info("Database connected successfully")

    return pool, nil
}