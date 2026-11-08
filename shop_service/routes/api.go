package routes

import (
	"fiber-app/handler"
	"fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupAPIRoutes configures all API routes with their middleware and handlers.
// Groups routes under /api prefix with compression, CORS, and rate limiting.
func SetupAPIRoutes(app *fiber.App, sliderHandler *handler.SliderHandler,
	finderHandler *handler.FinderHandler) {

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

	// Finder routes
	finder := api.Group("/finder", middleware.DefaultLimiter())

	// products
	finder.Get("/cheapest", finderHandler.FindCheapestProducts)
	finder.Get("/discount", finderHandler.FindProductsByDiscount)
	finder.Get("/best-selling", finderHandler.FindBestSellingProducts)
	finder.Post("/filter", finderHandler.FindCheapestProductsByFilter)

	// search
	finder.Get("/search", finderHandler.SearchProducts)
	finder.Get("/search/suggestions", finderHandler.GetSearchSuggestions)

	// articles
	finder.Get("/articles/:title", finderHandler.FindArticleByTitle)
	finder.Get("/articles/:slug", finderHandler.FindArticleBySlug)
}