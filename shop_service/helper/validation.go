package helper

import (
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Regexes are compiled once at startup instead of on every request.
var (
	FilterRegex = regexp.MustCompile(`^[\p{L}\p{N}\s\-_\x{200C}]{1,40}$`)
	TextRegex   = regexp.MustCompile(`^[\p{L}\p{N}\s\-_\x{200C}]{1,190}$`)
	SlugRegex   = regexp.MustCompile(`^[\p{L}\p{N}\-_\x{200C}]{1,200}$`)
)

// ValidFilterValue accepts an empty value (filter not set) or a value matching FilterRegex.
func ValidFilterValue(v string) bool {
	return v == "" || FilterRegex.MatchString(v)
}

// ValidText accepts a non-empty value matching TextRegex.
func ValidText(v string) bool {
	return TextRegex.MatchString(v)
}

// ValidSlug accepts a non-empty value matching SlugRegex.
func ValidSlug(v string) bool {
	return SlugRegex.MatchString(v)
}

func GetPage(c *fiber.Ctx) uint8 {
	pageStr := c.Query("page", "0")
	page, err := strconv.ParseUint(pageStr, 10, 8)
	if err != nil {
		log.Warnf("Invalid page parameter: %s, defaulting to 0", pageStr)
		return 0
	}
	return uint8(page)
}