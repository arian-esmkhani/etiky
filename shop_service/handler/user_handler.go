package handler

import (
	"context"
	"fiber-app/dto"
	"fiber-app/helper"
	"fiber-app/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler{
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 5*time.Second)
	defer cancel()

	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if badField := helper.ValidateLogin(loginDto); badField != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "invalid field: " + badField,
		})
	}

	err := h.service.CreateUser(ctx, loginDto) 

	if err!= nil {
		if err.Error() == "user already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"Message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Message": err.Error(),
		})
	}	

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (h *UserHandler) CreateAdminUser(c *fiber.Ctx) error {

	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 5*time.Second)
	defer cancel()

	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if badField := helper.ValidateLogin(loginDto); badField != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "invalid field: " + badField,
		})
	}

	err := h.service.CreateAdminUser(ctx, loginDto) 

	if err!= nil {
		if err.Error() == "Admin user already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"Message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Message": err.Error(),
		})
	}	

	return c.Status(201).JSON(fiber.Map{
		"message": "Admin User created successfully",
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {

	reqCtx := c.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 5*time.Second)
	defer cancel()

	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if badField := helper.ValidateLogin(loginDto); badField != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "invalid field: " + badField,
		})
	}

	access, err := h.service.Login(ctx, loginDto) 
	if err!= nil {
		switch err.Error() {
			case "user not found":
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"Message": err.Error(),
				})
			case "invalid password":
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Message": err.Error(),
				})
			case "failed to generate access token":
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Message": err.Error(),
				})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"Message": err.Error(),
				})
		}
		
	}


	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (h *UserHandler) CheckAuth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"authenticated": true,
	})
}