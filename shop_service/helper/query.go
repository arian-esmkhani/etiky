package helper

import (
	"strconv"

	"fiber-app/dto"

	"github.com/gofiber/fiber/v2"
)

const maxPrice = 999_999_999

// ReadProductFilter reads and validates product filter query params.
// On failure it returns the name of the invalid field.
func ReadProductFilter(c *fiber.Ctx) (dto.ProductFilterDto, string) {
	filter := dto.ProductFilterDto{
		Mark:    c.Query("mark", ""),
		Feature: c.Query("feature", ""),
		Tag:     c.Query("tag", ""),
	}

	fields := []struct {
		name  string
		value string
	}{
		{"mark", filter.Mark},
		{"feature", filter.Feature},
		{"tag", filter.Tag},
	}

	for _, f := range fields {
		if !ValidFilterValue(f.value) {
			return dto.ProductFilterDto{}, f.name
		}
	}

	return filter, ""
}

// ReadPriceRange reads low_price and max_price.
// On failure it returns the name of the invalid field.
func ReadPriceRange(c *fiber.Ctx) (low, high uint32, badField string) {
	lowV, err := strconv.ParseUint(c.Query("low_price", "0"), 10, 32)
	if err != nil || lowV > maxPrice {
		return 0, 0, "low_price"
	}

	highV, err := strconv.ParseUint(c.Query("max_price", "0"), 10, 32)
	if err != nil || highV > maxPrice {
		return 0, 0, "max_price"
	}

	if lowV > 0 && highV > 0 && lowV > highV {
		return 0, 0, "low_price"
    }

	return uint32(lowV), uint32(highV), ""
}