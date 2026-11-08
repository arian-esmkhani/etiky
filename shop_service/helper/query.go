package helper

import (
	"fiber-app/dto"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const maxPrice = 999_999_999

// ValidateProductFilter validates all fields of a product filter.
// It returns the name of the first invalid field, or "" if everything is valid.
func ValidateProductFilter(filter dto.ProductFilterDto) string {
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
			return f.name
		}
	}

	return ""
}

// ValidateLogin validates all fields of a login request.
// It returns the name of the first invalid field, or "" if everything is valid.
func ValidateLogin(login dto.LoginDto) string {
	fields := []struct {
		name  string
		value string
	}{
		{"number", login.Number},
		{"password", login.Password},
	}

	for _, f := range fields {
		var isValid bool
		
		switch f.name {
		case "number":
			isValid = PhoneRegex.MatchString(f.value)
		case "password":
			isValid = PasswordRegex.MatchString(f.value)
		}

		if !isValid {
			return f.name
		}
	}

	return ""
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