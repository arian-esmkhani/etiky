package handler

import (
	"context"
	"fiber-app/dto"
	"fiber-app/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type SliderHandler struct {
	service service.ISliderService
}

// NewSliderHandler creates a new slider HTTP handler with injected service.
func NewSliderHandler(service service.ISliderService) *SliderHandler {
	return &SliderHandler{service: service}
}

// Common Error massages
const internalServerErrorMassage string = "internal server error"
const noContentMassage string = "no products found"

// Returns products with highest discounts, sorted by discount DESC.
// Response: 200 OK with []dto.PreviewProductDto, 204 NoContent if empty, 500 on error.
func (h *SliderHandler) GetProductsByDiscount(c *fiber.Ctx) error {
	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, err := h.service.GetProductsByDiscount(ctx)
	if err != nil {
		log.Errorf("Failed to get discounted products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	if len(slider) < 1 {
		log.Warn("No discounted products found")
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(slider)
}

// Returns products with lowest final price, sorted by price ASC.
// Response: 200 OK with []dto.PreviewProductDto, 204 NoContent if empty, 500 on error.
func (h *SliderHandler) GetCheapestProducts(c *fiber.Ctx) error {
	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, err := h.service.GetCheapestProducts(ctx)
	if err != nil {
		log.Errorf("Failed to get cheapest products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	if len(slider) < 1 {
		log.Warn("No cheapest products found")
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(slider)
}

// Returns products with highest sales_count, sorted by sales DESC.
// Response: 200 OK with []dto.PreviewProductDto, 204 NoContent if empty, 500 on error.
func (h *SliderHandler) GetBestSellingProducts(c *fiber.Ctx) error {
	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, err := h.service.GetBestSellingProducts(ctx)
	if err != nil {
		log.Errorf("Failed to get best selling products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	if len(slider) < 1 {
		log.Warn("No best selling products found")
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(slider)
}

// Returns products with filter
// Response: 200 OK with []dto.PreviewProductDto, 204 NoContent if empty, 500 on error.
func (h *SliderHandler) GetCheapestProductsByFilter(c *fiber.Ctx) error {
	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	var filter dto.ProductFilterDto
	if err := c.BodyParser(&filter); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	slider, err := h.service.GetCheapestProductsByFilter(ctx, filter)
	if err != nil {
		log.Errorf("Failed to get cheapest products by filter products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	if len(slider) < 1 {
		log.Warn("No cheapest products by filter products found")
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": noContentMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(slider)
}