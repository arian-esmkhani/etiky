package handler

import (
	"context"
	"fiber-app/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CartHandler struct {
	service service.ICartService
}

func NewCartHandler(service service.ICartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	// Get user ID from context (set by auth middleware)
	userId, ok := c.Locals("userId").(uint64)
	if !ok {
		log.Errorf("User ID not found in context or have problem")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	} else if userId < 1 || userId > 999999 {
		log.Errorf("User ID not valid")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	// Parse product ID from params
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	} else if productId < 1 || productId > 999999 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	// Call service
	err = h.service.Add(ctx, userId, productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product added to cart successfully",
	})
}

func (h *CartHandler) UpdateCartItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	// Get user ID from context
	userId, ok := c.Locals("userId").(uint64)
	if !ok {
		log.Errorf("User ID not found in context or have problem")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	} else if userId < 1 || userId > 999999 {
		log.Errorf("User ID not valid")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	// Parse product ID from params
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	} else if productId < 1 || productId > 999999 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	// Parse status from query or body
	var req struct {
		Status *bool `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		log.Warnf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}
	if req.Status == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "status field is required",
		})
	}

	// Call service
	err = h.service.Update(ctx, userId, productId, *req.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cart item updated successfully",
	})
}

func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	userId, ok := c.Locals("userId").(uint64)
	if !ok {
		log.Errorf("User ID not found in context or have problem")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	} else if userId < 1 || userId > 999999 {
		log.Errorf("User ID not valid")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	// Parse product ID from params
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	} else if productId < 1 || productId > 999999 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	// Call service
	err = h.service.HardDelete(ctx, userId, productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product removed from cart successfully",
	})
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	// Get user ID from context
	userId, ok := c.Locals("userId").(uint64)
	if !ok {
		log.Errorf("User ID not found in context or have problem")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	} else if userId < 1 || userId > 999999 {
		log.Errorf("User ID not valid")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	// Call service
	cart, err := h.service.Get(ctx, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	if len(cart.CartDto) < 1 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "no products found",
		})
	}


	return c.Status(fiber.StatusOK).JSON(cart)
}

func (h *CartHandler) GetProductQuantity(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
    defer cancel()

    // Get user ID from context
	userId, ok := c.Locals("userId").(uint64)
	if !ok {
		log.Errorf("User ID not found in context or have problem")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	} else if userId < 1 || userId > 999999 {
		log.Errorf("User ID not valid")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

    // Parse product ID from params
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	} else if productId < 1 || productId > 999999 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

    // Call service
    quantity, err := h.service.GetProductQuantity(ctx, userId, productId)
    if err != nil {
        log.Errorf("Failed to get product quantity: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get product quantity",
		})
    }

    return c.Status(200).JSON(fiber.Map{
        "quantity":  quantity,
        "inCart":    quantity > 0,
    })
}