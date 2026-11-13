package handler

import (
	"context"
	"encoding/json"
	"fiber-app/dto"
	"fiber-app/service"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	// Get text fields
	name := c.FormValue("name")
	description := c.FormValue("description")
	mark := c.FormValue("mark")
	priceStr := c.FormValue("price")
	discountStr := c.FormValue("discount", "0")
	countStr := c.FormValue("count")
	salesCountStr := c.FormValue("sales", "0")
	featuresStr := c.FormValue("features")
	tagsStr := c.FormValue("tags")

	// Parse numeric fields
	price, err := strconv.ParseUint(priceStr, 10, 32)
	if err != nil {
		log.Warnf("Invalid price value: %s, error: %v", priceStr, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "price is not valid",
		})
	}

	discount, err := strconv.ParseUint(discountStr, 10, 8)
	if err != nil {
		log.Warnf("Invalid discount value: %s, defaulting to 0, error: %v", discountStr, err)
		discount = 0
	}

	count, err := strconv.ParseUint(countStr, 10, 16)
	if err != nil {
		log.Warnf("Invalid count value: %s, error: %v", countStr, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "count is not valid"})
	}

	salesCount, err := strconv.ParseUint(salesCountStr, 10, 64)
	if err != nil {
		log.Infof("Invalid sales count value: %s, defaulting to 0, error: %v", salesCountStr, err)
		salesCount = 0
	}

	// Parse array fields (JSON string to []string)
	var features []string
	if featuresStr != "" {
		if err := json.Unmarshal([]byte(featuresStr), &features); err != nil {
			log.Warnf("Invalid features JSON: %s, error: %v", featuresStr, err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "features must be a JSON array",
			})
		}
	}

	var tags []string
	if tagsStr != "" {
		if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
			log.Warnf("Invalid tags JSON: %s, error: %v", tagsStr, err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "tags must be a JSON array",
			})
		}
	}

	// Main image
	imageFile, err := c.FormFile("image")
	if err != nil {
		log.Errorf("Failed to get main image: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to process main image",
		})
	}

	// Gallery images
	var imageFiles []*multipart.FileHeader
	form, err := c.MultipartForm()
	if err != nil {
		log.Infof("No multipart form data received: %v", err)
		imageFiles = []*multipart.FileHeader{}
	} else {
		imageFiles = form.File["images"]
		if imageFiles == nil {
			log.Warnf("No 'images' field found in form data")
			imageFiles = []*multipart.FileHeader{}
		}
	}

	// Build request for service
	req := dto.CreateProductReq{
		Name:        name,
		Description: description,
		Mark:        mark,
		Price:       uint32(price),
		Discount:    uint8(discount),
		Count:       uint16(count),
		SalesCount:  salesCount,
		Features:    features,
		Tags:        tags,
		ImageFile:   imageFile,
		ImageFiles:  imageFiles,
	}

	// Call service
	err = h.service.CreateProduct(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	// Success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	// Get text fields
	name := c.FormValue("name")
	description := c.FormValue("description")
	mark := c.FormValue("mark")
	priceStr := c.FormValue("price")
	discountStr := c.FormValue("discount", "0")
	countStr := c.FormValue("count")
	salesCountStr := c.FormValue("sales", "0")
	featuresStr := c.FormValue("features")
	tagsStr := c.FormValue("tags")

	// Parse numeric fields
	price, err := strconv.ParseUint(priceStr, 10, 32)
	if err != nil {
		log.Warnf("Invalid price value: %s, error: %v", priceStr, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "price is not valid",
		})
	}

	discount, err := strconv.ParseUint(discountStr, 10, 8)
	if err != nil {
		log.Warnf("Invalid discount value: %s, defaulting to 0, error: %v", discountStr, err)
		discount = 0
	}

	count, err := strconv.ParseUint(countStr, 10, 16)
	if err != nil {
		log.Warnf("Invalid count value: %s, error: %v", countStr, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "count is not valid",
		})
	}

	salesCount, err := strconv.ParseUint(salesCountStr, 10, 64)
	if err != nil {
		log.Infof("Invalid sales count value: %s, defaulting to 0, error: %v", salesCountStr, err)
		salesCount = 0
	}

	// Parse array fields (JSON string to []string)
	var features []string
	if featuresStr != "" {
		if err := json.Unmarshal([]byte(featuresStr), &features); err != nil {
			log.Warnf("Invalid features JSON: %s, error: %v", featuresStr, err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "features must be a JSON array",
			})
		}
	}

	var tags []string
	if tagsStr != "" {
		if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
			log.Warnf("Invalid tags JSON: %s, error: %v", tagsStr, err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "tags must be a JSON array",
			})
		}
	}

	imageFile, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		log.Errorf("Failed to get main image: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to process main image",
		})
	}

	var imageFiles []*multipart.FileHeader
	form, err := c.MultipartForm()
	if err != nil {
		if err != http.ErrMissingFile && err != http.ErrNotMultipart {
			log.Errorf("Failed to parse multipart form: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "failed to process gallery images",
			})
		}
		imageFiles = []*multipart.FileHeader{}
	} else {
		imageFiles = form.File["images"]
		if imageFiles == nil {
			imageFiles = []*multipart.FileHeader{}
		}
	}

	req := dto.CreateProductReq{
		Name:        name,
		Description: description,
		Mark:        mark,
		Price:       uint32(price),
		Discount:    uint8(discount),
		Count:       uint16(count),
		SalesCount:  salesCount,
		Features:    features,
		Tags:        tags,
		ImageFile:   imageFile,
		ImageFiles:  imageFiles,
	}

	err = h.service.UpdateProduct(ctx, id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (h *ProductHandler) SoftDeleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	err = h.service.SoftDeleteProduct(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product soft deleted successfully",
	})
}

func (h *ProductHandler) HardDeleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	err = h.service.HardDeleteProduct(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product hard deleted successfully",
	})
}

func (h *ProductHandler) RestoreProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	err = h.service.RestoreProduct(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product restored successfully",
	})
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	} else if id < 1 || id > 999999 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	product, err := h.service.GetProductByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *ProductHandler) AddProductImages(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		log.Warnf("No multipart form data received: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "no images provided",
		})
	}

	imageFiles := form.File["images"]
	if imageFiles == nil || len(imageFiles) == 0 {
		log.Warnf("No 'images' field found in form data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "no images provided",
		})
	}

	err = h.service.AddProductImages(ctx, id, imageFiles)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": internalServerErrorMassage,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product images added successfully",
	})
}

func (h *ProductHandler) RemoveProductImages(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
	defer cancel()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid product ID: %s, error: %v", c.Params("id"), err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	var req struct {
		ImagePaths []string `json:"imagePaths"`
	}

	if err := c.BodyParser(&req); err != nil {
		log.Warnf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if len(req.ImagePaths) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "no image paths provided",
		})
	}

	err = h.service.RemoveProductImages(ctx, id, req.ImagePaths)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product images removed successfully",
	})
}