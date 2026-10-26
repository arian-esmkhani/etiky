package main

import (
	"fiber-app/config"
	"fiber-app/handler"
	"fiber-app/migration"
	"fiber-app/repo"
	"fiber-app/routes"
	"fiber-app/service"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Errorf("No .env file found, using system env")
	}

	// PostgreSQL connection configuration
	dbConfig := &config.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "pass"),
		Database: getEnv("DB_NAME", "etiky"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
	
	// Establish database connection pool
	db, err := config.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Database connection failed: ", err)
	}
	defer db.Close()
	
	// Apply database schema migrations
	if err := migration.RunMigrations(db); err != nil {
		log.Fatalf("Migration failed: ", err)
	}

	// Initialize Fiber with sonic JSON and production settings
	config.SetupSonic()
	app   := config.SetupFiber()
    
	// Dependency Injection Setup
	sliderRepo 	   := repo.NewSliderRepo(db)
	sliderService  := service.NewSliderService(sliderRepo)
	sliderHandler  := handler.NewSliderHandler(sliderService)

	// Register all API routes with the configured handler
	routes.SetupAPIRoutes(app, sliderHandler)

    // Expose Prometheus-style metrics endpoint
    app.Get("/metrics", monitor.New(monitor.Config{
        Title: "My App Metrics",
    }))

	// Start HTTP server
	log.Fatal(app.Listen(getEnv("APP_PORT", "5002")))
}

// getEnv retrieves environment variable or returns default value if not set.
// Logs warnings for missing variables to help with configuration debugging.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		log.Infof("Loaded environment variable: %s = %s", key, value)
		return value
	}
	log.Warnf("Environment variable %s not set, using default: %s", key, defaultValue)
	return defaultValue
}

// getEnvInt retrieves integer environment variable or returns default on parse failure.
// Validates format and logs warnings for non-integer values.
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intVal int
		if _, err := fmt.Sscanf(value, "%d", &intVal); err == nil {
			log.Infof("Loaded environment variable: %s = %d", key, intVal)
			return intVal
		}
		log.Warnf("Invalid integer format for %s: %s, using default: %d", key, value, defaultValue)
	}
	log.Warnf("Environment variable %s not set, using default: %d", key, defaultValue)
	return defaultValue
}