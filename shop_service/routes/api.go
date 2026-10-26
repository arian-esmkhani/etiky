package routes

import (
	"fiber-app/handler"
	"fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupAPIRoutes configures all API routes with their middleware and handlers.
// Groups routes under /api prefix with compression, CORS, and rate limiting.
func SetupAPIRoutes(app *fiber.App, sliderHandler *handler.SliderHandler) {

	api := app.Group("/api",
		middleware.Compression(),
		middleware.CORS(),
	)

	// Serve static images from storage directory
	api.Static("/images", "../storage/images")

	// Slider routes
	slider := api.Group("/slider", middleware.DefaultLimiter())
	slider.Get("/discount", sliderHandler.GetProductsByDiscount)
	slider.Get("/best-selling", sliderHandler.GetBestSellingProducts)
	slider.Get("/cheapest", sliderHandler.GetCheapestProducts)
	slider.Post("/filter", sliderHandler.GetCheapestProductsByFilter)
}