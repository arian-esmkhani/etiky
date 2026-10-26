package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// Compression returns a middleware that compresses responses using gzip/deflate.
// Uses BestSpeed level for minimal CPU overhead and faster responses.
func Compression() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
}