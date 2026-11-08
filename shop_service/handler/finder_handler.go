package handler

import (
	"context"
	"fiber-app/dto"
	"fiber-app/helper"
	"fiber-app/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// FinderHandler handles HTTP requests for product discovery, search and articles.
type FinderHandler struct {
	service service.IFinderService
}

func NewFinderHandler(service service.IFinderService) *FinderHandler {
	return &FinderHandler{service: service}
}

const invalidCategory string = "invalid category format"

// ---------- products ----------

func (h *FinderHandler) FindCheapestProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	page := helper.GetPage(c)

	result, err := h.service.FindCheapestProducts(ctx, page)
	if err != nil {
		log.Errorf("Failed to find cheapest products (page: %d): %v", page, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	if len(result.Products) < 1 {
		log.Warn("No cheapest products found with page: %d", page)
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *FinderHandler) FindProductsByDiscount(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	category := c.Query("category", "")
	if !helper.ValidFilterValue(category) {
		log.Warn(invalidCategory)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": invalidCategory,
		})
	}

	page := helper.GetPage(c)

	result, err := h.service.FindProductsByDiscount(ctx, category, page)
	if err != nil {
		log.Errorf("Failed to find discounted products (category: %s, page: %d): %v", category, page, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}

	if len(result.Products) < 1 {
		log.Warnf("No discount products found with page: %d", page)
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *FinderHandler) FindBestSellingProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	category := c.Query("category", "")
	if !helper.ValidFilterValue(category) {
		log.Warn(invalidCategory)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": invalidCategory,
		})
	}

	var page uint8 = helper.GetPage(c)

	result, err := h.service.FindBestSellingProducts(ctx, category, page)
	if err != nil {
		log.Errorf("failed to find best selling products (category: %s, page: %d): %v", category, page, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}

	if len(result.Products) < 1 {
		log.Warnf("No best selling products found with page: %d", page)
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *FinderHandler) FindCheapestProductsByFilter(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	var filter dto.ProductFilterDto
	if err := c.BodyParser(&filter); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	var inValidFiled string = helper.ValidateProductFilter(filter)
	if inValidFiled != "" {
		log.Error(invalidCategory + ":" + inValidFiled)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	page := helper.GetPage(c)

	result, err := h.service.FindCheapestProductsByFilter(ctx, filter, page)
	if err != nil {
		log.Errorf("Failed to find cheapest products by filter (filter: %+v, page: %d): %v", filter, page, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}

	if len(result.Products) < 1 {
		log.Warnf("No best cheapest products by filter found with page: %d", page)
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// ---------- search ----------

func (h *FinderHandler) SearchProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	category := c.Query("category", "")
	if !helper.ValidFilterValue(category) {
		log.Warn(invalidCategory)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": invalidCategory,
		})
	}

	queryStr := c.Query("q", "")
	if queryStr != "" && !helper.ValidText(queryStr) {
		log.Warn("Invalid query format")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query format",
		})
	}

	lowPrice, maxPrice, badField := helper.ReadPriceRange(c)
	if badField != "" {
		log.Warn("Invalid price format"+badField)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid price format"+badField,
		})
	}

	page := helper.GetPage(c)

	result, err := h.service.SearchProducts(ctx, queryStr, category, lowPrice, maxPrice, page)
	if err != nil {
		log.Errorf("Failed to search products (query: %s, category: %s, page: %d): %v",
			queryStr, category, page, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}

	if len(result.Products) < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"massage" : "no product found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *FinderHandler) GetSearchSuggestions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 2*time.Second)
	defer cancel()

	queryStr := c.Query("q", "")
	if queryStr != "" && !helper.ValidText(queryStr) {
		log.Warn("Invalid query format")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query format",
		})
	}

	result, err := h.service.GetSearchSuggestions(ctx, queryStr)
	if err != nil {
		log.Errorf("Failed to get search suggestions (query: %s): %v", queryStr, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// ---------- articles ----------

func (h *FinderHandler) FindArticleBySlug(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	slug := c.Params("slug")
	if !helper.ValidSlug(slug) {
		log.Warn("invalid slug format"+slug)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid slug format"+slug,
		})
	}

	article, err := h.service.FindArticleBySlug(ctx, slug)
	if err != nil {
		log.Errorf("Failed to find article by slug (slug: %s): %v", slug, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}
	if article == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"massage" : "article not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(article)
}

func (h *FinderHandler) FindArticleByTitle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	title := c.Params("title")
	if !helper.ValidText(title) {
		log.Warn("invalid title format"+title)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid title format"+title,
		})
	}

	article, err := h.service.FindArticleByTitle(ctx, title)
	if err != nil {
		log.Errorf("Failed to find article by title (title: %s): %v", title, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage" : internalServerErrorMassage,
		})
	}
	if article == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"massage" : "article not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(article)
}