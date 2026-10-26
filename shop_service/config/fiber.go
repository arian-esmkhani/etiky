package config

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

// SetupFiber initializes and returns a new Fiber app instance with production-optimized defaults.
// Uses sonic for JSON (faster than default encoding/json) and configures timeouts for API stability.
func SetupFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       		  false,
		CaseSensitive: 		  true,
		StrictRouting: 		  true,
		JSONEncoder:   		  sonic.Marshal, // 2x faster JSON encoding
		JSONDecoder:   	  	  sonic.Unmarshal, // 2x faster JSON decoding
		Network:       		  "tcp4",
		BodyLimit:             1536 * 1024,  // 1.5MB = 1536KB
		Concurrency:           30, // Max concurrent connections per CPU
		DisableStartupMessage: false,
		ReduceMemoryUsage:     true,  // Lower memory footprint at cost of slight speed
		Immutable:             false,
		EnablePrintRoutes:     true, // Useful for debugging during dev
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second, // Prevents slow client attacks
		IdleTimeout:  30 * time.Second,
		AppName:               "etiky",
		StreamRequestBody:     false,
	})
	
	App = app
	return app
}